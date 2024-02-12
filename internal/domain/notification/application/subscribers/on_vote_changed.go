package subscribers

import (
	"github.com/nitoba/poll-voting/internal/domain/core"
	"github.com/nitoba/poll-voting/internal/domain/notification/application/usecases"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/repositories"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/events"
)

type OnVoteChangedHandler struct {
	updateVotingCountUseCase usecases.UpdateVotingCountUseCaseInterface
	countingVotesRepository  repositories.CountingVotesRepository
}

func (h *OnVoteChangedHandler) SetupSubscriptions() {
	core.DomainEvents().Register(h.sendVoteCreatedNotification, "event.vote.changed")
}

func (h *OnVoteChangedHandler) sendVoteCreatedNotification(event interface{}) {
	voteChangedEvent := event.(*events.VoteChangedEvent)
	countOfVotes, err := h.countingVotesRepository.CountVotesByOptionId(voteChangedEvent.PollId.String(), voteChangedEvent.OptionId.String())

	if err != nil {
		return
	}

	err = h.updateVotingCountUseCase.Execute(&usecases.UpdateVotingCountUseCaseRequest{
		PollId:       voteChangedEvent.PollId.String(),
		PollOptionId: voteChangedEvent.OptionId.String(),
		CountOfVotes: countOfVotes,
	})

	if err != nil {
		println("Error to send notification: ", err.Error())
	}
}

func NewOnVoteChangedHandler(updateVotingCountUseCase usecases.UpdateVotingCountUseCaseInterface, countingVotesRepository repositories.CountingVotesRepository) {
	handler := &OnVoteChangedHandler{
		updateVotingCountUseCase: updateVotingCountUseCase,
		countingVotesRepository:  countingVotesRepository,
	}
	handler.SetupSubscriptions()
}
