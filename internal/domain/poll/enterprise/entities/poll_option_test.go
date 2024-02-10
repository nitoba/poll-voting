package entities_test

import (
	"testing"

	"github.com/nitoba/poll-voting/internal/domain/core"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
	"github.com/stretchr/testify/assert"
)

func TestNewPollOption(t *testing.T) {
	t.Run("with title", func(t *testing.T) {
		title := "Option 1"
		option, err := entities.NewPollOption(title)

		assert.NoError(t, err)
		assert.NotEmpty(t, option.Id)
		assert.Equal(t, title, option.Title)
	})

	t.Run("with title and id", func(t *testing.T) {
		id := core.NewUniqueEntityId()
		title := "Option 1"

		option, err := entities.NewPollOption(title, id)

		assert.NoError(t, err)
		assert.Equal(t, id, option.Id)
		assert.Equal(t, title, option.Title)
	})
}

func TestPollOption_Equals(t *testing.T) {
	t.Run("same object is equal", func(t *testing.T) {
		option := entities.PollOption{}
		assert.True(t, option.Equals(&option))
	})

	t.Run("different objects with same ID are equal", func(t *testing.T) {
		id := core.NewUniqueEntityId()
		option1 := entities.PollOption{Id: id}
		option2 := entities.PollOption{Id: id}
		assert.True(t, option1.Equals(&option2))
	})

	t.Run("different objects with different IDs are not equal", func(t *testing.T) {
		option1 := entities.PollOption{Id: core.NewUniqueEntityId()}
		option2 := entities.PollOption{Id: core.NewUniqueEntityId()}
		assert.False(t, option1.Equals(&option2))
	})
}
