package controllers_test

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/nitoba/poll-voting/internal/infra/http/server"
	"github.com/nitoba/poll-voting/test"
	"github.com/stretchr/testify/suite"
)

type RegisterVoterControllerTestSuite struct {
	suite.Suite
	e *httpexpect.Expect
}

// Run this function before the all tests
func (s *RegisterVoterControllerTestSuite) SetupSuite() {
	test.BeforeAll()
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
	println("Run after all tests")
}

// Run this function before every test
// func (suite *RegisterVoterControllerTestSuite) SetupTest() {

// }

// Run this function after every test
// func (suite *RegisterVoterControllerTestSuite) TearDownTest() {
// 	println("teardown")
// }

func TestSuit(t *testing.T) {
	// Register the test suite
	suite.Run(t, new(RegisterVoterControllerTestSuite))
}

func (suite *RegisterVoterControllerTestSuite) TestHandle() {
	suite.Run("should return 200", func() {
		suite.e.POST("/auth/register").Expect().Status(http.StatusOK).JSON().Object().HasValue("message", "register voter")
	})
}

func (suite *RegisterVoterControllerTestSuite) TestHandle2() {
	suite.Run("should return 404", func() {
		suite.e.POST("/auth/authenticate").Expect().Status(http.StatusNotFound)
	})
}
