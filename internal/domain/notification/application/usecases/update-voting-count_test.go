package usecases_test

import (
	"testing"

	"github.com/nitoba/poll-voting/internal/domain/notification/application/usecases"
	"github.com/nitoba/poll-voting/test/factories"
	messaging_test "github.com/nitoba/poll-voting/test/messaging"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UpdateVotingCountUseCaseConfig struct {
	sut              *usecases.UpdateVotingCountUseCase
	messagePublisher *messaging_test.InMemoryMessagePublisher
}

func makeCreateUpdateVotingCountUseCase() UpdateVotingCountUseCaseConfig {
	messagePublisher := &messaging_test.InMemoryMessagePublisher{}
	sut := usecases.NewUpdateVotingCountUseCase(messagePublisher)
	return UpdateVotingCountUseCaseConfig{
		sut:              sut,
		messagePublisher: messagePublisher,
	}
}

func TestUpdateVotingCount(t *testing.T) {

	t.Run("it should be able to send a notification of the update counting votes", func(t *testing.T) {
		config := makeCreateUpdateVotingCountUseCase()

		config.messagePublisher.On("Publish", mock.Anything).Return(nil)

		poll := factories.MakePool()
		req := usecases.UpdateVotingCountUseCaseRequest{
			PollId:       poll.Id.String(),
			PollOptionId: poll.Options[0].Id.String(),
			CountOfVotes: 1,
		}
		err := config.sut.Execute(&req)
		assert.Nil(t, err)
		config.messagePublisher.AssertCalled(t, "Publish", mock.Anything)
		config.messagePublisher.AssertNumberOfCalls(t, "Publish", 1)
	})

}
