package entities_test

import (
	"testing"

	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/value_objects"
	"github.com/stretchr/testify/assert"
)

func TestParticipant(t *testing.T) {
	t.Run("it should create a new participant", func(t *testing.T) {
		email, _ := value_objects.NewEmail("john.doe@gmail.com")
		p, err := entities.NewParticipant("John Doe", email, "123456")
		assert.Nil(t, err)
		assert.NotEmpty(t, p.Id.String())
		assert.Equal(t, p.Name, "John Doe")
		assert.Equal(t, p.Email.Value(), "john.doe@gmail.com")
		assert.Equal(t, p.Password, "123456")
	})

	t.Run("it should not be able create a new participant", func(t *testing.T) {
		email, _ := value_objects.NewEmail("john.doe@gmail.com")
		p, err := entities.NewParticipant("", email, "123456")
		assert.Nil(t, p)
		assert.ErrorIs(t, err, entities.ErrMissingArgument)
		p, err = entities.NewParticipant("John Doe", email, "")
		assert.Nil(t, p)
		assert.ErrorIs(t, err, entities.ErrMissingArgument)
	})
}
