package usecases

import (
	"github.com/nitoba/poll-voting/internal/domain/poll/application/repositories"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases/errors"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
)

type CreatePollUseCase struct {
	repo      repositories.PollsRepository
	voterRepo repositories.VotersRepository
}

type CreatePollRequest struct {
	Title   string
	Options []string
	OwnerId string
}

func (uc *CreatePollUseCase) Execute(req CreatePollRequest) error {
	p, err := uc.voterRepo.FindById(req.OwnerId)
	if p == nil && err != nil {
		return errors.ErrInvalidOwner
	}

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

	poll, err := entities.NewPoll(req.Title, pollOptions, p.Id)

	if err != nil {
		return errors.ErrInvalidPoll
	}

	if err = uc.repo.Create(poll); err != nil {
		return err
	}

	return nil
}

func NewCreatePollUseCase(repo repositories.PollsRepository, voterRepo repositories.VotersRepository) *CreatePollUseCase {
	return &CreatePollUseCase{
		repo:      repo,
		voterRepo: voterRepo,
	}
}
