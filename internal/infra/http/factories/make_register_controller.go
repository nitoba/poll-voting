package factories

import "github.com/nitoba/poll-voting/internal/infra/http/controllers"

func MakeRegisterController() *controllers.RegisterVoterController {
	return controllers.NewRegisterVoterController()
}
