package usecases_test

import (
	"testing"

	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases/errors"
	repositories_test "github.com/nitoba/poll-voting/test/repositories"
	"github.com/stretchr/testify/assert"
)

type TestCreatePollUseCaseConfig struct {
	sut  *usecases.CreatePollUseCase
	repo *repositories_test.InMemoryPollsRepository
}

func makeCreatePollUseCase() TestCreatePollUseCaseConfig {
	repo := &repositories_test.InMemoryPollsRepository{}
	sut := usecases.NewCreatePollUseCase(repo)

	return TestCreatePollUseCaseConfig{
		sut:  sut,
		repo: repo,
	}
}

func TestCreatePollUseCase(t *testing.T) {
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

	t.Run("returns error if poll has no options", func(t *testing.T) {
		uc := makeCreatePollUseCase()

		err := uc.sut.Execute(usecases.CreatePollRequest{
			Title: "test",
			Options: []string{
				"option1",
				"option2",
			},
		})

		assert.Nil(t, err)
		assert.NotEmpty(t, uc.repo.Polls[0].Id.String())
		assert.Equal(t, uc.repo.Polls[0].Title, "test")
		assert.Len(t, uc.repo.Polls[0].Options, 2)
		assert.Equal(t, uc.repo.Polls[0].Options[0].Title, "option1")
		assert.Equal(t, uc.repo.Polls[0].Options[1].Title, "option2")
	})
}
