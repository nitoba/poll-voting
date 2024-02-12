package subscribers_test

import (
	"testing"

	"github.com/nitoba/poll-voting/internal/domain/notification/application/subscribers"
	"github.com/nitoba/poll-voting/internal/domain/notification/application/usecases"
	"github.com/nitoba/poll-voting/test/factories"
	messaging_test "github.com/nitoba/poll-voting/test/messaging"
	repositories_test "github.com/nitoba/poll-voting/test/repositories"
	"github.com/stretchr/testify/mock"
)

type UseCaseMock struct {
	mock.Mock
}

func (u *UseCaseMock) Execute(req *usecases.UpdateVotingCountUseCaseRequest) error {
	u.Called(req)
	return nil
}

type OnVoteCreatedTestConfig struct {
	votesRepository         *repositories_test.InMemoryVotesRepository
	countingVotesRepository *repositories_test.InMemoryCountingVotesRepository
	messagePublisher        *messaging_test.InMemoryMessagePublisher
	usecase                 *UseCaseMock
}

func makeCreateOnVoteCreatedTestConfig() OnVoteCreatedTestConfig {
	votesRepository := &repositories_test.InMemoryVotesRepository{}
	countingVotesRepository := &repositories_test.InMemoryCountingVotesRepository{}
	messagePublisher := &messaging_test.InMemoryMessagePublisher{}
	usecase := &UseCaseMock{}

	subscribers.NewOnVoteCreatedHandler(usecase, countingVotesRepository)

	return OnVoteCreatedTestConfig{
		votesRepository:         votesRepository,
		countingVotesRepository: countingVotesRepository,
		messagePublisher:        messagePublisher,
		usecase:                 usecase,
	}
}

func TestOnVoteCreated(t *testing.T) {
	t.Run("it should be able to send a notification of the update counting votes", func(t *testing.T) {
		config := makeCreateOnVoteCreatedTestConfig()
		config.usecase.On("Execute", mock.Anything).Return(nil)

		poll := factories.MakePool()
		// mask the aggregate vote to dispatch the event
		vote := factories.MakeVote(factories.OptionalVoteParams{
			PollId:   &poll.Id,
			OptionId: &poll.Options[0].Id,
		})
		// when the vote is created, the domain events must be dispatched
		config.votesRepository.Create(vote)

		config.usecase.AssertCalled(t, "Execute", mock.Anything)
		config.usecase.AssertNumberOfCalls(t, "Execute", 1)
	})
}
