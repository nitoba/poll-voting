package usecases

import (
	"github.com/nitoba/poll-voting/internal/domain/poll/application/repositories"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases/errors"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
)

type GetPollByIdUseCase struct {
	repo              repositories.PollsRepository
	countingVotesRepo repositories.CountingVotesRepository
}

func (uc *GetPollByIdUseCase) Execute(id string) (*entities.Poll, error) {
	poll, err := uc.repo.FindById(id)
	if err != nil || poll == nil {
		return nil, errors.ErrPollNotFound
	}

	votes, err := uc.countingVotesRepo.CountVotes(poll.Id.String())

	if err != nil {
		return nil, err
	}

	poll.Votes = votes

	return poll, nil
}

func NewGetPollByIdUseCase(repo repositories.PollsRepository, countingVotesRepo repositories.CountingVotesRepository) *GetPollByIdUseCase {
	return &GetPollByIdUseCase{
		repo:              repo,
		countingVotesRepo: countingVotesRepo,
	}
}
