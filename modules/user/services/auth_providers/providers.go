package auth_providers

import (
	"github.com/amrebada/go-modules/modules/user/config"
	"github.com/amrebada/go-modules/modules/user/entities"
)

type AuthProvider interface {
	Login(identifier string, passwordOrCode string) (token string, user entities.UserPublicData, error error)
	Register()
	Logout()
	ForgotPassword()
	ResetPassword()
	// Invite()
}

func GetProvider(provider config.AuthenticationProvider, _config *config.UserModuleConfiguration) AuthProvider {
	switch provider {
	case config.AuthenticationProviderEmail:
		return &EmailProvider{
			configuration: _config,
		}
	case config.AuthenticationProviderGoogle:
		return &GoogleProvider{
			configuration: _config,
		}
	case config.AuthenticationProviderGithub:
		return &GithubProvider{
			configuration: _config,
		}
	case config.AuthenticationProviderPhone:
		return &PhoneProvider{
			configuration: _config,
		}
	default:
		return nil
	}
}
