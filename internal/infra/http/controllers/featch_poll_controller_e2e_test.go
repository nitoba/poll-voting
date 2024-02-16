package controllers_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/gavv/httpexpect/v2"
	configs "github.com/nitoba/poll-voting/config"
	"github.com/nitoba/poll-voting/internal/domain/core"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
	"github.com/nitoba/poll-voting/internal/infra/cryptography"
	"github.com/nitoba/poll-voting/internal/infra/database/prisma"
	http_module "github.com/nitoba/poll-voting/internal/infra/http"
	"github.com/nitoba/poll-voting/internal/infra/http/rest"
	"github.com/nitoba/poll-voting/pkg/di"
	"github.com/nitoba/poll-voting/prisma/db"
	"github.com/nitoba/poll-voting/test"
	"github.com/nitoba/poll-voting/test/factories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FetchPollsControllerTestSuite struct {
	suite.Suite
	e          *httpexpect.Expect
	httpModule *http_module.HttpModule
}

// Run this function before the all tests
func (s *FetchPollsControllerTestSuite) SetupSuite() {
	di.InitContainer()

	httpModule := http_module.NewHttpModule()

	httpModule.Build()

	di.RegisterModuleProviders(httpModule.GetProviders())

	di.BuildDependencies()

	test.SetupDatabase()

	server := rest.GetServer()
	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(server),
			Jar:       httpexpect.NewCookieJar(),
		},
		Reporter: httpexpect.NewAssertReporter(s.T()),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(s.T(), true),
		},
	})
	s.e = e
	s.httpModule = httpModule
}

// Run this function after the all tests
func (suite *FetchPollsControllerTestSuite) TearDownSuite() {
	test.AfterAll()
}

// Run this function before every test
func (suite *FetchPollsControllerTestSuite) SetupTest() {
	test.TruncateTables()
}

// Run this function after every test
// func (suite *FetchPollsControllerTestSuite) TearDownTest() {
// 	test.TruncateTables()
// }

func TestFetchPollsControllerSuite(t *testing.T) {
	// Register the test suite
	if os.Getenv("IGNORE_E2E") != "" {
		t.Skip("Ignorando testes E2E")
	}
	suite.Run(t, new(FetchPollsControllerTestSuite))
}

func (suite *FetchPollsControllerTestSuite) TestE2EHandle() {
	suite.Run("should return a list of polls", func() {
		userID := core.NewUniqueEntityId()

		factories.MakePrismaVoter(factories.OptionalVoterParams{
			Id: &userID,
		})
		factories.MakePrismaPoll(factories.OptionalPollParams{
			Title:   "Poll example",
			OwnerId: &userID,
			Options: []*entities.PollOption{
				factories.MakePoolOption(factories.OptionalPollOptionParams{
					Title: "Option 1",
				}),
				factories.MakePoolOption(factories.OptionalPollOptionParams{
					Title: "Option 2",
				}),
			},
		})

		token := di.GetContainer().Get("encrypter").(*cryptography.JWTEncrypter).Encrypt(map[string]interface{}{
			"sub": userID.String(),
		})

		suite.e.GET("/polls/").WithHeader("Authorization", "Bearer "+token).Expect().Status(http.StatusOK).JSON().Array().NotEmpty()

		poll, _ := prisma.GetDB().Poll.FindFirst(db.Poll.Title.Equals("Poll example")).Exec(configs.GetConfig().Ctx)
		assert.NotEmpty(suite.T(), poll.ID)
		assert.Equal(suite.T(), "Poll example", poll.Title)
	})
}

// func (suite *FetchPollsControllerTestSuite) TestE2EHandleInvalidData() {
// 	suite.Run("should return 400 if voter data is not valid", func() {
// 		suite.e.POST("/auth/register").WithJSON(map[string]interface{}{
// 			"email":    "john.doe@gmail.com",
// 			"password": "123456",
// 		}).Expect().Status(http.StatusBadRequest).JSON().Object().ContainsKey("message")
// 	})
// }
