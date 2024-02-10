package usecases_test

import (
	"testing"

	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases/errors"
	"github.com/nitoba/poll-voting/test/factories"
	repositories_test "github.com/nitoba/poll-voting/test/repositories"
	"github.com/stretchr/testify/assert"
)

type TestCreatePollUseCaseConfig struct {
	sut                    *usecases.CreatePollUseCase
	repo                   *repositories_test.InMemoryPollsRepository
	participantsRepository *repositories_test.InMemoryParticipantsRepository
}

func makeCreatePollUseCase() TestCreatePollUseCaseConfig {
	participantsRepository := &repositories_test.InMemoryParticipantsRepository{}
	repo := &repositories_test.InMemoryPollsRepository{}
	sut := usecases.NewCreatePollUseCase(repo, participantsRepository)

	return TestCreatePollUseCaseConfig{
		sut:                    sut,
		repo:                   repo,
		participantsRepository: participantsRepository,
	}
}

func TestCreatePollUseCase(t *testing.T) {
	t.Run("returns error if owner does not exists", func(t *testing.T) {
		uc := makeCreatePollUseCase()

		err := uc.sut.Execute(usecases.CreatePollRequest{
			Title:   "test",
			Options: []string{},
			OwnerId: "invalid-owner-id",
		})

		assert.ErrorIs(t, err, errors.ErrInvalidOwner)
	})

	t.Run("returns error if poll has no at least 2 options", func(t *testing.T) {
		uc := makeCreatePollUseCase()

		err := uc.sut.Execute(usecases.CreatePollRequest{
			Title:   "test",
			Options: []string{},
		})

		assert.ErrorIs(t, err, errors.ErrInvalidPoll)

		err = uc.sut.Execute(usecases.CreatePollRequest{
			Title: "test",
			Options: []string{
				"test",
			},
		})

		assert.ErrorIs(t, err, errors.ErrInvalidPoll)
	})

	t.Run("it should be able to create a new poll", func(t *testing.T) {
		uc := makeCreatePollUseCase()
		owner := factories.MakeParticipant()

		uc.participantsRepository.Participants = append(uc.participantsRepository.Participants, owner)

		err := uc.sut.Execute(usecases.CreatePollRequest{
			Title: "test",
			Options: []string{
				"option1",
				"option2",
			},
			OwnerId: owner.Id.String(),
		})

		assert.Nil(t, err)
		assert.NotEmpty(t, uc.repo.Polls[0].Id.String())
		assert.Equal(t, uc.repo.Polls[0].OwnerId.String(), owner.Id.String())
		assert.Equal(t, uc.repo.Polls[0].Title, "test")
		assert.Len(t, uc.repo.Polls[0].Options, 2)
		assert.Equal(t, uc.repo.Polls[0].Options[0].Title, "option1")
		assert.Equal(t, uc.repo.Polls[0].Options[1].Title, "option2")
	})
}
