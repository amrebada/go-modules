package modules

import (
	user "github.com/amrebada/go-modules/modules/user"
	core "github.com/amrebada/neon-core"
)

type AppModule = core.Module

func NewAppModule() *AppModule {
	return core.NewModule().
		SetName("App").
		AddImport(user.New())
}
