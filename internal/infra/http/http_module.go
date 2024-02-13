package http

import (
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases"
	infra_cryptography "github.com/nitoba/poll-voting/internal/infra/cryptography"
	infra_repositories "github.com/nitoba/poll-voting/internal/infra/database/prisma/repositories"
	"github.com/nitoba/poll-voting/internal/infra/http/controllers"
	"github.com/nitoba/poll-voting/prisma/db"
	"github.com/sarulabs/di"
)

type HttpModule struct {
	Hasher               di.Def
	VoterRepository      di.Def
	RegisterVoterUseCase di.Def
	RegisterController   di.Def
}

func NewHttpModule() *HttpModule {
	return &HttpModule{
		Hasher: di.Def{
			Name:  "hasher",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return infra_cryptography.CreateBCryptHasher(), nil
			},
		},
		VoterRepository: di.Def{
			Name:  "voterRepository",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return infra_repositories.NewVotersRepositoryPrisma(ctn.Get("db").(*db.PrismaClient)), nil
			},
		},
		RegisterVoterUseCase: di.Def{
			Name:  "registerVoterUseCase",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return usecases.NewRegisterVoterUseCase(ctn.Get("voterRepository").(*infra_repositories.VotersRepositoryPrisma), ctn.Get("hasher").(*infra_cryptography.BCryptHasher)), nil
			},
		},
		RegisterController: di.Def{
			Name:  "registerController",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return controllers.NewRegisterVoterController(ctn.Get("registerVoterUseCase").(*usecases.RegisterVoterUseCase)), nil
			},
		},
	}
}
