package subscribers_test

import (
	"testing"

	"github.com/nitoba/poll-voting/internal/domain/notification/application/subscribers"
	"github.com/nitoba/poll-voting/internal/domain/notification/application/usecases"
	"github.com/nitoba/poll-voting/test/factories"
	messaging_test "github.com/nitoba/poll-voting/test/messaging"
	repositories_test "github.com/nitoba/poll-voting/test/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UseCaseVoteCreatedMock struct {
	mock.Mock
}

func (u *UseCaseVoteCreatedMock) Execute(req *usecases.UpdateVotingCountUseCaseRequest) error {
	u.Called(req)
	return nil
}

type OnVoteCreatedTestConfig struct {
	votesRepository         *repositories_test.InMemoryVotesRepository
	countingVotesRepository *repositories_test.InMemoryCountingVotesRepository
	messagePublisher        *messaging_test.InMemoryMessagePublisher
	usecase                 *UseCaseVoteCreatedMock
}

func makeCreateOnVoteCreatedTestConfig() OnVoteCreatedTestConfig {
	countingVotesRepository := &repositories_test.InMemoryCountingVotesRepository{
		Votes: make(map[string]map[string]int),
	}
	votesRepository := &repositories_test.InMemoryVotesRepository{
		CountingRepository: *countingVotesRepository,
	}
	messagePublisher := &messaging_test.InMemoryMessagePublisher{}
	usecase := &UseCaseVoteCreatedMock{}

	subscribers.NewOnVoteCreatedHandler(usecase, countingVotesRepository)

	return OnVoteCreatedTestConfig{
		votesRepository:         votesRepository,
		countingVotesRepository: countingVotesRepository,
		messagePublisher:        messagePublisher,
		usecase:                 usecase,
	}
}

func TestOnVoteCreated(t *testing.T) {
	t.Run("it should be able to send a notification of the update counting votes when created one", func(t *testing.T) {
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
		assert.Equal(t, 1, config.countingVotesRepository.Votes[poll.Id.String()][poll.Options[0].Id.String()])
	})
}