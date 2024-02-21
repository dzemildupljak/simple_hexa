package oauth

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/dzemildupljak/simple_hexa/internal/app/domain"
)

type OAuthRepositoryImpl struct {
	clientId           string
	clientSecret       string
	redirectUrl        string
	oAuthTokenEndpoint string
}

func NewOAuthRepository() *OAuthRepositoryImpl {
	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectUrl := os.Getenv("GOOGLE_REDIRECT_URL")
	return &OAuthRepositoryImpl{
		clientId:     clientId,
		clientSecret: clientSecret,
		redirectUrl:  redirectUrl,
	}
}

func (oauth *OAuthRepositoryImpl) FetchAuthenticatedUser(token string) (user *domain.User, err error) {
	return nil, nil
}

// TokenResponse represents the structure of the response from the OAuth provider
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	// Include other fields if needed, like RefreshToken, ExpiresIn, etc.
}

func (oauth *OAuthRepositoryImpl) AuthenticateWithCode(code string) (token string, err error) {
	// Prepare the request data
	data := url.Values{}
	data.Set("client_id", oauth.clientId)
	data.Set("client_secret", oauth.clientSecret)
	//data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", oauth.redirectUrl) // Replace with your redirect URI

	// Make the request
	req, err := http.NewRequest("POST", oauth.oAuthTokenEndpoint, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	err = resp.Body.Close()
	if err != nil {
		return "", err
	}

	// Handle the response
	if resp.StatusCode != http.StatusOK {
		// Handle non-200 responses here
		return "", errors.New("authentication failed")
	}

	// Parse the response body
	var tokenResponse TokenResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}
