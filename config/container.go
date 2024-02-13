package configs

import "github.com/sarulabs/di"

var ctn di.Container
var builder *di.Builder

func InitContainer() {
	b, _ := di.NewBuilder()
	builder = b
}

func BuildDependencies() {
	ctn = builder.Build()
}

func RegisterDependency(dependency ...di.Def) {
	builder.Add(dependency...)
}

func GetContainer() di.Container {
	return ctn
}
