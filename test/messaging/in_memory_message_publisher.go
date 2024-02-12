package messaging_test

import (
	"github.com/nitoba/poll-voting/internal/domain/notification/enterprise/entities"
	"github.com/stretchr/testify/mock"
)

type InMemoryMessagePublisher struct {
	mock.Mock
}

func (p *InMemoryMessagePublisher) Publish(message *entities.Notification) error {
	println("Message published with id:", message.Id.String())
	p.Called(message)
	return nil
}
