package usecases

import (
	"github.com/nitoba/poll-voting/internal/domain/poll/application/repositories"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
)

type FetchPollsByOwnerRequest struct {
	OwnerId string
}

type FetchPollsByOwnerResponse struct {
	Polls []*entities.Poll
}

type FetchPollsByOwnerUseCase struct {
	repo repositories.PollsRepository
}

func (uc *FetchPollsByOwnerUseCase) Execute(req FetchPollsByOwnerRequest) (*FetchPollsByOwnerResponse, error) {
	polls, err := uc.repo.FindManyByOwnerId(req.OwnerId)
	if err != nil {
		return nil, err
	}

	return &FetchPollsByOwnerResponse{
		Polls: polls,
	}, nil
}

func NewFetchPollsByOwnerUseCase(repo repositories.PollsRepository) *FetchPollsByOwnerUseCase {
	return &FetchPollsByOwnerUseCase{
		repo: repo,
	}
}
