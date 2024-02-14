package di

import (
	"github.com/nitoba/poll-voting/pkg/module"
	"github.com/sarulabs/di/v2"
)

var ctn di.Container
var builder *di.Builder

func InitContainer() {
	if b, err := di.NewBuilder(); err != nil {
		panic(err)
	} else {
		builder = b
	}
}

func BuildDependencies() {
	ctn = builder.Build()
}

func wrapContainerFunc(p module.Provide) func(di.Container) (interface{}, error) {
	return func(ctn di.Container) (interface{}, error) {
		t, err := p(ctn)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
}

func RegisterModuleProviders(providers module.Providers) {
	for _, dep := range providers {
		builder.Add(di.Def{
			Unshared: !dep.IsSingleton,
			Name:     dep.Name,
			Build:    wrapContainerFunc(dep.Provide),
		})
	}
}

func GetContainer() di.Container {
	return ctn
}
