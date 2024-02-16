package controllers_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/gavv/httpexpect/v2"
	configs "github.com/nitoba/poll-voting/config"
	"github.com/nitoba/poll-voting/internal/domain/core"
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

type CreatePollControllerTestSuite struct {
	suite.Suite
	e          *httpexpect.Expect
	httpModule *http_module.HttpModule
}

// Run this function before the all tests
func (s *CreatePollControllerTestSuite) SetupSuite() {
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
func (suite *CreatePollControllerTestSuite) TearDownSuite() {
	test.AfterAll()
}

// Run this function before every test
func (suite *CreatePollControllerTestSuite) SetupTest() {
	test.TruncateTables()
}

// Run this function after every test
// func (suite *CreatePollControllerTestSuite) TearDownTest() {
// 	test.TruncateTables()
// }

func TestCreatePollControllerSuite(t *testing.T) {
	// Register the test suite
	if os.Getenv("IGNORE_E2E") != "" {
		t.Skip("Ignorando testes E2E")
	}
	suite.Run(t, new(CreatePollControllerTestSuite))
}

func (suite *CreatePollControllerTestSuite) TestE2EHandle() {
	suite.Run("should return 201 if polls was created", func() {
		userID := core.NewUniqueEntityId()

		factories.MakePrismaVoter(factories.OptionalVoterParams{
			Id: &userID,
		})

		token := di.GetContainer().Get("encrypter").(*cryptography.JWTEncrypter).Encrypt(map[string]interface{}{
			"sub": userID.String(),
		})

		suite.e.POST("/polls/").WithHeader("Authorization", "Bearer "+token).WithJSON(map[string]interface{}{
			"title":   "Poll example",
			"options": []string{"Option 1", "Option 2"},
		}).Expect().Status(http.StatusCreated)

		poll, _ := prisma.GetDB().Poll.FindFirst(db.Poll.Title.Equals("Poll example")).Exec(configs.GetConfig().Ctx)
		assert.NotEmpty(suite.T(), poll.ID)
		assert.Equal(suite.T(), "Poll example", poll.Title)
	})
}

// func (suite *CreatePollControllerTestSuite) TestE2EHandleInvalidData() {
// 	suite.Run("should return 400 if voter data is not valid", func() {
// 		suite.e.POST("/auth/register").WithJSON(map[string]interface{}{
// 			"email":    "john.doe@gmail.com",
// 			"password": "123456",
// 		}).Expect().Status(http.StatusBadRequest).JSON().Object().ContainsKey("message")
// 	})
// }
