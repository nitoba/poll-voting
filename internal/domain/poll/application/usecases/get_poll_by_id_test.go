package usecases_test

import (
	"testing"

	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases/errors"
	"github.com/nitoba/poll-voting/test/factories"
	repositories_test "github.com/nitoba/poll-voting/test/repositories"
	"github.com/stretchr/testify/assert"
)

type TestGetPollByIdUseCaseConfig struct {
	sut  *usecases.GetPollByIdUseCase
	repo *repositories_test.InMemoryPollsRepository
}

func makeGetPollsByIdUseCase() TestGetPollByIdUseCaseConfig {
	repo := &repositories_test.InMemoryPollsRepository{}
	countingVotesRepo := &repositories_test.InMemoryCountingVotesRepository{
		Votes: make(map[string]map[string]int),
	}
	sut := usecases.NewGetPollByIdUseCase(repo, countingVotesRepo)
	return TestGetPollByIdUseCaseConfig{
		sut:  sut,
		repo: repo,
	}
}

func TestGetPollByIdUseCase(t *testing.T) {
	t.Run("it should returns a poll", func(t *testing.T) {
		uc := makeGetPollsByIdUseCase()

		poll := factories.MakePool()

		uc.repo.Polls = append(uc.repo.Polls, poll)

		res, err := uc.sut.Execute(poll.Id.String())

		assert.NoError(t, err)
		assert.True(t, poll.Equals(res))
	})

	t.Run("it should return a error if no existing any poll", func(t *testing.T) {
		uc := makeGetPollsByIdUseCase()
		poll, err := uc.sut.Execute("invalid-poll-id")
		assert.Nil(t, poll)
		assert.ErrorIs(t, err, errors.ErrPollNotFound)
	})
}
