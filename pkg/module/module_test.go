package module_test

import (
	"testing"

	"github.com/nitoba/poll-voting/pkg/module"
	"github.com/sarulabs/di/v2"
	"github.com/stretchr/testify/assert"
)

type FooModule struct {
	module.Module
}

func NewFooModule() *FooModule {
	m := &FooModule{
		Module: module.Module{
			Providers: module.Providers{
				{
					Name: "foo",
					Provide: func(ctn di.Container) (interface{}, error) {
						return "foo", nil
					},
				},
			},
		},
	}
	m.Build()
	return m
}

type BarModule struct {
	module.Module
}

func NewBarModule() *BarModule {
	m := &BarModule{
		Module: module.Module{
			Imports: module.Imports{},
			Providers: module.Providers{
				{
					Name: "bar",
					Provide: func(ctn di.Container) (interface{}, error) {
						return "bar", nil
					},
				},
			},
		},
	}
	m.Build()
	return m
}

type AppModule struct {
	module.Module
}

func NewAppModule(module module.Module) *AppModule {
	m := &AppModule{module}
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
		appModule := NewAppModule(module.Module{
			Imports: module.Imports{NewBarModule(), NewFooModule()},
		})
		p1, _ := appModule.Providers[0].Provide(nil)
		p2, _ := appModule.Providers[1].Provide(nil)
		assert.Len(t, appModule.Providers, 2)
		assert.Equal(t, "bar", appModule.Providers[0].Name)
		assert.Equal(t, "foo", appModule.Providers[1].Name)
		assert.Equal(t, "bar", p1)
		assert.Equal(t, "foo", p2)
	})

	t.Run("should build all modules and resolve repeated providers", func(t *testing.T) {
		barModule := NewBarModule()
		fooModule := NewFooModule()
		appModule := NewAppModule(module.Module{
			Imports: module.Imports{barModule, fooModule},
			Providers: module.Providers{
				{
					Name: "bar",
					Provide: func(ctn di.Container) (interface{}, error) {
						return "bar", nil
					},
				},
			},
		})
		assert.Len(t, appModule.Providers, 2)
		assert.Len(t, barModule.Providers, 1)
		assert.Len(t, fooModule.Providers, 1)
		assert.Equal(t, "bar", appModule.Providers[0].Name)
		assert.Equal(t, "foo", appModule.Providers[1].Name)
	})
}
