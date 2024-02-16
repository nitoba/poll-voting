package messaging

import (
	"github.com/nitoba/poll-voting/internal/infra/messaging/redis"
	"github.com/nitoba/poll-voting/pkg/module"
)

type MessagingModule struct {
	module.Module
}

func NewMessagingModule() *MessagingModule {
	m := &MessagingModule{
		Module: module.Module{
			Providers: module.Providers{
				{
					Name: "redis",
					Provide: func(ctn module.Container) (interface{}, error) {
						return redis.GetRedis(), nil
					},
				},
			},
		},
	}
	m.Build()
	return m
}
