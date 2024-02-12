package usecases_test

import (
	"testing"

	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases"
	"github.com/nitoba/poll-voting/test/factories"
	repositories_test "github.com/nitoba/poll-voting/test/repositories"
	"github.com/stretchr/testify/assert"
)

type VoteOnPollUseCaseConfig struct {
	countingRepo *repositories_test.InMemoryCountingVotesRepository
	pollRepo     *repositories_test.InMemoryPollsRepository
	voteRepo     *repositories_test.InMemoryVotesRepository
	voterRepo    *repositories_test.InMemoryVotersRepository
	sut          *usecases.VoteOnPollUseCase
}

func makeVoteOnPollUseCase() VoteOnPollUseCaseConfig {
	countingRepo := &repositories_test.InMemoryCountingVotesRepository{}
	pollRepo := &repositories_test.InMemoryPollsRepository{}
	voteRepo := &repositories_test.InMemoryVotesRepository{
		CountingRepository: *countingRepo,
	}
	voterRepo := &repositories_test.InMemoryVotersRepository{}
	sut := usecases.NewVoteOnPollUseCase(voteRepo, pollRepo, voterRepo)

	return VoteOnPollUseCaseConfig{
		pollRepo:  pollRepo,
		voteRepo:  voteRepo,
		voterRepo: voterRepo,
		sut:       sut,
	}
}

func TestVoteOnPollUseCase(t *testing.T) {
	t.Run("it should vote on a poll", func(t *testing.T) {
		voter := factories.MakeVoter()
		poll := factories.MakePool()
		res := makeVoteOnPollUseCase()

		res.voterRepo.Create(voter)
		res.pollRepo.Create(poll)

		err := res.sut.Execute(&usecases.VoteOnPollUseCaseRequest{
			PollId:       poll.Id.String(),
			VoterId:      voter.Id.String(),
			PollOptionId: poll.Options[0].Id.String(),
		})

		assert.Nil(t, err)
		assert.NotEmpty(t, res.voteRepo.Votes[0].Id.String())
		assert.Equal(t, res.voteRepo.Votes[0].PollId.String(), poll.Id.String())
		assert.Equal(t, res.voteRepo.Votes[0].VoterId.String(), voter.Id.String())
		assert.Equal(t, res.voteRepo.Votes[0].OptionId.String(), poll.Options[0].Id.String())
	})

}
func TestVoteOnPollAdditionalCases(t *testing.T) {
	t.Run("returns error if poll does not exist", func(t *testing.T) {
		res := makeVoteOnPollUseCase()
		// Arrange
		req := &usecases.VoteOnPollUseCaseRequest{
			PollId: "invalid-poll-id",
		}

		// Act
		err := res.sut.Execute(req)

		// Assert
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "poll not found")
	})

	t.Run("returns error if voter does not exist", func(t *testing.T) {
		res := makeVoteOnPollUseCase()
		poll := factories.MakePool()
		res.pollRepo.Create(poll)
		// Arrange
		req := &usecases.VoteOnPollUseCaseRequest{
			VoterId:      "invalid-voter-id",
			PollId:       poll.Id.String(),
			PollOptionId: poll.Options[0].Id.String(),
		}

		// Act
		err := res.sut.Execute(req)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "voter not found")
	})

	t.Run("returns error if poll option does not exist", func(t *testing.T) {
		// Arrange
		res := makeVoteOnPollUseCase()
		poll := factories.MakePool()
		voter := factories.MakeVoter()
		res.pollRepo.Create(poll)
		res.voterRepo.Create(voter)

		req := &usecases.VoteOnPollUseCaseRequest{
			PollOptionId: "invalid-option-id",
			PollId:       poll.Id.String(),
			VoterId:      voter.Id.String(),
		}

		// Act
		err := res.sut.Execute(req)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid poll option")
	})

	t.Run("returns error if voter has already voted on this poll option", func(t *testing.T) {
		// Arrange
		res := makeVoteOnPollUseCase()
		poll := factories.MakePool()
		voter := factories.MakeVoter()
		res.pollRepo.Create(poll)
		res.voterRepo.Create(voter)

		vote := factories.MakeVote(factories.OptionalVoteParams{
			PollId:   &poll.Id,
			VoterId:  &voter.Id,
			OptionId: &poll.Options[0].Id,
		})
		res.voteRepo.Create(vote)

		req := usecases.VoteOnPollUseCaseRequest{
			PollOptionId: poll.Options[0].Id.String(),
			PollId:       poll.Id.String(),
			VoterId:      voter.Id.String(),
		}

		// Act
		err := res.sut.Execute(&req)

		// Assert
		assert.Error(t, err, "voter has already voted on this poll option")
	})

	t.Run("it should be change the vote", func(t *testing.T) {
		// Arrange
		res := makeVoteOnPollUseCase()
		poll := factories.MakePool()
		voter := factories.MakeVoter()
		res.pollRepo.Create(poll)
		res.voterRepo.Create(voter)

		vote := factories.MakeVote(factories.OptionalVoteParams{
			PollId:   &poll.Id,
			VoterId:  &voter.Id,
			OptionId: &poll.Options[0].Id,
		})

		res.voteRepo.Create(vote)

		req := usecases.VoteOnPollUseCaseRequest{
			PollOptionId: poll.Options[1].Id.String(),
			PollId:       poll.Id.String(),
			VoterId:      voter.Id.String(),
		}

		// Act

		err := res.sut.Execute(&req)

		// Assert

		assert.Nil(t, err)
		assert.NotEmpty(t, res.voteRepo.Votes[0].Id.String())
		assert.Equal(t, res.voteRepo.Votes[0].PollId.String(), poll.Id.String())
		assert.Equal(t, res.voteRepo.Votes[0].VoterId.String(), voter.Id.String())
		assert.Equal(t, res.voteRepo.Votes[0].OptionId.String(), poll.Options[1].Id.String())
	})

}
