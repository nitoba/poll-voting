package repositories_test

import "github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"

type InMemoryVotesRepository struct {
	Votes []*entities.Vote
}

func (repo *InMemoryVotesRepository) Create(vote *entities.Vote) error {
	repo.Votes = append(repo.Votes, vote)
	return nil
}
