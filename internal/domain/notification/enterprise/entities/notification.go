package entities

import (
	"time"

	"github.com/nitoba/poll-voting/internal/domain/core"
	"github.com/nitoba/poll-voting/internal/domain/notification/enterprise/value_objects"
)

type Notification struct {
	core.Entity
	Title     string
	Content   value_objects.NewVote
	CreatedAt time.Time
}

func NewNotification(title string, content value_objects.NewVote) *Notification {
	return &Notification{
		Entity:    *core.NewEntity(),
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}
}
