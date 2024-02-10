package repositories

import "github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"

type ParticipantsRepository interface {
	Create(participant *entities.Participant) error
	FindById(id string) (*entities.Participant, error)
	FindByEmail(email string) (*entities.Participant, error)
}
