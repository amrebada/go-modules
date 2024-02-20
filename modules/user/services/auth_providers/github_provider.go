package auth_providers

import (
	"github.com/amrebada/go-modules/modules/user/config"
	"gorm.io/gorm"
)

type GithubProvider struct {
	db            *gorm.DB
	configuration *config.UserModuleConfiguration
}

// ForgotPassword implements AuthProvider.
func (*GithubProvider) ForgotPassword() {
	panic("unimplemented")
}

// Login implements AuthProvider.
func (*GithubProvider) Login(identifier string, passwordOrCode string) {
	panic("unimplemented")
}

// Logout implements AuthProvider.
func (*GithubProvider) Logout() {
	panic("unimplemented")
}

// Register implements AuthProvider.
func (*GithubProvider) Register() {
	panic("unimplemented")
}

// ResetPassword implements AuthProvider.
func (*GithubProvider) ResetPassword() {
	panic("unimplemented")
}
