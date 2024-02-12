package repositories

import (
	configs "github.com/nitoba/poll-voting/config"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
	"github.com/nitoba/poll-voting/internal/infra/database/prisma/mappers"
	"github.com/nitoba/poll-voting/prisma/db"
)

type VotersRepositoryPrisma struct {
	db *db.PrismaClient
}

func (r *VotersRepositoryPrisma) Create(voter *entities.Voter) error {
	ctx := configs.GetConfig().Ctx
	if _, err := r.db.Voter.CreateOne(
		db.Voter.Name.Set(voter.Name),
		db.Voter.Email.Set(voter.Email.Value()),
		db.Voter.Password.Set(voter.Password),
		db.Voter.ID.Set(voter.Id.String()),
	).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r *VotersRepositoryPrisma) FindByEmail(email string) (*entities.Voter, error) {
	ctx := configs.GetConfig().Ctx
	voter, err := r.db.Voter.FindUnique(db.Voter.Email.Equals(email)).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return mappers.VoterToDomain(voter), nil
}

func (r *VotersRepositoryPrisma) FindById(id string) (*entities.Voter, error) {
	ctx := configs.GetConfig().Ctx
	voter, err := r.db.Voter.FindUnique(db.Voter.ID.Equals(id)).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return mappers.VoterToDomain(voter), nil
}

func NewVotersRepositoryPrisma(db *db.PrismaClient) *VotersRepositoryPrisma {
	return &VotersRepositoryPrisma{
		db: db,
	}
}
