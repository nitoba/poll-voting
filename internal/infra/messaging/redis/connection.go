package redis

import (
	"context"
	"fmt"

	configs "github.com/nitoba/poll-voting/config"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Conn

func Connect() error {
	logger := configs.GetLogger("redis")
	logger.Info("connecting with redis")
	config := configs.GetConfig()

	addr := fmt.Sprintf("%s:%d", config.RedisHost, config.RedisPort)
	r := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.RedisPassword,
		DB:       0, // use default DB
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			logger.Info("connected with redis")
			rdb = cn
			return nil
		},
	})

	if err := r.Ping(config.Ctx).Err(); err != nil {
		return err
	}
	return nil
}

func Disconnect() error {
	logger := configs.GetLogger("redis")
	logger.Info("disconnecting with redis")
	if err := rdb.Close(); err != nil {
		return err
	}
	return nil
}

func GetRedis() *redis.Conn {
	return rdb
}
