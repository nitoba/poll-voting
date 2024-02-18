package repositories

import (
	configs "github.com/nitoba/poll-voting/config"
	"github.com/redis/go-redis/v9"
)

type CountingVotingRepositoryRedis struct {
	rdb *redis.Conn
}

func (c *CountingVotingRepositoryRedis) IncrementCountVotesByOptionId(pollId string, optionId string) (int, error) {
	conf := configs.GetConfig()
	cmd := c.rdb.ZIncrBy(conf.Ctx, pollId, 1, optionId)

	if err := cmd.Err(); err != nil {
		return 0, err
	}
	return int(cmd.Val()), nil
}
func (c *CountingVotingRepositoryRedis) DecrementCountVotesByOptionId(pollId string, optionId string) (int, error) {
	conf := configs.GetConfig()
	cmd := c.rdb.ZIncrBy(conf.Ctx, pollId, -1, optionId)

	if err := cmd.Err(); err != nil {
		return 0, err
	}
	return int(cmd.Val()), nil
}
func (c *CountingVotingRepositoryRedis) CountVotesByOptionId(pollId string, optionId string) (int, error) {
	conf := configs.GetConfig()
	cmd := c.rdb.ZScore(conf.Ctx, pollId, optionId)

	if err := cmd.Err(); err != nil {
		return 0, err
	}

	return int(cmd.Val()), nil
}

func (c *CountingVotingRepositoryRedis) CountVotes(pollId string) (int, error) {
	totalOfVotes := 0
	conf := configs.GetConfig()
	cmd := c.rdb.ZRangeWithScores(conf.Ctx, pollId, 0, -1)
	if err := cmd.Err(); err != nil {
		return 0, err
	}

	for _, v := range cmd.Val() {
		totalOfVotes += int(v.Score)
	}

	return totalOfVotes, nil
}

func NewCountingVotingRepositoryRedis(rdb *redis.Conn) *CountingVotingRepositoryRedis {
	return &CountingVotingRepositoryRedis{rdb: rdb}
}
