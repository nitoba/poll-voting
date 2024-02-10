package usecases

import (
	"github.com/nitoba/poll-voting/internal/domain/poll/application/repositories"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
)

type FetchPollsUseCase struct {
	repo repositories.PollsRepository
}

func (uc *FetchPollsUseCase) Execute() ([]*entities.Poll, error) {
	return uc.repo.FindMany()
}

func NewFetchPollsUseCase(repo repositories.PollsRepository) *FetchPollsUseCase {
	return &FetchPollsUseCase{
		repo: repo,
	}
}
