package ws

import (
	"encoding/json"

	"github.com/nitoba/poll-voting/internal/domain/notification/enterprise/entities"
)

type UpdateCountingVotesPublisher struct{}

type Messages struct {
	Title        string `json:"title"`
	VoteCount    int    `json:"vote_count"`
	PollOptionId string `json:"poll_option_id"`
}

func (p *UpdateCountingVotesPublisher) Publish(message *entities.Notification) error {
	msg := &Messages{
		Title:        message.Title,
		VoteCount:    message.Content.Count,
		PollOptionId: message.Content.PollOptionId,
	}

	msgBytes, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	Manager.send(msgBytes, nil)
	return nil
}

func NewUpdateCountingVotesPublisher() *UpdateCountingVotesPublisher {
	return &UpdateCountingVotesPublisher{}
}
