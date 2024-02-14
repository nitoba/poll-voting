package module

type Container interface {
	Get(name string) interface{}
}

type Module interface {
	GetDependencies() []Provider
	Build()
}

type Provider struct {
	Name    string
	Provide func(ctn Container) (interface{}, error)
}

type NewModule struct {
	Imports   []Module
	Providers []Provider
}
