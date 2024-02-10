package factories

import (
	"github.com/jaswdr/faker"
	"github.com/nitoba/poll-voting/internal/domain/core"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/value_objects"
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
		Id:       id,
		Name:     name,
		Email:    emailVo,
		Password: password,
	}
}
