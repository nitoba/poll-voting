package subscribers

import (
	"sync"

	"github.com/nitoba/poll-voting/internal/domain/core"
	"github.com/nitoba/poll-voting/internal/domain/notification/application/usecases"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/repositories"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/events"
)

type OnVoteCreatedHandler struct {
	updateVotingCountUseCase usecases.UpdateVotingCountUseCaseInterface
	countingVotesRepository  repositories.CountingVotesRepository
}

func (h *OnVoteCreatedHandler) SetupSubscriptions() {
	core.DomainEvents().Register(h.sendVoteCreatedNotification, "event.vote.created")
}

func (h *OnVoteCreatedHandler) sendVoteCreatedNotification(event interface{}, wg *sync.WaitGroup) {
	voteCreatedEvent := event.(*events.VoteCreatedEvent)
	countOfVotes, err := h.countingVotesRepository.CountVotesByOptionId(voteCreatedEvent.PollId.String(), voteCreatedEvent.OptionId.String())

	if err != nil {
		wg.Done()
		return
	}

	err = h.updateVotingCountUseCase.Execute(&usecases.UpdateVotingCountUseCaseRequest{
		PollId:       voteCreatedEvent.PollId.String(),
		PollOptionId: voteCreatedEvent.OptionId.String(),
		CountOfVotes: countOfVotes,
	})

	if err != nil {
		println("Error to send notification: ", err.Error())
	}

	wg.Done()

}

func NewOnVoteCreatedHandler(updateVotingCountUseCase usecases.UpdateVotingCountUseCaseInterface, countingVotesRepository repositories.CountingVotesRepository) {
	handler := &OnVoteCreatedHandler{
		updateVotingCountUseCase: updateVotingCountUseCase,
		countingVotesRepository:  countingVotesRepository,
	}
	handler.SetupSubscriptions()
}
