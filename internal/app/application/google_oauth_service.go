package application

import (
	"github.com/dzemildupljak/simple_hexa/internal/app/domain"
	"github.com/dzemildupljak/simple_hexa/internal/app/ports/outbound"
)

type OAuthServiceImpl struct {
	oAuthRepository outbound.OAuthRepository
}

func NewOAuthService(oAuthRepo outbound.OAuthRepository) *OAuthServiceImpl {
	return &OAuthServiceImpl{
		oAuthRepository: oAuthRepo,
	}
}

func (oAuthservice *OAuthServiceImpl) ExchangeCodeForToken(code string) (token string, err error) {
	return "", nil
}

func (oAuthservice *OAuthServiceImpl) GetUserInfo(token string) (user *domain.User, err error) {
	return nil, nil
}
