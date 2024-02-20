package user

import (
	"github.com/amrebada/go-modules/modules/user/config"
	"github.com/amrebada/go-modules/modules/user/controllers"
	"github.com/amrebada/go-modules/modules/user/entities"
	core "github.com/amrebada/neon-core"
)

type UserModule = core.Module

func New(_config *config.UserModuleConfiguration) *UserModule {
	return core.NewModule().
		SetName("User").
		SetDescription("User Module for authentication and authorization").
		AddController(controllers.NewAuthController(_config)).
		AddEntity(entities.User{})
}
