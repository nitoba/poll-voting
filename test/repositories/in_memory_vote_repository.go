package repositories_test

import (
	"errors"

	"github.com/nitoba/poll-voting/internal/domain/core"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
)

type InMemoryVotesRepository struct {
	Votes []*entities.Vote
}

func (repo *InMemoryVotesRepository) Create(vote *entities.Vote) error {
	repo.Votes = append(repo.Votes, vote)
	// TODO: before dispatch events, store the count votes in redis
	core.DomainEvents().DispatchEventsForAggregate(vote.Id)
	return nil
}

func (repo *InMemoryVotesRepository) FindByOptionId(id string) (*entities.Vote, error) {
	for _, p := range repo.Votes {
		if p.OptionId.String() == id {
			return p, nil
		}
	}
	return nil, errors.New("vote not found")
}

func (repo *InMemoryVotesRepository) Delete(vote *entities.Vote) error {
	for i, v := range repo.Votes {
		if v.OptionId.String() == vote.OptionId.String() {
			repo.Votes = append(repo.Votes[:i], repo.Votes[i+1:]...)
			return nil
		}
	}
	return errors.New("vote not found")
}

func (repo *InMemoryVotesRepository) FindByPollIdAndVoterId(pollId string, voterId string) (*entities.Vote, error) {
	for _, p := range repo.Votes {
		if p.PollId.String() == pollId && p.VoterId.String() == voterId {
			return p, nil
		}
	}
	return nil, errors.New("vote not found")
}
