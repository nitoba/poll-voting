package usecases

import (
	"github.com/nitoba/poll-voting/internal/domain/poll/application/repositories"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases/errors"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
)

type GetPollByIdUseCase struct {
	repo repositories.PollsRepository
}

func (uc *GetPollByIdUseCase) Execute(id string) (*entities.Poll, error) {
	poll, err := uc.repo.FindById(id)
	if err != nil || poll == nil {
		return nil, errors.ErrPollNotFound
	}
	return poll, nil
}

func NewGetPollByIdUseCase(repo repositories.PollsRepository) *GetPollByIdUseCase {
	return &GetPollByIdUseCase{
		repo: repo,
	}
}
