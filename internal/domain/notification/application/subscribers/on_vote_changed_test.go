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

type UseCaseVoteChangedMock struct {
	mock.Mock
}

func (u *UseCaseVoteChangedMock) Execute(req *usecases.UpdateVotingCountUseCaseRequest) error {
	u.Called(req)
	return nil
}

type OnVoteChangedTestConfig struct {
	votesRepository         *repositories_test.InMemoryVotesRepository
	countingVotesRepository *repositories_test.InMemoryCountingVotesRepository
	messagePublisher        *messaging_test.InMemoryMessagePublisher
	usecase                 *UseCaseVoteChangedMock
}

func makeCreateOnVoteChangedTestConfig() OnVoteChangedTestConfig {
	countingVotesRepository := &repositories_test.InMemoryCountingVotesRepository{
		Votes: make(map[string]map[string]int),
	}
	votesRepository := &repositories_test.InMemoryVotesRepository{
		CountingRepository: *countingVotesRepository,
	}
	messagePublisher := &messaging_test.InMemoryMessagePublisher{}
	usecase := &UseCaseVoteChangedMock{}

	subscribers.NewOnVoteChangedHandler(usecase, countingVotesRepository)
	subscribers.NewOnVoteCreatedHandler(usecase, countingVotesRepository)

	return OnVoteChangedTestConfig{
		votesRepository:         votesRepository,
		countingVotesRepository: countingVotesRepository,
		messagePublisher:        messagePublisher,
		usecase:                 usecase,
	}
}

func TestOnVoteChanged(t *testing.T) {
	t.Run("it should be able to send a notification of the update counting votes when changed one", func(t *testing.T) {
		config := makeCreateOnVoteChangedTestConfig()
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

		assert.Equal(t, 1, config.countingVotesRepository.Votes[poll.Id.String()][poll.Options[0].Id.String()])

		// when the vote is changed, the domain events must be dispatched
		config.votesRepository.Delete(vote)

		vote.ChangeVoteOption(poll.Options[1].Id.String())

		config.votesRepository.Create(vote)

		config.usecase.AssertCalled(t, "Execute", mock.Anything)
		config.usecase.AssertNumberOfCalls(t, "Execute", 2)
		assert.Equal(t, 0, config.countingVotesRepository.Votes[poll.Id.String()][poll.Options[0].Id.String()])
		assert.Equal(t, 1, config.countingVotesRepository.Votes[poll.Id.String()][poll.Options[1].Id.String()])

	})
}
