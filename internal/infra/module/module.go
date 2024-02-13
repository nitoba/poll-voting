package module

import "github.com/sarulabs/di"

type ModuleConfigOptions struct {
	Dependencies []di.Def
}

func NewModuleConfigOptions(dependencies ...di.Def) *ModuleConfigOptions {
	return &ModuleConfigOptions{
		Dependencies: dependencies,
	}
}

type IModule interface {
	GetDependencies() []di.Def
}

type Module struct {
	ModuleConfigOptions
}

func (module *Module) GetDependencies() []di.Def {
	return module.Dependencies
}

func NewModule(moduleOptions ModuleConfigOptions) *Module {
	return &Module{
		ModuleConfigOptions: *NewModuleConfigOptions(moduleOptions.Dependencies...),
	}
}
