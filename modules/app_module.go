package modules

import (
	"github.com/amrebada/go-modules/core"
	user "github.com/amrebada/go-modules/modules/user"
)

type AppModule = core.Module

func NewAppModule() *AppModule {
	return core.NewModule().
		SetName("App").
		AddImport(user.New())
}
