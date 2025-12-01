package messages

import (
	"context"
	"frogsmash/internal/app/shared"
	"log"
)

type RedisClient interface {
	InitStream(ctx context.Context) error
	GetMessages(ctx context.Context) ([]*Message, error)
	AcknowledgeMessage(messageID string, ctx context.Context) error
	AddMessage(ctx context.Context, message map[string]interface{}) error
}

type MessageConsumer interface {
	SetUpAndRunWorker(ctx context.Context) error
}

type messageConsumer struct {
	client     RedisClient
	dispatcher Dispatcher
	db         shared.DBWithTxStarter
}

func NewMessageConsumer(ctx context.Context, client RedisClient, dispatcher Dispatcher, db shared.DBWithTxStarter) (MessageConsumer, error) {
	return &messageConsumer{
		client:     client,
		dispatcher: dispatcher,
		db:         db,
	}, nil
}

func (r *messageConsumer) SetUpAndRunWorker(ctx context.Context) error {
	err := r.client.InitStream(ctx)
	if err != nil {
		return err
	}

	log.Println("Redis worker set up")

	for {
		msgs, err := r.client.GetMessages(ctx)

		if err != nil {
			log.Printf("Error reading messages: %v", err)
			continue
		}

		if len(msgs) == 0 {
			log.Println("no new messages, continuing...")
			continue
		}

		for _, msg := range msgs {
			r.processMessage(msg, ctx)
		}
	}
}

func (r *messageConsumer) processMessage(msg *Message, ctx context.Context) {
	messageType, ok := msg.Values["type"].(string)
	if !ok {
		log.Printf("Redis: message ID %s missing type field or type field is not a string", msg.ID)
		r.acknowledgeMessage(msg.ID, ctx)
		return
	}
	r.dispatcher.DispatchMessage(ctx, messageType, msg.Values, r.db)
	r.acknowledgeMessage(msg.ID, ctx)
}

func (r *messageConsumer) acknowledgeMessage(messageID string, ctx context.Context) {
	err := r.client.AcknowledgeMessage(messageID, ctx)
	if err != nil {
		log.Printf("Redis: error acknowledging message ID %s: %v", messageID, err)
	}
}
