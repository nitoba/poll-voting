package di

import (
	"github.com/nitoba/poll-voting/pkg/module"
	"github.com/sarulabs/di"
)

var ctn di.Container
var builder *di.Builder

func InitContainer() {
	b, _ := di.NewBuilder()
	builder = b
}

func BuildDependencies() {
	ctn = builder.Build()
}

func RegisterModuleProviders(dependencies ...module.Provider) {
	dependency := []di.Def{}

	for _, dep := range dependencies {
		dependency = append(dependency, di.Def{
			Name:  dep.Name,
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return dep.Provide(ctn)
			},
		})
	}

	builder.Add(dependency...)
}

func GetContainer() di.Container {
	return ctn
}
