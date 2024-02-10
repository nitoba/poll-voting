package repositories

import "github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"

type VotersRepository interface {
	Create(voter *entities.Voter) error
	FindById(id string) (*entities.Voter, error)
	FindByEmail(email string) (*entities.Voter, error)
}
