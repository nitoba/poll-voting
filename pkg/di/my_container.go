package di

var container DIContainer

type DIContainer interface {
	Get(name string) interface{}
	RegisterModuleProviders(dependencies []Dep)
}

type CTN interface {
	Get(name string) interface{}
}

type Dep struct {
	Name    string
	Provide func(ctn CTN) (interface{}, error)
}

type myContainer struct {
	Dependencies map[string]Dep
}

func (c *myContainer) Get(name string) interface{} {
	if _, existsIn := c.Dependencies[name]; existsIn {
		p, _ := c.Dependencies[name].Provide(c)
		return p
	}
	return nil
}

func GetMyContainer() DIContainer {
	return container
}

func (c *myContainer) RegisterModuleProviders(dependencies []Dep) {
	for _, dep := range dependencies {
		if _, existsIn := c.Dependencies[dep.Name]; !existsIn {
			c.Dependencies[dep.Name] = dep
		}
	}
}

func CreateContainer() {
	container = &myContainer{
		Dependencies: make(map[string]Dep),
	}
}
