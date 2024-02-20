package auth_providers

type InvalidCredentialsError struct{}

func (e InvalidCredentialsError) Error() string {
	return "invalid credentials"
}
