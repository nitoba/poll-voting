package mappers

import (
	"github.com/nitoba/poll-voting/internal/domain/core"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/value_objects"
	"github.com/nitoba/poll-voting/prisma/db"
)

func VoterToDomain(voter *db.VoterModel) *entities.Voter {
	email, _ := value_objects.NewEmail(voter.Email)
	return &entities.Voter{
		Entity:   *core.NewEntity(core.NewUniqueEntityId(voter.ID)),
		Name:     voter.Name,
		Email:    email,
		Password: voter.Password,
	}
}
