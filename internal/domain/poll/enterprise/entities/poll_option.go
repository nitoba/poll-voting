package entities

import "github.com/nitoba/poll-voting/internal/domain/core"

type PollOption struct {
	core.Entity
	Title string
}

func (p *PollOption) Equals(other *PollOption) bool {
	if p == nil || other == nil {
		return false
	}
	if p == other {
		return true
	}
	return p.Id.String() == other.Id.String()
}

func NewPollOption(title string, id ...core.UniqueEntityId) (*PollOption, error) {
	var ID core.UniqueEntityId
	if len(id) > 0 {
		ID = id[0]
	} else {
		ID = core.NewUniqueEntityId()
	}
	p := &PollOption{
		Entity: core.Entity{
			Id: ID,
		},
		Title: title,
	}

	return p, nil
}
