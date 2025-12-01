package container

import (
	"context"
	"frogsmash/internal/infrastructure/messages"
)

type WorkerContainer struct {
	*BaseContainer
	MessageConsumer messages.MessageConsumer
}

func NewWorkerContainer(c *BaseContainer, ctx context.Context) (*WorkerContainer, error) {
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
		BaseContainer:   c,
		MessageConsumer: messageConsumer,
	}, nil
}
