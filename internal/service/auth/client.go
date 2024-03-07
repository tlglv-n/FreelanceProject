package auth

import (
	"errors"
	"github.com/go-chi/oauth"
	"net/http"
)

// ValidateUser checks validation of username and password, it returns error if credentials are wrong
func (*Service) ValidateUser(username, password, scope string, r *http.Request) error {
	if username == "user01" && password == "12345" {
		return nil
	}

	return errors.New("wrong user")
}

// ValidateClient checks validations of clientID and secret if credentials are wrong then returns error
func (*Service) ValidateClient(clientID, clientSecret, scope string, r *http.Request) error {
	if clientID == "abcdef" && clientSecret == "12345" {
		return nil
	}

	return errors.New("wrong client")
}

// ValidateCode Validates token ID
func (*Service) ValidateCode(clientID, clientSecret, code, redirectURI string, r *http.Request) (string, error) {
	return "", nil
}

// AddClaims provides additional claims to the token
func (*Service) AddClaims(tokenType oauth.TokenType, credential, tokenID, scope string, r *http.Request) (map[string]string, error) {
	claims := make(map[string]string)
	claims["id"] = "1001"
	claims["data"] = `{"order_date":"2016-12-14","order_id":"9999"}`
	return claims, nil
}

// AddProperties provides additional information to the token response
func (*Service) AddProperties(tokenType oauth.TokenType, credential, tokenID, scope string, r *http.Request) (map[string]string, error) {
	props := make(map[string]string)
	props["name"] = "Gopher"
	return props, nil
}

// ValidateTokenID validates token ID
func (*Service) ValidateTokenID(tokenType oauth.TokenType, credential, tokenID, refreshTokenID string) error {
	return nil
}

// StoreTokenID saves user's client id
func (*Service) StoreTokenID(tokenType oauth.TokenType, credential, tokenID, refreshTokenID string) error {
	return nil
}
