package modules

import (
	user "github.com/amrebada/go-modules/modules/user"
	userConfig "github.com/amrebada/go-modules/modules/user/config"
	core "github.com/amrebada/neon-core"
)

type AppModule = core.Module

func NewAppModule() *AppModule {

	return core.NewModule().
		SetName("App").
		AddImport(user.New(&userConfig.UserModuleConfiguration{
			Authentication: userConfig.AuthenticationConfiguration{
				Providers: []userConfig.AuthenticationProvider{
					userConfig.AuthenticationProviderEmail,
					userConfig.AuthenticationProviderGoogle,
				},
			},
		}))
}
