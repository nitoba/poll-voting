package ws

import (
	"github.com/gorilla/websocket"
	"github.com/nitoba/poll-voting/internal/domain/notification/enterprise/entities"
)

type UpdateCountingVotesPublisher struct {
	conn *websocket.Conn
}

type Message struct {
	Title        string `json:"title"`
	VoteCount    int    `json:"vote_count"`
	PollOptionId string `json:"poll_option_id"`
}

func (p *UpdateCountingVotesPublisher) Publish(message *entities.Notification) error {
	msg := &Message{
		Title:        message.Title,
		VoteCount:    message.Content.Count,
		PollOptionId: message.Content.PollOptionId,
	}
	if err := p.conn.WriteJSON(msg); err != nil {
		return err
	}
	return nil
}

func NewUpdateCountingVotesPublisher(conn *websocket.Conn) *UpdateCountingVotesPublisher {
	return &UpdateCountingVotesPublisher{
		conn: conn,
	}
}
