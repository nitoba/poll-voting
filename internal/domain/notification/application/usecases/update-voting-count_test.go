package usecases_test

import (
	"testing"

	"github.com/nitoba/poll-voting/internal/domain/notification/application/usecases"
	"github.com/nitoba/poll-voting/test/factories"
	repositories_test "github.com/nitoba/poll-voting/test/repositories"
	"github.com/stretchr/testify/assert"
)

type UpdateVotingCountUseCaseConfig struct {
	sut  *usecases.UpdateVotingCountUseCase
	repo *repositories_test.InMemoryNotificationsRepository
}

func makeCreateUpdateVotingCountUseCase() UpdateVotingCountUseCaseConfig {
	repo := &repositories_test.InMemoryNotificationsRepository{}
	sut := usecases.NewUpdateVotingCountUseCase(repo)
	return UpdateVotingCountUseCaseConfig{
		sut:  sut,
		repo: repo,
	}
}

func TestUpdateVotingCount(t *testing.T) {

	t.Run("it should be able to send a notification of the update counting votes", func(t *testing.T) {
		config := makeCreateUpdateVotingCountUseCase()
		poll := factories.MakePool()
		req := usecases.UpdateVotingCountUseCaseRequest{
			PollId:       poll.Id.String(),
			PollOptionId: poll.Options[0].Id.String(),
			CountOfVotes: 1,
		}
		err := config.sut.Execute(&req)

		assert.Nil(t, err)
		assert.Len(t, config.repo.Notifications, 1)
	})

}
