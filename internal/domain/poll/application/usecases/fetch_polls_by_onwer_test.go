package usecases_test

import (
	"testing"

	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases"
	"github.com/nitoba/poll-voting/test/factories"
	repositories_test "github.com/nitoba/poll-voting/test/repositories"
	"github.com/stretchr/testify/assert"
)

type TestFetchPollsByOwnerUseCaseConfig struct {
	sut  *usecases.FetchPollsByOwnerUseCase
	repo *repositories_test.InMemoryPollsRepository
}

func makeFetchPollsByOwnerUseCase() TestFetchPollsByOwnerUseCaseConfig {
	repo := &repositories_test.InMemoryPollsRepository{}
	sut := usecases.NewFetchPollsByOwnerUseCase(repo)
	return TestFetchPollsByOwnerUseCaseConfig{
		sut:  sut,
		repo: repo,
	}
}

func TestFetchPollsByOwnerUseCase(t *testing.T) {
	t.Run("it should returns polls of your owner", func(t *testing.T) {
		uc := makeFetchPollsByOwnerUseCase()
		owner := factories.MakeVoter()

		for i := 0; i < 3; i++ {
			poll := factories.MakePool(factories.OptionalPollParams{
				OwnerId: &owner.Id,
			})
			uc.repo.Polls = append(uc.repo.Polls, poll)
		}

		polls, err := uc.sut.Execute(usecases.FetchPollsByOwnerRequest{
			OwnerId: owner.Id.String(),
		})

		assert.NoError(t, err)
		assert.Len(t, polls.Polls, 3)
	})

	t.Run("it should return a empty list if not exiting polls", func(t *testing.T) {
		uc := makeFetchPollsByOwnerUseCase()
		polls, err := uc.sut.Execute(usecases.FetchPollsByOwnerRequest{
			OwnerId: factories.MakeVoter().Id.String(),
		})
		assert.NoError(t, err)
		assert.Len(t, polls.Polls, 0)
	})
}
