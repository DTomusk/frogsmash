package redis

import (
	"context"
	"crypto/tls"
	"frogsmash/internal/infrastructure/messages"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	InitStream(ctx context.Context) error
	GetMessages(ctx context.Context) ([]*messages.Message, error)
	AcknowledgeMessage(messageID string, ctx context.Context) error
	AddMessage(ctx context.Context, message map[string]interface{}) error
	IncrementAndGet(ctx context.Context, key string, expirationSeconds int) (int64, error)
}

type redisClient struct {
	client     *redis.Client
	streamName string
	groupName  string
	consumerID string
}

func NewRedisClient(redisAddress, redisUsername, redisPassword, streamName, groupName, consumerID string) RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:      redisAddress,
		Username:  redisUsername,
		Password:  redisPassword,
		TLSConfig: &tls.Config{},
	})
	return &redisClient{
		client:     rdb,
		streamName: streamName,
		groupName:  groupName,
		consumerID: consumerID,
	}
}

func (r *redisClient) InitStream(ctx context.Context) error {
	err := r.client.XGroupCreateMkStream(
		ctx,
		r.streamName,
		r.groupName,
		"$",
	).Err()
	if err != nil && !isGroupExistsErr(err) {
		return err
	}
	return nil
}

func isGroupExistsErr(err error) bool {
	return err != nil &&
		(err.Error() == "BUSYGROUP Consumer Group name already exists")
}

func (r *redisClient) GetMessages(ctx context.Context) ([]*messages.Message, error) {
	msgs, err := r.client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    r.groupName,
		Consumer: r.consumerID,
		Streams:  []string{r.streamName, ">"},
		Count:    10,
		Block:    0,
	}).Result()

	if err != nil && err != redis.Nil {
		return nil, err
	}

	var messageList []*messages.Message
	for _, msg := range msgs {
		for _, m := range msg.Messages {
			messageList = append(messageList, &messages.Message{
				ID:     m.ID,
				Values: m.Values,
			})
		}
	}

	return messageList, nil
}

func (r *redisClient) AcknowledgeMessage(messageID string, ctx context.Context) error {
	return r.client.XAck(ctx, r.streamName, r.groupName, messageID).Err()
}

func (r *redisClient) AddMessage(ctx context.Context, message map[string]interface{}) error {
	return r.client.XAdd(ctx, &redis.XAddArgs{
		Stream: r.streamName,
		Values: message,
	}).Err()
}

func (r *redisClient) IncrementAndGet(ctx context.Context, key string, expirationSeconds int) (int64, error) {
	// Lua script to increment a key and set expiration if it's new
	var incrWithExpireScript = redis.NewScript(`
	local current = redis.call("INCR", KEYS[1])
	if current == 1 then
		redis.call("EXPIRE", KEYS[1], ARGV[1])
	end
	return current
	`)

	result, err := incrWithExpireScript.Run(ctx, r.client, []string{key}, expirationSeconds).Result()

	return result.(int64), err
}
