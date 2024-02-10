package usecases_test

import (
	"testing"

	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases"
	"github.com/nitoba/poll-voting/test/factories"
	repositories_test "github.com/nitoba/poll-voting/test/repositories"
	"github.com/stretchr/testify/assert"
)

type TestFetchPollsUseCaseConfig struct {
	sut  *usecases.FetchPollsUseCase
	repo *repositories_test.InMemoryPollsRepository
}

func makeFetchPollsUseCase() TestFetchPollsUseCaseConfig {
	repo := &repositories_test.InMemoryPollsRepository{}
	sut := usecases.NewFetchPollsUseCase(repo)
	return TestFetchPollsUseCaseConfig{
		sut:  sut,
		repo: repo,
	}
}

func TestFetchPollsUseCase(t *testing.T) {
	t.Run("it should returns polls", func(t *testing.T) {
		uc := makeFetchPollsUseCase()

		for i := 0; i < 3; i++ {
			poll := factories.MakePool()
			uc.repo.Polls = append(uc.repo.Polls, poll)
		}

		polls, err := uc.sut.Execute()

		assert.NoError(t, err)
		assert.Len(t, polls, 3)
	})

	t.Run("it should return a empty list if not exiting polls", func(t *testing.T) {
		uc := makeFetchPollsUseCase()
		polls, err := uc.sut.Execute()
		assert.NoError(t, err)
		assert.Len(t, polls, 0)
	})
}
