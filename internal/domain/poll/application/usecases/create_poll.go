package usecases

import (
	"github.com/nitoba/poll-voting/internal/domain/poll/application/repositories"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases/errors"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
)

type CreatePollUseCase struct {
	repo repositories.PollsRepository
}

type CreatePollRequest struct {
	Title   string
	Options []string
}

func (uc *CreatePollUseCase) Execute(req CreatePollRequest) error {
	if len(req.Options) == 0 || len(req.Options) < 2 {
		return errors.ErrInvalidPoll
	}

	var pollOptions []*entities.PollOption

	for _, option := range req.Options {
		currentOption, err := entities.NewPollOption(option)
		if err != nil {
			return errors.ErrInvalidOption
		}
		pollOptions = append(pollOptions, currentOption)
	}

	poll, err := entities.NewPoll(req.Title, pollOptions)

	if err != nil {
		return errors.ErrInvalidPoll
	}

	if err = uc.repo.Create(poll); err != nil {
		return err
	}

	return nil
}

func NewCreatePollUseCase(repo repositories.PollsRepository) *CreatePollUseCase {
	return &CreatePollUseCase{
		repo: repo,
	}
}
