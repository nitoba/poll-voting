package module_test

import (
	"testing"

	"github.com/nitoba/poll-voting/pkg/module"
	"github.com/stretchr/testify/assert"
)

type FooModule struct {
	Imports   []module.Module
	Providers []module.Provider
}

func (m *FooModule) Build() {
	m.Providers = module.RevolveProvidersFromImports(m.Imports, m.Providers)
}

func (m *FooModule) GetDependencies() []module.Provider {
	return m.Providers
}

func NewFooModule() *FooModule {
	m := &FooModule{
		Providers: []module.Provider{
			{
				Name: "foo",
				Provide: func(ctn module.Container) (interface{}, error) {
					return "foo", nil
				},
			},
		},
	}
	return m
}

type BarModule struct {
	Imports   []module.Module
	Providers []module.Provider
}

func (m *BarModule) Build() {
	m.Providers = module.RevolveProvidersFromImports(m.Imports, m.Providers)
}

func (m *BarModule) GetDependencies() []module.Provider {
	return m.Providers
}

func NewBarModule() *BarModule {
	m := &BarModule{
		Providers: []module.Provider{
			{
				Name: "far",
				Provide: func(ctn module.Container) (interface{}, error) {
					return "far", nil
				},
			},
		},
	}
	m.Build()
	return m
}

type AppModule struct {
	Imports   []module.Module
	Providers []module.Provider
}

func (m *AppModule) Build() {
	m.Providers = module.RevolveProvidersFromImports(m.Imports, m.Providers)
}

func NewAppModule(options module.NewModule) *AppModule {
	m := &AppModule{
		Imports:   options.Imports,
		Providers: options.Providers,
	}
	m.Build()
	return m
}

func TestModule_Build(t *testing.T) {
	t.Run("should build a module", func(t *testing.T) {
		barModule := NewBarModule()
		fooModule := NewFooModule()
		assert.Len(t, barModule.Providers, 1)
		assert.Len(t, fooModule.Providers, 1)
	})
	t.Run("should build all modules", func(t *testing.T) {
		appModule := NewAppModule(module.NewModule{
			Imports:   []module.Module{NewBarModule(), NewFooModule()},
			Providers: []module.Provider{},
		})
		assert.Len(t, appModule.Providers, 2)
	})

	t.Run("should build all modules and resolve repeated providers", func(t *testing.T) {
		appModule := NewAppModule(module.NewModule{
			Imports: []module.Module{NewBarModule(), NewFooModule()},
			Providers: []module.Provider{
				{
					Name: "far",
					Provide: func(ctn module.Container) (interface{}, error) {
						return "far", nil
					},
				},
			},
		})
		assert.Len(t, appModule.Providers, 2)
	})
}
