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
	"github.com/nitoba/poll-voting/internal/infra/http/server"
	"github.com/nitoba/poll-voting/internal/infra/messaging/redis"
	"github.com/nitoba/poll-voting/pkg/di"
	"github.com/nitoba/poll-voting/prisma/db"
	"github.com/nitoba/poll-voting/test"
	"github.com/nitoba/poll-voting/test/factories"
	"github.com/stretchr/testify/suite"
)

type VoteOnPollControllerTestSuite struct {
	suite.Suite
	e          *httpexpect.Expect
	httpModule *http_module.HttpModule
}

// Run this function before the all tests
func (s *VoteOnPollControllerTestSuite) SetupSuite() {
	di.InitContainer()

	httpModule := http_module.NewHttpModule()

	httpModule.Build()

	di.RegisterModuleProviders(httpModule.GetProviders())

	di.BuildDependencies()

	test.SetupDatabase()
	test.SetupRedis()

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
	s.httpModule = httpModule
}

// Run this function after the all tests
func (suite *VoteOnPollControllerTestSuite) TearDownSuite() {
	test.AfterAll()
}

// Run this function before every test
func (suite *VoteOnPollControllerTestSuite) SetupTest() {
	test.TruncateTables()
	test.TruncateRedis()
}

// Run this function after every test
func (suite *VoteOnPollControllerTestSuite) TearDownTest() {
	test.TruncateRedis()
	test.TruncateTables()
}

func TestVoteOnPollControllerSuite(t *testing.T) {
	// Register the test suite
	if os.Getenv("IGNORE_E2E") != "" {
		t.Skip("Ignorando testes E2E")
	}
	suite.Run(t, new(VoteOnPollControllerTestSuite))
}

func (suite *VoteOnPollControllerTestSuite) TestE2EHandle() {
	suite.Run("should return 201 if vote is created", func() {
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

		routeToVote := "/polls/" + pollId.String() + "/vote"

		suite.e.POST(routeToVote).WithHeader("Authorization", "Bearer "+token).WithJSON(map[string]interface{}{
			"option_id": option1.Id.String(),
		}).Expect().Status(http.StatusCreated)

		vote, _ := prisma.GetDB().Votes.FindFirst(db.Votes.PollID.Equals(pollId.String())).Exec(configs.GetConfig().Ctx)
		suite.Equal(option1.Id.String(), vote.PollOptionID)

		count := redis.GetRedis().ZScore(configs.GetConfig().Ctx, pollId.String(), option1.Id.String())
		suite.Equal(float64(1), count.Val())
	})
}

func (suite *VoteOnPollControllerTestSuite) TestE2EHandleWithVoteChanged() {
	suite.Run("should return 201 if vote is changed", func() {
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

		routeToVote := "/polls/" + pollId.String() + "/vote"

		suite.e.POST(routeToVote).WithHeader("Authorization", "Bearer "+token).WithJSON(map[string]interface{}{
			"option_id": option1.Id.String(),
		}).Expect().Status(http.StatusCreated)

		suite.e.POST(routeToVote).WithHeader("Authorization", "Bearer "+token).WithJSON(map[string]interface{}{
			"option_id": option2.Id.String(),
		}).Expect().Status(http.StatusCreated)
		vote, _ := prisma.GetDB().Votes.FindFirst(db.Votes.PollID.Equals(pollId.String())).Exec(configs.GetConfig().Ctx)
		suite.Equal(option2.Id.String(), vote.PollOptionID)
		votes, _ := prisma.GetDB().Votes.FindMany().Exec(configs.GetConfig().Ctx)
		suite.Len(votes, 1)
		count1 := redis.GetRedis().ZScore(configs.GetConfig().Ctx, pollId.String(), option1.Id.String())
		count2 := redis.GetRedis().ZScore(configs.GetConfig().Ctx, pollId.String(), option2.Id.String())
		suite.Equal(float64(0), count1.Val())
		suite.Equal(float64(1), count2.Val())
	})
}
