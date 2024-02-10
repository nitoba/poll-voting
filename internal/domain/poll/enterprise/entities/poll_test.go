package entities_test

import (
	"testing"
	"time"

	"github.com/nitoba/poll-voting/internal/domain/core"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
	"github.com/stretchr/testify/assert"
)

func TestNewPoll(t *testing.T) {
	t.Run("with required params", func(t *testing.T) {
		options := []*entities.PollOption{
			{Title: "Option 1"},
			{Title: "Option 2"},
		}

		poll, err := entities.NewPoll("Test Poll", options)

		assert.NoError(t, err)
		assert.NotEmpty(t, poll.Id)
		assert.Equal(t, "Test Poll", poll.Title)
		assert.Equal(t, options, poll.Options)
		assert.WithinDuration(t, time.Now(), poll.CreatedAt, time.Second)
	})

	t.Run("with optional params", func(t *testing.T) {
		id := core.NewUniqueEntityId()
		createdAt := time.Now().Add(-1 * time.Hour)

		options := []*entities.PollOption{
			{Title: "Option 1"},
			{Title: "Option 2"},
		}

		optional := entities.OptionalParams{
			Id:        &id,
			CreatedAt: &createdAt,
		}

		poll, err := entities.NewPoll("Test Poll", options, optional)

		assert.NoError(t, err)
		assert.Equal(t, id, poll.Id)
		assert.Equal(t, "Test Poll", poll.Title)
		assert.Equal(t, options, poll.Options)
		assert.NotNil(t, poll.CreatedAt)
		assert.Equal(t, poll.CreatedAt, createdAt)
	})
}
func TestPoll_AddOption(t *testing.T) {
	t.Run("add first option", func(t *testing.T) {
		poll := entities.Poll{}

		option, _ := entities.NewPollOption("Option 1")
		poll.AddOption(option)

		assert.Equal(t, 1, len(poll.Options))
		assert.Equal(t, option, poll.Options[0])
	})

	t.Run("add duplicate option", func(t *testing.T) {
		existingOption := &entities.PollOption{Title: "Option 1"}
		poll := entities.Poll{
			Options: []*entities.PollOption{existingOption},
		}

		duplicateOption := &entities.PollOption{Title: "Option 1"}
		poll.AddOption(duplicateOption)

		assert.Equal(t, 1, len(poll.Options))
		assert.Equal(t, existingOption, poll.Options[0])
	})

	t.Run("add unique option", func(t *testing.T) {
		existingOption := &entities.PollOption{Title: "Option 1"}
		poll := entities.Poll{
			Options: []*entities.PollOption{existingOption},
		}

		newOption := &entities.PollOption{Title: "Option 2"}
		poll.AddOption(newOption)

		assert.Equal(t, 2, len(poll.Options))
		assert.Equal(t, existingOption, poll.Options[0])
		assert.Equal(t, newOption, poll.Options[1])
	})
}

func TestPoll_Equals(t *testing.T) {
	t.Run("same object is equal", func(t *testing.T) {
		poll := entities.Poll{}
		assert.True(t, poll.Equals(&poll))
	})

	t.Run("different objects with same ID are equal", func(t *testing.T) {
		id := core.NewUniqueEntityId()
		poll1 := entities.Poll{Id: id}
		poll2 := entities.Poll{Id: id}
		assert.True(t, poll1.Equals(&poll2))
	})

	t.Run("different objects with different IDs are not equal", func(t *testing.T) {
		poll1 := entities.Poll{Id: core.NewUniqueEntityId()}
		poll2 := entities.Poll{Id: core.NewUniqueEntityId()}
		assert.False(t, poll1.Equals(&poll2))
	})
}
