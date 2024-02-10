package usecases_test

import (
	"testing"

	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases/errors"
	cryptography_test "github.com/nitoba/poll-voting/test/cryptography"
	"github.com/nitoba/poll-voting/test/factories"
	repositories_test "github.com/nitoba/poll-voting/test/repositories"
	"github.com/stretchr/testify/assert"
)

type TestRegisterUseCaseConfig struct {
	sut            *usecases.RegisterParticipantUseCase
	userRepository *repositories_test.InMemoryParticipantsRepository
	hashGenerator  *cryptography_test.FakeHasher
}

func makeRegisterUseCase() TestRegisterUseCaseConfig {
	userRepository := &repositories_test.InMemoryParticipantsRepository{}
	hashGenerator := &cryptography_test.FakeHasher{}
	sut := usecases.NewRegisterParticipantUseCase(userRepository, hashGenerator)

	return TestRegisterUseCaseConfig{
		sut:            sut,
		userRepository: userRepository,
		hashGenerator:  hashGenerator,
	}
}

func TestRegisterParticipantUseCase(t *testing.T) {
	t.Run("it should register a new participant", func(t *testing.T) {
		res := makeRegisterUseCase()

		err := res.sut.Execute(&usecases.RegisterParticipantRequest{
			Name:     "John Doe",
			Email:    "john.doe@gmail.com",
			Password: "123456",
		})

		assert.Nil(t, err)
		assert.NotEmpty(t, res.userRepository.Participants[0].Id.String())
		assert.Equal(t, res.userRepository.Participants[0].Name, "John Doe")
		assert.Equal(t, res.userRepository.Participants[0].Email.Value(), "john.doe@gmail.com")
		assert.Equal(t, res.userRepository.Participants[0].Password, "hashed:123456")
	})

	t.Run("it should no be able register a new participant if already exists", func(t *testing.T) {
		res := makeRegisterUseCase()

		p := factories.MakeParticipant(map[string]interface{}{"email": "john.doe@gmail.com"})

		res.userRepository.Participants = append(res.userRepository.Participants, p)

		err := res.sut.Execute(&usecases.RegisterParticipantRequest{
			Name:     "John Doe",
			Email:    p.Email.Value(),
			Password: "123456",
		})

		assert.ErrorIs(t, err, errors.ErrParticipantAlreadyExists)
	})
}
