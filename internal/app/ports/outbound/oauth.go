package outbound

import "github.com/dzemildupljak/simple_hexa/internal/app/domain"

type OAuthRepository interface {
	AuthenticateWithCode(code string) (token string, err error)
	FetchAuthenticatedUser(token string) (user *domain.User, err error)
}
