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

type MessageService interface {
	SetUpAndRunWorker(ctx context.Context) error
	EnqueueMessage(ctx context.Context, message map[string]interface{}) error
}

type messageService struct {
	client     RedisClient
	dispatcher Dispatcher
	db         shared.DBWithTxStarter
}

func NewMessageService(ctx context.Context, client RedisClient, dispatcher Dispatcher, db shared.DBWithTxStarter) (MessageService, error) {
	return &messageService{
		client:     client,
		dispatcher: dispatcher,
		db:         db,
	}, nil
}

func (r *messageService) SetUpAndRunWorker(ctx context.Context) error {
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

func (r *messageService) processMessage(msg *Message, ctx context.Context) {
	messageType, ok := msg.Values["type"].(string)
	if !ok {
		log.Printf("Redis: message ID %s missing type field or type field is not a string", msg.ID)
		r.acknowledgeMessage(msg.ID, ctx)
		return
	}
	r.dispatcher.DispatchMessage(ctx, messageType, msg.Values, r.db)
	r.acknowledgeMessage(msg.ID, ctx)
}

func (r *messageService) acknowledgeMessage(messageID string, ctx context.Context) {
	err := r.client.AcknowledgeMessage(messageID, ctx)
	if err != nil {
		log.Printf("Redis: error acknowledging message ID %s: %v", messageID, err)
	}
}

func (r *messageService) EnqueueMessage(ctx context.Context, message map[string]interface{}) error {
	log.Printf("Enqueuing message %v", message)
	return r.client.AddMessage(ctx, message)
}

// TODO move model
type Message struct {
	ID     string
	Values map[string]interface{}
}
