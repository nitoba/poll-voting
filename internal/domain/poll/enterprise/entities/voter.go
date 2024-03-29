package entities

import (
	"errors"

	"github.com/nitoba/poll-voting/internal/domain/core"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/value_objects"
)

var ErrMissingArgument = errors.New("missing argument")

type Voter struct {
	core.Entity
	Name     string
	Email    *value_objects.Email
	Password string
}

func (p *Voter) Equals(other *Voter) bool {
	if p == nil || other == nil {
		return false
	}
	if p == other {
		return true
	}
	return p.Id.String() == other.Id.String()
}

func (*Voter) validate(name, password string) error {
	if name == "" {
		return ErrMissingArgument
	}
	if password == "" {
		return ErrMissingArgument
	}
	return nil
}

func NewVoter(name string, email *value_objects.Email, password string, id ...core.UniqueEntityId) (*Voter, error) {
	var ID core.UniqueEntityId
	if len(id) > 0 {
		ID = id[0]
	} else {
		ID = core.NewUniqueEntityId()
	}

	p := &Voter{
		Entity:   *core.NewEntity(ID),
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err := p.validate(name, password); err != nil {
		return nil, err
	}

	return p, nil
}
