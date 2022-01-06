package user

import "github.com/amrebada/go-template/core"

type UserModule = core.Module

func New() *UserModule {
	return core.NewModule().
		SetName("User").
		AddController(NewAuthController()).
		AddEntity(UserEntity{})
}
