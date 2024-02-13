package factories

import (
	"github.com/jaswdr/faker"
	configs "github.com/nitoba/poll-voting/config"
	"github.com/nitoba/poll-voting/internal/domain/core"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/value_objects"
	"github.com/nitoba/poll-voting/internal/infra/database/prisma"
	"github.com/nitoba/poll-voting/prisma/db"
)

func MakeVoter(props ...map[string]interface{}) *entities.Voter {
	fake := faker.New()
	id := core.NewUniqueEntityId()
	name := fake.Lorem().Word()
	email := fake.Internet().Email()
	password := fake.Internet().Password()

	if len(props) > 0 && props[0]["id"] != nil {
		id = core.NewUniqueEntityId(props[0]["id"].(string))
	}

	if len(props) > 0 && props[0]["name"] != nil {
		name = props[0]["name"].(string)
	}

	if len(props) > 0 && props[0]["email"] != nil {
		email = props[0]["email"].(string)
	}

	if len(props) > 0 && props[0]["password"] != nil {
		password = props[0]["password"].(string)
	}

	emailVo, _ := value_objects.NewEmail(email)

	return &entities.Voter{
		Entity: core.Entity{
			Id: id,
		},
		Name:     name,
		Email:    emailVo,
		Password: password,
	}
}

type OptionalVoterParams struct {
	Id       *core.UniqueEntityId
	Name     string
	Email    string
	Password string
}

func MakePrismaVoter(props ...OptionalVoterParams) error {
	fake := faker.New()
	id := core.NewUniqueEntityId()
	name := fake.Lorem().Word()
	email := fake.Internet().Email()
	password := fake.Internet().Password()

	if len(props) > 0 && props[0].Id != nil {
		id = core.NewUniqueEntityId(props[0].Id.String())
	}

	if len(props) > 0 && props[0].Name != "" {
		name = props[0].Name
	}

	if len(props) > 0 && props[0].Email != "" {
		email = props[0].Email
	}

	if len(props) > 0 && props[0].Password != "" {
		password = props[0].Password
	}

	_, err := prisma.GetDB().Voter.CreateOne(
		db.Voter.Name.Set(name),
		db.Voter.Email.Set(email),
		db.Voter.Password.Set(password),
		db.Voter.ID.Set(id.String()),
	).Exec(configs.GetConfig().Ctx)

	if err != nil {
		println("Error to create voter: ", err.Error())
		panic("Error to create voter")
	}

	return nil
}
