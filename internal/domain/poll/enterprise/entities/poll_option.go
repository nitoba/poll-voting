package entities

import "github.com/nitoba/poll-voting/internal/domain/core/entity"

type PollOption struct {
	entity.Entity
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

func NewPollOption(title string, id ...entity.UniqueEntityId) (*PollOption, error) {
	var ID entity.UniqueEntityId
	if len(id) > 0 {
		ID = id[0]
	} else {
		ID = entity.NewUniqueEntityId()
	}
	p := &PollOption{
		Entity: entity.Entity{
			Id: ID,
		},
		Title: title,
	}

	return p, nil
}
