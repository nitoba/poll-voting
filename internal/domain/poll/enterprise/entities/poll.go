package entities

import (
	"time"

	"github.com/nitoba/poll-voting/internal/domain/core/entity"
)

type Poll struct {
	entity.Entity
	Title     string
	Options   []*PollOption
	OwnerId   entity.UniqueEntityId
	CreatedAt time.Time
}

func (p *Poll) AddOption(option *PollOption) {
	if p.Options == nil || len(p.Options) == 0 {
		p.Options = []*PollOption{}
	}

	for _, o := range p.Options {
		if o.Title == option.Title {
			return
		}
	}

	p.Options = append(p.Options, option)
}

func (p *Poll) Equals(other *Poll) bool {
	if p == nil || other == nil {
		return false
	}
	if p == other {
		return true
	}
	return p.Id.String() == other.Id.String()
}

type OptionalParams struct {
	Id        *entity.UniqueEntityId
	CreatedAt *time.Time
}

func NewPoll(title string, options []*PollOption, ownerId entity.UniqueEntityId, optional ...OptionalParams) (*Poll, error) {
	var id entity.UniqueEntityId
	if len(optional) > 0 && optional[0].Id != nil {
		id = entity.NewUniqueEntityId(optional[0].Id.String())
	} else {
		id = entity.NewUniqueEntityId()
	}

	var createdAt time.Time

	if len(optional) > 0 && optional[0].CreatedAt != nil {
		createdAt = *optional[0].CreatedAt
	} else {
		createdAt = time.Now()
	}

	p := &Poll{
		Entity: entity.Entity{
			Id: id,
		},
		Title:     title,
		Options:   options,
		OwnerId:   ownerId,
		CreatedAt: createdAt,
	}

	return p, nil
}
