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
	sut              *usecases.CreatePollUseCase
	repo             *repositories_test.InMemoryPollsRepository
	votersRepository *repositories_test.InMemoryVotersRepository
}

func makeCreatePollUseCase() TestCreatePollUseCaseConfig {
	votersRepository := &repositories_test.InMemoryVotersRepository{}
	repo := &repositories_test.InMemoryPollsRepository{}
	sut := usecases.NewCreatePollUseCase(repo, votersRepository)

	return TestCreatePollUseCaseConfig{
		sut:              sut,
		repo:             repo,
		votersRepository: votersRepository,
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
		owner := factories.MakeVoter()
		uc.votersRepository.Voters = append(uc.votersRepository.Voters, owner)

		err := uc.sut.Execute(usecases.CreatePollRequest{
			Title:   "test",
			Options: []string{},
			OwnerId: owner.Id.String(),
		})

		println(err.Error())

		assert.ErrorIs(t, err, errors.ErrInvalidPoll)

		err = uc.sut.Execute(usecases.CreatePollRequest{
			Title: "test",
			Options: []string{
				"test",
			},
			OwnerId: owner.Id.String(),
		})

		assert.ErrorIs(t, err, errors.ErrInvalidPoll)
	})

	t.Run("it should be able to create a new poll", func(t *testing.T) {
		uc := makeCreatePollUseCase()
		owner := factories.MakeVoter()

		uc.votersRepository.Voters = append(uc.votersRepository.Voters, owner)

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
