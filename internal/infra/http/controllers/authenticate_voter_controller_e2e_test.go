package controllers_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/nitoba/poll-voting/internal/infra/cryptography"
	http_module "github.com/nitoba/poll-voting/internal/infra/http"
	"github.com/nitoba/poll-voting/internal/infra/http/rest"
	"github.com/nitoba/poll-voting/pkg/di"
	"github.com/nitoba/poll-voting/test"
	"github.com/nitoba/poll-voting/test/factories"
	"github.com/stretchr/testify/suite"
)

type AuthenticateVoterControllerTestSuite struct {
	suite.Suite
	e          *httpexpect.Expect
	httpModule *http_module.HttpModule
}

// Run this function before the all tests
func (s *AuthenticateVoterControllerTestSuite) SetupSuite() {
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
func (suite *AuthenticateVoterControllerTestSuite) TearDownSuite() {
	test.AfterAll()
}

// Run this function before every test
func (suite *AuthenticateVoterControllerTestSuite) SetupTest() {
	test.TruncateTables()
}

// Run this function after every test
// func (suite *AuthenticateVoterControllerTestSuite) TearDownTest() {
// 	test.TruncateTables()
// }

func TestAuthenticateVoterControllerSuite(t *testing.T) {
	// Register the test suite
	if os.Getenv("IGNORE_E2E") != "" {
		t.Skip("Ignorando testes E2E")
	}
	suite.Run(t, new(AuthenticateVoterControllerTestSuite))
}

func (suite *AuthenticateVoterControllerTestSuite) TestE2EHandle() {
	suite.Run("should return 200 if voter was authenticated", func() {
		password, _ := di.GetContainer().Get("hasher").(*cryptography.BCryptHasher).Hash("123456")

		factories.MakePrismaVoter(factories.OptionalVoterParams{
			Email:    "john.doe@gmail.com",
			Password: password,
		})
		suite.e.POST("/auth/authenticate").WithJSON(map[string]interface{}{
			"email":    "john.doe@gmail.com",
			"password": "123456",
		}).Expect().Status(http.StatusOK).JSON().Object().Value("access_token").String().NotEmpty()
	})
}

func (suite *AuthenticateVoterControllerTestSuite) TestE2EHandleInvalidData() {
	suite.Run("should return 400 if voter data is not valid", func() {
		suite.e.POST("/auth/register").WithJSON(map[string]interface{}{
			"email":    "john.doe@gmail.com",
			"password": "123456",
		}).Expect().Status(http.StatusBadRequest).JSON().Object().ContainsKey("message")
	})
}
