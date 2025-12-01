package container

import (
	"context"
	"frogsmash/internal/infrastructure/messages"
)

type WorkerContainer struct {
	*Container
	MessageConsumer messages.MessageConsumer
}

func NewWorkerContainer(c *Container, ctx context.Context) (*WorkerContainer, error) {
	dispatcher := messages.NewDispatcher()
	dispatcher.RegisterHandler("send_verification_email", c.Verification.VerificationService)

	messageConsumer, err := messages.NewMessageConsumer(
		ctx,
		c.InfraServices.RedisClient,
		dispatcher,
		c.InfraServices.DB,
	)
	if err != nil {
		return nil, err
	}

	return &WorkerContainer{
		Container:       c,
		MessageConsumer: messageConsumer,
	}, nil
}
