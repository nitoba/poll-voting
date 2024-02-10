package repositories

import "github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"

type PollsRepository interface {
	Create(poll *entities.Poll) error
	FindById(id string) (*entities.Poll, error)
	FindMany() ([]*entities.Poll, error)
}
