package repositories

import "github.com/nitoba/poll-voting/internal/domain/notification/enterprise/entities"

type NotificationsRepository interface {
	Create(notification *entities.Notification) error
	// FindByPollId(pollId string) ([]*entities.Notification, error)
}
