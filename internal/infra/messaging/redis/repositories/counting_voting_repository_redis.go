package repositories

import "github.com/redis/go-redis/v9"

type CountingVotingRepositoryRedis struct {
	rdb *redis.Conn
}

func (c *CountingVotingRepositoryRedis) IncrementCountVotesByOptionId(pollId string, optionId string) (int, error) {
	return 0, nil
}
func (c *CountingVotingRepositoryRedis) DecrementCountVotesByOptionId(pollId string, optionId string) (int, error) {
	return 0, nil
}
func (c *CountingVotingRepositoryRedis) CountVotesByOptionId(pollId string, optionId string) (int, error) {
	return 0, nil
}

func NewCountingVotingRepositoryRedis(rdb *redis.Conn) *CountingVotingRepositoryRedis {
	return &CountingVotingRepositoryRedis{rdb: rdb}
}
