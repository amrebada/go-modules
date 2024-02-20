package auth_providers

import (
	"github.com/amrebada/go-modules/modules/user/config"
	"gorm.io/gorm"
)

type PhoneProvider struct {
	db            *gorm.DB
	configuration *config.UserModuleConfiguration
}

// ForgotPassword implements AuthProvider.
func (*PhoneProvider) ForgotPassword() {
	panic("unimplemented")
}

// Login implements AuthProvider.
func (*PhoneProvider) Login(identifier string, passwordOrCode string) {
	panic("unimplemented")
}

// Logout implements AuthProvider.
func (*PhoneProvider) Logout() {
	panic("unimplemented")
}

// Register implements AuthProvider.
func (*PhoneProvider) Register() {
	panic("unimplemented")
}

// ResetPassword implements AuthProvider.
func (*PhoneProvider) ResetPassword() {
	panic("unimplemented")
}
