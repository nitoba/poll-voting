package entities

import (
	"errors"

	"github.com/nitoba/poll-voting/internal/domain/core/entity"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/value_objects"
)

var ErrMissingArgument = errors.New("missing argument")

type Voter struct {
	entity.Entity
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

func NewVoter(name string, email *value_objects.Email, password string, id ...entity.UniqueEntityId) (*Voter, error) {
	var ID entity.UniqueEntityId
	if len(id) > 0 {
		ID = id[0]
	} else {
		ID = entity.NewUniqueEntityId()
	}

	p := &Voter{
		Entity: entity.Entity{
			Id: ID,
		},
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err := p.validate(name, password); err != nil {
		return nil, err
	}

	return p, nil
}
