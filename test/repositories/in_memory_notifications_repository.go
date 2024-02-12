package repositories_test

import "github.com/nitoba/poll-voting/internal/domain/notification/enterprise/entities"

type InMemoryNotificationsRepository struct {
	Notifications []*entities.Notification
}

func (repo *InMemoryNotificationsRepository) Create(notification *entities.Notification) error {
	repo.Notifications = append(repo.Notifications, notification)
	return nil
}
