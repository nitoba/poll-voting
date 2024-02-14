package controllers_test

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	configs "github.com/nitoba/poll-voting/config"
	"github.com/nitoba/poll-voting/internal/infra/database/prisma"
	http_module "github.com/nitoba/poll-voting/internal/infra/http"
	"github.com/nitoba/poll-voting/internal/infra/http/server"
	"github.com/nitoba/poll-voting/pkg/di"
	"github.com/nitoba/poll-voting/prisma/db"
	"github.com/nitoba/poll-voting/test"
	"github.com/nitoba/poll-voting/test/factories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RegisterVoterControllerTestSuite struct {
	suite.Suite
	e *httpexpect.Expect
}

// Run this function before the all tests
func (s *RegisterVoterControllerTestSuite) SetupSuite() {
	di.InitContainer()

	httpModule := http_module.NewHttpModule()

	httpModule.Build()

	di.RegisterModuleProviders(httpModule.GetDependencies())

	di.BuildDependencies()

	test.SetupDatabase()

	server := server.GetServer()
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
}

// Run this function after the all tests
func (suite *RegisterVoterControllerTestSuite) TearDownSuite() {
	test.AfterAll()
}

// Run this function before every test
func (suite *RegisterVoterControllerTestSuite) SetupTest() {
	test.TruncateTables()
}

// Run this function after every test
// func (suite *RegisterVoterControllerTestSuite) TearDownTest() {
// 	test.TruncateTables()
// }

func TestSuit(t *testing.T) {
	// Register the test suite
	suite.Run(t, new(RegisterVoterControllerTestSuite))
}

func (suite *RegisterVoterControllerTestSuite) TestHandle() {
	suite.Run("should return 204 if voter was created", func() {
		suite.e.POST("/auth/register").WithJSON(map[string]interface{}{
			"name":     "John Doe",
			"email":    "john.doe@gmail.com",
			"password": "123456",
		}).Expect().Status(http.StatusCreated)

		voter, _ := prisma.GetDB().Voter.FindUnique(db.Voter.Email.Equals("john.doe@gmail.com")).Exec(configs.GetConfig().Ctx)

		assert.NotEmpty(suite.T(), voter.ID)
		assert.Equal(suite.T(), "John Doe", voter.Name)
		assert.Equal(suite.T(), "john.doe@gmail.com", voter.Email)
	})
}

func (suite *RegisterVoterControllerTestSuite) TestHandleInvalidData() {
	suite.Run("should return 400 if the data if not valid", func() {
		suite.e.POST("/auth/register").WithJSON(map[string]interface{}{
			"name":     "",
			"email":    "john.doe@gmail.com",
			"password": "123456",
		}).Expect().Status(http.StatusBadRequest).JSON().Object().ContainsKey("message")
	})
}

func (suite *RegisterVoterControllerTestSuite) TestHandleVoterAlreadyExists() {
	suite.Run("should return 409 if the voter already exists", func() {
		factories.MakePrismaVoter(factories.OptionalVoterParams{
			Email: "john.doe@gmail.com",
		})

		suite.e.POST("/auth/register").WithJSON(map[string]interface{}{
			"name":     "John Doe",
			"email":    "john.doe@gmail.com",
			"password": "123456",
		}).Expect().Status(http.StatusConflict).JSON().Object().HasValue("message", "voter already exists")
	})
}
