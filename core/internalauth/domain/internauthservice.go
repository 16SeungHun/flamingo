package domain

import (
	"context"

	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// InternalAuthService interface for internal oauth clients
// todo necessary?
type InternalAuthService interface {
	GetConfig(TokenEndpointPath string, ClientID string, ClientSecret string, GrantType string) clientcredentials.Config
	GetOauthToken(ctx context.Context, config *clientcredentials.Config) (*oauth2.Token, error)
	GetClaimsFromToken(tokenString string) jwt.MapClaims
}
