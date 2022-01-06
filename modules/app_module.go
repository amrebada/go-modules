package modules

import (
	"github.com/amrebada/go-template/core"
	user "github.com/amrebada/go-template/modules/user"
)

type AppModule = core.Module

func NewAppModule() *AppModule {
	return core.NewModule().
		SetName("App").
		AddImport(user.New())
}
