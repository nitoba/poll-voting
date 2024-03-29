package repositories

import "github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"

type VotesRepository interface {
	Create(vote *entities.Vote) error
	Delete(vote *entities.Vote) error
	FindByPollIdAndVoterId(pollId string, voterId string) (*entities.Vote, error)
}
