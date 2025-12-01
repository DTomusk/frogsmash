package messages

import (
	"context"
	"log"
)

type MessageProducer interface {
	EnqueueMessage(ctx context.Context, message map[string]interface{}) error
}

type messageProducer struct {
	client RedisClient
}

func NewMessageProducer(client RedisClient) (MessageProducer, error) {
	return &messageProducer{
		client: client,
	}, nil
}

func (r *messageProducer) EnqueueMessage(ctx context.Context, message map[string]interface{}) error {
	log.Printf("Enqueuing message %v", message)
	return r.client.AddMessage(ctx, message)
}
