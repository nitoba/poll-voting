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
	"github.com/nitoba/poll-voting/internal/infra/database/redis"
	http_module "github.com/nitoba/poll-voting/internal/infra/http"
	"github.com/nitoba/poll-voting/internal/infra/http/rest"
	"github.com/nitoba/poll-voting/pkg/di"
	"github.com/nitoba/poll-voting/prisma/db"
	"github.com/nitoba/poll-voting/test"
	"github.com/nitoba/poll-voting/test/factories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GetPollByIdControllerTestSuite struct {
	suite.Suite
	e          *httpexpect.Expect
	httpModule *http_module.HttpModule
}

// Run this function before the all tests
func (s *GetPollByIdControllerTestSuite) SetupSuite() {
	di.InitContainer()

	httpModule := http_module.NewHttpModule()

	httpModule.Build()

	di.RegisterModuleProviders(httpModule.GetProviders())

	di.BuildDependencies()

	test.SetupDatabase()
	test.SetupRedis()

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
func (suite *GetPollByIdControllerTestSuite) TearDownSuite() {
	test.AfterAll()
}

// Run this function before every test
func (suite *GetPollByIdControllerTestSuite) SetupTest() {
	test.TruncateTables()
	test.TruncateRedis()
}

// Run this function after every test
func (suite *GetPollByIdControllerTestSuite) TearDownTest() {
	test.TruncateTables()
	test.TruncateRedis()
}

func TestGetPollByIdControllerSuite(t *testing.T) {
	// Register the test suite
	if os.Getenv("IGNORE_E2E") != "" {
		t.Skip("Ignorando testes E2E")
	}
	suite.Run(t, new(GetPollByIdControllerTestSuite))
}

func (suite *GetPollByIdControllerTestSuite) TestE2EHandle() {
	suite.Run("should return a poll by id", func() {
		userID := core.NewUniqueEntityId()

		factories.MakePrismaVoter(factories.OptionalVoterParams{
			Id: &userID,
		})

		pollId := core.NewUniqueEntityId()

		option1 := factories.MakePoolOption(factories.OptionalPollOptionParams{
			Title: "Option 1",
		})

		option2 := factories.MakePoolOption(factories.OptionalPollOptionParams{
			Title: "Option 2",
		})

		factories.MakePrismaPoll(factories.OptionalPollParams{
			Id:      &pollId,
			Title:   "Poll example",
			OwnerId: &userID,
			Options: []*entities.PollOption{
				option1,
				option2,
			},
		})

		token := di.GetContainer().Get("encrypter").(*cryptography.JWTEncrypter).Encrypt(map[string]interface{}{
			"sub": userID.String(),
		})

		pollResponse := map[string]interface{}{
			"id":    pollId.String(),
			"title": "Poll example",
			"total": 0,
			"options": []map[string]interface{}{
				{
					"id":    option1.Id.String(),
					"title": "Option 1",
				},
				{
					"id":    option2.Id.String(),
					"title": "Option 2",
				},
			},
		}

		suite.e.GET("/polls/"+pollId.String()).WithCookie("auth", token).Expect().Status(http.StatusOK).JSON().Object().IsEqual(pollResponse)

		poll, _ := prisma.GetDB().Poll.FindFirst(db.Poll.Title.Equals("Poll example")).Exec(configs.GetConfig().Ctx)
		assert.NotEmpty(suite.T(), poll.ID)
		assert.Equal(suite.T(), "Poll example", poll.Title)
	})
}

func (suite *GetPollByIdControllerTestSuite) TestE2EHandleWithTotalVotes() {
	suite.Run("should return a poll by id with total votes", func() {
		userID := core.NewUniqueEntityId()

		factories.MakePrismaVoter(factories.OptionalVoterParams{
			Id: &userID,
		})

		pollId := core.NewUniqueEntityId()

		option1 := factories.MakePoolOption(factories.OptionalPollOptionParams{
			Title: "Option 1",
		})

		option2 := factories.MakePoolOption(factories.OptionalPollOptionParams{
			Title: "Option 2",
		})

		factories.MakePrismaPoll(factories.OptionalPollParams{
			Id:      &pollId,
			Title:   "Poll example",
			OwnerId: &userID,
			Options: []*entities.PollOption{
				option1,
				option2,
			},
		})

		redis.GetRedis().ZIncrBy(configs.GetConfig().Ctx, pollId.String(), 1, option1.Id.String())
		redis.GetRedis().ZIncrBy(configs.GetConfig().Ctx, pollId.String(), 1, option1.Id.String())
		redis.GetRedis().ZIncrBy(configs.GetConfig().Ctx, pollId.String(), 1, option2.Id.String())
		redis.GetRedis().ZIncrBy(configs.GetConfig().Ctx, pollId.String(), 1, option2.Id.String())
		redis.GetRedis().ZIncrBy(configs.GetConfig().Ctx, pollId.String(), 1, option2.Id.String())

		token := di.GetContainer().Get("encrypter").(*cryptography.JWTEncrypter).Encrypt(map[string]interface{}{
			"sub": userID.String(),
		})

		pollResponse := map[string]interface{}{
			"id":    pollId.String(),
			"title": "Poll example",
			"total": 5,
			"options": []map[string]interface{}{
				{
					"id":    option1.Id.String(),
					"title": "Option 1",
				},
				{
					"id":    option2.Id.String(),
					"title": "Option 2",
				},
			},
		}

		suite.e.GET("/polls/"+pollId.String()).WithCookie("auth", token).Expect().Status(http.StatusOK).JSON().Object().IsEqual(pollResponse)

		poll, _ := prisma.GetDB().Poll.FindFirst(db.Poll.Title.Equals("Poll example")).Exec(configs.GetConfig().Ctx)
		assert.NotEmpty(suite.T(), poll.ID)
		assert.Equal(suite.T(), "Poll example", poll.Title)
	})
}
