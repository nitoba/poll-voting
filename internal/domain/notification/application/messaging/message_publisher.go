package messaging

import "github.com/nitoba/poll-voting/internal/domain/notification/enterprise/entities"

type MessagePublisher interface {
	Publish(message *entities.Notification) error
}
