package auth_providers

import (
	"github.com/amrebada/go-modules/modules/user/config"
	"github.com/amrebada/go-modules/modules/user/entities"
	"github.com/amrebada/go-modules/modules/user/utils"
	"gorm.io/gorm"
)

type EmailProvider struct {
	db            *gorm.DB
	configuration *config.UserModuleConfiguration
}

// ForgotPassword implements AuthProvider.
func (*EmailProvider) ForgotPassword() {
	panic("unimplemented")
}

// Login implements AuthProvider.
func (e *EmailProvider) Login(identifier string, passwordOrCode string) (token string, user map[string]any, error error) {
	userRepo := entities.NewUserRepository(e.db)
	userRepo.Email = identifier
	userRepo.GetByEmail()
	if userRepo.ID == "" {
		return "", map[string]any{}, InvalidCredentialsError{}
	}
	if !utils.ComparePassword(userRepo.Password, passwordOrCode) {
		return "", map[string]any{}, InvalidCredentialsError{}
	}

	//TODO: check stateless or stateful token generation
	token, err := utils.GenerateToken(userRepo.GetPublicData(), e.configuration)
	if err != nil {
		return "", map[string]any{}, err
	}

	return token, userRepo.GetPublicData(), nil

}

// Logout implements AuthProvider.
func (*EmailProvider) Logout() {
	panic("unimplemented")
}

// Register implements AuthProvider.
func (*EmailProvider) Register() {
	panic("unimplemented")
}

// ResetPassword implements AuthProvider.
func (*EmailProvider) ResetPassword() {
	panic("unimplemented")
}
