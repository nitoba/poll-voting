package entities

import (
	"errors"

	"github.com/nitoba/poll-voting/internal/domain/core"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/value_objects"
)

var ErrMissingArgument = errors.New("missing argument")

type Participant struct {
	Id       core.UniqueEntityId
	Name     string
	Email    *value_objects.Email
	Password string
}

func (p *Participant) Equals(other *Participant) bool {
	if p == nil || other == nil {
		return false
	}
	if p == other {
		return true
	}
	return p.Id.String() == other.Id.String()
}

func (*Participant) validate(name, password string) error {
	if name == "" {
		return ErrMissingArgument
	}
	if password == "" {
		return ErrMissingArgument
	}
	return nil
}

func NewParticipant(name string, email *value_objects.Email, password string, id ...core.UniqueEntityId) (*Participant, error) {
	var ID core.UniqueEntityId
	if len(id) > 0 {
		ID = id[0]
	} else {
		ID = core.NewUniqueEntityId()
	}

	p := &Participant{
		Id:       ID,
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err := p.validate(name, password); err != nil {
		return nil, err
	}

	return p, nil
}
