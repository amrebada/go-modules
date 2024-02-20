package auth_providers

import (
	"github.com/amrebada/go-modules/modules/user/config"
	"gorm.io/gorm"
)

type GoogleProvider struct {
	db            *gorm.DB
	configuration *config.UserModuleConfiguration
}

// ForgotPassword implements AuthProvider.
func (*GoogleProvider) ForgotPassword() {
	panic("unimplemented")
}

// Login implements AuthProvider.
func (*GoogleProvider) Login(identifier string, passwordOrCode string) {
	panic("unimplemented")
}

// Logout implements AuthProvider.
func (*GoogleProvider) Logout() {
	panic("unimplemented")
}

// Register implements AuthProvider.
func (*GoogleProvider) Register() {
	panic("unimplemented")
}

// ResetPassword implements AuthProvider.
func (*GoogleProvider) ResetPassword() {
	panic("unimplemented")
}
