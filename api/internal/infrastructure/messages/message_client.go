package messages

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type MessageClient interface {
	IncrementAndGet(ctx context.Context, key string, expirationSeconds int) (int64, error)
	SetUpAndRunWorker(ctx context.Context, streamName, groupName, consumerID string) error
	EnqueueMessage(ctx context.Context, message map[string]interface{}) error
}

type messageClient struct {
	client *redis.Client
}

func NewMessageClient(ctx context.Context, redisAddress string) (MessageClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddress,
	})
	err := rdb.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}
	return &messageClient{client: rdb}, nil
}

func (r *messageClient) IncrementAndGet(ctx context.Context, key string, expirationSeconds int) (int64, error) {
	// If key doesn't exist, create and set to 1
	val, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	// Set expiration only when key is newly created
	// Expiration is expirationSeconds after it gets created
	if val == 1 {
		err = r.client.Expire(ctx, key, time.Duration(expirationSeconds)*time.Second).Err()
		if err != nil {
			return 0, err
		}
	}
	return val, nil
}

func (r *messageClient) SetUpAndRunWorker(ctx context.Context, streamName, groupName, consumerID string) error {
	// This creates a stream if it doesn't exist and a consumer group for the stream
	// $ means only consume new messages
	err := r.client.XGroupCreateMkStream(
		ctx,
		streamName,
		groupName,
		"$",
	).Err()
	if err != nil && !isGroupExistsErr(err) {
		log.Fatalf("Redis: failed to created grouop: %v", err)
	}

	log.Println("Redis worker set up")

	for {
		msgs, err := r.client.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    groupName,
			Consumer: consumerID,
			Streams:  []string{streamName, ">"},
			Count:    1,
			Block:    30 * time.Second,
		}).Result()

		if err != nil && err != redis.Nil {
			log.Printf("Redis: error reading from stream: %v", err)
			continue
		}

		if len(msgs) == 0 {
			log.Println("Redis: no new messages, continuing...")
			continue
		}

		for _, msg := range msgs[0].Messages {
			// TODO: process message
			log.Printf("Redis: processing message ID %s with values %v", msg.ID, msg.Values)
			// Acknowledge message
			err = r.client.XAck(ctx, streamName, groupName, msg.ID).Err()
			if err != nil {
				log.Printf("Redis: error acknowledging message ID %s: %v", msg.ID, err)
			}
		}
	}
}

func isGroupExistsErr(err error) bool {
	return err != nil &&
		(err.Error() == "BUSYGROUP Consumer Group name already exists")
}

func (r *messageClient) EnqueueMessage(ctx context.Context, message map[string]interface{}) error {
	log.Printf("Enqueuing message %v", message)
	_, err := r.client.XAdd(ctx, &redis.XAddArgs{
		Stream: "mystream",
		Values: message,
	}).Result()
	return err
}
