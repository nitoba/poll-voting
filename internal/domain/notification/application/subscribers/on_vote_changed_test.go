package subscribers_test

import (
	"testing"

	"github.com/nitoba/poll-voting/internal/domain/notification/application/subscribers"
	"github.com/nitoba/poll-voting/internal/domain/notification/application/usecases"
	us "github.com/nitoba/poll-voting/internal/domain/poll/application/usecases"
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
	votersRepository        *repositories_test.InMemoryVotersRepository
	pollsRepository         *repositories_test.InMemoryPollsRepository
	votesRepository         *repositories_test.InMemoryVotesRepository
	countingVotesRepository *repositories_test.InMemoryCountingVotesRepository
	messagePublisher        *messaging_test.InMemoryMessagePublisher
	usecase                 *UseCaseVoteChangedMock
	voteOnPoll              *us.VoteOnPollUseCase
}

func makeCreateOnVoteChangedTestConfig() OnVoteChangedTestConfig {
	pollsRepository := &repositories_test.InMemoryPollsRepository{}
	votersRepository := &repositories_test.InMemoryVotersRepository{}
	countingVotesRepository := &repositories_test.InMemoryCountingVotesRepository{
		Votes: make(map[string]map[string]int),
	}
	votesRepository := &repositories_test.InMemoryVotesRepository{
		CountingRepository: *countingVotesRepository,
	}
	messagePublisher := &messaging_test.InMemoryMessagePublisher{}
	usecase := &UseCaseVoteChangedMock{}

	voteOnPoll := us.NewVoteOnPollUseCase(votesRepository, pollsRepository, votersRepository, countingVotesRepository)

	subscribers.NewOnVoteChangedHandler(usecase, countingVotesRepository)

	return OnVoteChangedTestConfig{
		votesRepository:         votesRepository,
		countingVotesRepository: countingVotesRepository,
		messagePublisher:        messagePublisher,
		usecase:                 usecase,
		votersRepository:        votersRepository,
		pollsRepository:         pollsRepository,
		voteOnPoll:              voteOnPoll,
	}
}

func TestOnVoteChanged(t *testing.T) {
	t.Run("it should be able to send a notification of the update counting votes when changed one", func(t *testing.T) {
		config := makeCreateOnVoteChangedTestConfig()
		config.usecase.On("Execute", mock.Anything).Return(nil)
		voter := factories.MakeVoter()
		poll := factories.MakePool()
		// mask the aggregate vote to dispatch the event
		previousVote := factories.MakeVote(factories.OptionalVoteParams{
			PollId:   &poll.Id,
			OptionId: &poll.Options[0].Id,
			VoterId:  &voter.Id,
		})
		config.votersRepository.Create(voter)
		config.pollsRepository.Create(poll)
		config.votesRepository.Create(previousVote)
		config.countingVotesRepository.IncrementCountVotesByOptionId(poll.Id.String(), poll.Options[0].Id.String())

		assert.Equal(t, 1, config.countingVotesRepository.Votes[poll.Id.String()][poll.Options[0].Id.String()])
		assert.Len(t, config.votesRepository.Votes, 1)

		// when the vote is changed, the domain events must be dispatched
		err := config.voteOnPoll.Execute(&us.VoteOnPollUseCaseRequest{
			PollId:       poll.Id.String(),
			VoterId:      voter.Id.String(),
			PollOptionId: poll.Options[1].Id.String(),
		})

		assert.Nil(t, err)

		config.usecase.AssertCalled(t, "Execute", mock.Anything)
		config.usecase.AssertNumberOfCalls(t, "Execute", 1)
		assert.Equal(t, 0, config.countingVotesRepository.Votes[poll.Id.String()][poll.Options[0].Id.String()])
		assert.Equal(t, 1, config.countingVotesRepository.Votes[poll.Id.String()][poll.Options[1].Id.String()])

	})
}
