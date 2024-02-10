package entities_test

import (
	"testing"
	"time"

	"github.com/nitoba/poll-voting/internal/domain/core"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
	"github.com/stretchr/testify/assert"
)

func TestVote_Equals(t *testing.T) {
	t.Run("same object is equal", func(t *testing.T) {
		vote := entities.Vote{}
		assert.True(t, vote.Equals(&vote))
	})

	t.Run("different objects with same ID are equal", func(t *testing.T) {
		id := core.NewUniqueEntityId()
		vote1 := entities.Vote{Id: id}
		vote2 := entities.Vote{Id: id}
		assert.True(t, vote1.Equals(&vote2))
	})

	t.Run("different objects with different IDs are not equal", func(t *testing.T) {
		vote1 := entities.Vote{Id: core.NewUniqueEntityId()}
		vote2 := entities.Vote{Id: core.NewUniqueEntityId()}
		assert.False(t, vote1.Equals(&vote2))
	})

	t.Run("objects with different poll IDs are not equal", func(t *testing.T) {
		vote1 := entities.Vote{PollId: core.NewUniqueEntityId()}
		vote2 := entities.Vote{PollId: core.NewUniqueEntityId()}
		assert.False(t, vote1.Equals(&vote2))
	})

	t.Run("objects with different option IDs are not equal", func(t *testing.T) {
		vote1 := entities.Vote{OptionId: core.NewUniqueEntityId()}
		vote2 := entities.Vote{OptionId: core.NewUniqueEntityId()}
		assert.False(t, vote1.Equals(&vote2))
	})

	t.Run("objects with different voter IDs are not equal", func(t *testing.T) {
		vote1 := entities.Vote{VoterId: core.NewUniqueEntityId()}
		vote2 := entities.Vote{VoterId: core.NewUniqueEntityId()}
		assert.False(t, vote1.Equals(&vote2))
	})

	t.Run("nil objects are not equal", func(t *testing.T) {
		var vote1 *entities.Vote
		var vote2 *entities.Vote
		assert.False(t, vote1.Equals(vote2))
	})
}
func TestVote_NewVote(t *testing.T) {
	t.Run("with no optional params", func(t *testing.T) {
		pollId := core.NewUniqueEntityId()
		optionId := core.NewUniqueEntityId()
		voterId := core.NewUniqueEntityId()

		vote, err := entities.NewVote(pollId, optionId, voterId)

		assert.NoError(t, err)
		assert.NotEmpty(t, vote.Id)
		assert.Equal(t, pollId, vote.PollId)
		assert.Equal(t, optionId, vote.OptionId)
		assert.Equal(t, voterId, vote.VoterId)
		assert.WithinDuration(t, time.Now(), vote.CreatedAt, time.Second)
	})

	t.Run("with custom id", func(t *testing.T) {
		id := core.NewUniqueEntityId()
		pollId := core.NewUniqueEntityId()
		optionId := core.NewUniqueEntityId()
		voterId := core.NewUniqueEntityId()

		vote, err := entities.NewVote(pollId, optionId, voterId, entities.OptionalVoteParams{
			Id: &id,
		})

		assert.NoError(t, err)
		assert.Equal(t, id, vote.Id)
	})

	t.Run("with custom poll id", func(t *testing.T) {
		pollId := core.NewUniqueEntityId()
		customPollId := core.NewUniqueEntityId()
		optionId := core.NewUniqueEntityId()
		voterId := core.NewUniqueEntityId()

		vote, err := entities.NewVote(pollId, optionId, voterId, entities.OptionalVoteParams{
			PollId: &customPollId,
		})

		assert.NoError(t, err)
		assert.Equal(t, customPollId, vote.PollId)
	})

	t.Run("with custom option id", func(t *testing.T) {
		pollId := core.NewUniqueEntityId()
		optionId := core.NewUniqueEntityId()
		customOptionId := core.NewUniqueEntityId()
		voterId := core.NewUniqueEntityId()

		vote, err := entities.NewVote(pollId, optionId, voterId, entities.OptionalVoteParams{
			OptionId: &customOptionId,
		})

		assert.NoError(t, err)
		assert.Equal(t, customOptionId, vote.OptionId)
	})

	t.Run("with custom voter id", func(t *testing.T) {
		pollId := core.NewUniqueEntityId()
		optionId := core.NewUniqueEntityId()
		voterId := core.NewUniqueEntityId()
		customVoterId := core.NewUniqueEntityId()

		vote, err := entities.NewVote(pollId, optionId, voterId, entities.OptionalVoteParams{
			VoterId: &customVoterId,
		})

		assert.NoError(t, err)
		assert.Equal(t, customVoterId, vote.VoterId)
	})

	t.Run("with custom created at", func(t *testing.T) {
		pollId := core.NewUniqueEntityId()
		optionId := core.NewUniqueEntityId()
		voterId := core.NewUniqueEntityId()
		createdAt := time.Now().Add(-1 * time.Hour)

		vote, err := entities.NewVote(pollId, optionId, voterId, entities.OptionalVoteParams{
			CreatedAt: &createdAt,
		})

		assert.NoError(t, err)
		assert.Equal(t, createdAt, vote.CreatedAt)
	})

}
