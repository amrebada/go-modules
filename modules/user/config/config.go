package config

type AuthenticationProvider string

const (
	AuthenticationProviderEmail  = "email"
	AuthenticationProviderPhone  = "phone"
	AuthenticationProviderGoogle = "google"
	AuthenticationProviderGithub = "github"
)

type AuthenticationConfiguration struct {
	Providers                      []AuthenticationProvider `json:"providers"`
	EnableRegistration             bool                     `json:"enableRegistration"`
	EnableInvitation               bool                     `json:"enableInvitation"`
	EnableResetPassword            bool                     `json:"enableResetPassword"`
	EnableForgotPassword           bool                     `json:"enableForgotPassword"`
	Stateless                      bool                     `json:"stateless"`
	JwtSecret                      string                   `json:"jwtSecret"`
	JwtExpirationDurationInSeconds int                      `json:"jwtExpirationDurationInSeconds"`
}

type UserModuleConfiguration struct {
	Authentication AuthenticationConfiguration `json:"authentication"`
}
