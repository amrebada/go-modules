package user

import core "github.com/amrebada/neon-core"

type UserModule = core.Module

func New() *UserModule {
	return core.NewModule().
		SetName("User").
		SetDescription("User Module for authentication and authorization").
		AddController(NewAuthController()).
		AddEntity(UserEntity{})
}
