package usecases

import (
	"github.com/nitoba/poll-voting/internal/domain/notification/application/messaging"
	"github.com/nitoba/poll-voting/internal/domain/notification/enterprise/entities"
	"github.com/nitoba/poll-voting/internal/domain/notification/enterprise/value_objects"
)

type UpdateVotingCountUseCaseInterface interface {
	Execute(req *UpdateVotingCountUseCaseRequest) error
}

type UpdateVotingCountUseCase struct {
	messagePublisher messaging.MessagePublisher
}

type UpdateVotingCountUseCaseRequest struct {
	PollId       string
	PollOptionId string
	CountOfVotes int
}

func (u *UpdateVotingCountUseCase) Execute(req *UpdateVotingCountUseCaseRequest) error {

	newVote := value_objects.CreateNewVote(req.PollOptionId, req.CountOfVotes)
	notification := entities.NewNotification("New voting count", newVote)

	if err := u.messagePublisher.Publish(notification); err != nil {
		return err
	}
	return nil
}

func NewUpdateVotingCountUseCase(messagePublisher messaging.MessagePublisher) *UpdateVotingCountUseCase {
	return &UpdateVotingCountUseCase{
		messagePublisher: messagePublisher,
	}
}
