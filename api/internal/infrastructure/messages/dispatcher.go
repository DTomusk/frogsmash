package messages

import (
	"context"
	"fmt"
	"frogsmash/internal/app/shared"
)

type MessageHandler interface {
	HandleMessage(ctx context.Context, values map[string]interface{}, db shared.DBWithTxStarter) error
}

type Dispatcher interface {
	RegisterHandler(messageType string, handler MessageHandler)
	DispatchMessage(ctx context.Context, messageType string, values map[string]interface{}, db shared.DBWithTxStarter) error
}

type dispatcher struct {
	handlers map[string]MessageHandler
}

func NewDispatcher() Dispatcher {
	return &dispatcher{
		handlers: make(map[string]MessageHandler),
	}
}

func (d *dispatcher) RegisterHandler(messageType string, handler MessageHandler) {
	d.handlers[messageType] = handler
}

func (d *dispatcher) DispatchMessage(ctx context.Context, messageType string, values map[string]interface{}, db shared.DBWithTxStarter) error {
	handler, exists := d.handlers[messageType]
	if !exists {
		return fmt.Errorf("no handler registered for message type: %s", messageType)
	}
	return handler.HandleMessage(ctx, values, db)
}
