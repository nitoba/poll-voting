package usecases

import (
	"github.com/nitoba/poll-voting/internal/domain/notification/application/repositories"
	"github.com/nitoba/poll-voting/internal/domain/notification/enterprise/entities"
	"github.com/nitoba/poll-voting/internal/domain/notification/enterprise/value_objects"
)

type UpdateVotingCountUseCase struct {
	repo repositories.NotificationsRepository
}

type UpdateVotingCountUseCaseRequest struct {
	PollId       string
	PollOptionId string
	CountOfVotes int
}

func (u *UpdateVotingCountUseCase) Execute(req *UpdateVotingCountUseCaseRequest) error {

	newVote := value_objects.CreateNewVote(req.PollOptionId, req.CountOfVotes)
	notification := entities.NewNotification("New voting count", newVote)

	if err := u.repo.Create(notification); err != nil {
		return err
	}
	return nil
}

func NewUpdateVotingCountUseCase(repo repositories.NotificationsRepository) *UpdateVotingCountUseCase {
	return &UpdateVotingCountUseCase{
		repo: repo,
	}
}
