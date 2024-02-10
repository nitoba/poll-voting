package entities

import (
	"time"

	"github.com/nitoba/poll-voting/internal/domain/core"
)

type Poll struct {
	Id        core.UniqueEntityId
	Title     string
	Options   []*PollOption
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
	Id        *core.UniqueEntityId
	CreatedAt *time.Time
}

func NewPoll(title string, options []*PollOption, optional ...OptionalParams) (*Poll, error) {
	var id core.UniqueEntityId
	if len(optional) > 0 && optional[0].Id != nil {
		id = core.NewUniqueEntityId(optional[0].Id.String())
	} else {
		id = core.NewUniqueEntityId()
	}

	var createdAt time.Time

	if len(optional) > 0 && optional[0].CreatedAt != nil {
		createdAt = *optional[0].CreatedAt
	} else {
		createdAt = time.Now()
	}

	p := &Poll{
		Id:        id,
		Title:     title,
		Options:   options,
		CreatedAt: createdAt,
	}

	return p, nil
}
