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

type TestAuthenticateUseCaseConfig struct {
	sut                    *usecases.AuthenticateUseCase
	participantsRepository *repositories_test.InMemoryParticipantsRepository
	HashComparer           *cryptography_test.FakeHasher
	encrypter              *cryptography_test.FakeEncrypter
}

func makeAuthenticateUseCase() TestAuthenticateUseCaseConfig {
	participantsRepository := &repositories_test.InMemoryParticipantsRepository{}
	HashComparer := &cryptography_test.FakeHasher{}
	encrypter := &cryptography_test.FakeEncrypter{}
	sut := usecases.NewAuthenticateUseCase(participantsRepository, HashComparer, encrypter)

	return TestAuthenticateUseCaseConfig{
		sut:                    sut,
		participantsRepository: participantsRepository,
		HashComparer:           HashComparer,
		encrypter:              encrypter,
	}
}

func TestAuthenticateUseCase(t *testing.T) {
	t.Run("returns access token if credentials are valid", func(t *testing.T) {
		uc := makeAuthenticateUseCase()

		password, _ := uc.HashComparer.Hash("secret")

		participant := factories.MakeParticipant(map[string]interface{}{
			"email":    "john.doe@gmail.com",
			"password": password,
		})

		uc.participantsRepository.Participants = append(uc.participantsRepository.Participants, participant)

		req := usecases.AuthenticateRequest{
			Email:    participant.Email.Value(),
			Password: "secret",
		}

		res, err := uc.sut.Execute(req)

		assert.Nil(t, err)
		assert.NotEmpty(t, res.AccessToken)
	})

	t.Run("returns error if credentials are invalid", func(t *testing.T) {
		uc := makeAuthenticateUseCase()

		_, err := uc.sut.Execute(usecases.AuthenticateRequest{
			Email:    "invalid@email.com",
			Password: "<PASSWORD>",
		})

		assert.ErrorIs(t, err, errors.ErrWrongCredentials)

		password, _ := uc.HashComparer.Hash("secret")

		participant := factories.MakeParticipant(map[string]interface{}{
			"email":    "john.doe@gmail.com",
			"password": password,
		})

		uc.participantsRepository.Participants = append(uc.participantsRepository.Participants, participant)

		_, err = uc.sut.Execute(usecases.AuthenticateRequest{
			Email:    "invalid@email.com",
			Password: "secret",
		})

		assert.ErrorIs(t, err, errors.ErrWrongCredentials)
	})
}
