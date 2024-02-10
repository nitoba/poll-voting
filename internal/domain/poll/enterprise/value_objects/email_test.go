package value_objects_test

import (
	"testing"

	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/value_objects"
	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	t.Run("it should create a new email", func(t *testing.T) {
		email, err := value_objects.NewEmail("john.doe@gmail.com")
		assert.Nil(t, err)
		assert.Equal(t, email.Value(), "john.doe@gmail.com")
	})

	t.Run("it should not be able create a new email if invalid", func(t *testing.T) {
		email, err := value_objects.NewEmail("john.doegmail.com")
		assert.Nil(t, email)
		assert.ErrorIs(t, err, value_objects.ErrInvalidEmail)
	})
}
