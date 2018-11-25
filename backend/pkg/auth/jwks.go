package auth

import (
	"context"

	"github.com/go-chi/jwtauth"
	"github.com/lestrrat/go-jwx/jwk"
	"github.com/pkg/errors"
)

type JWKS struct {
	*jwk.Set
}

func New() *JWKS {
	jwks, err := loadCerts()

	if err != nil {
		panic(err)
	}

	return &JWKS{
		jwks,
	}
}

// Converts the first JWK into a public key.
// Auth0 only returns more than one JWK if you've rotated keys.
func (jwks *JWKS) GetFirstValidationKey() (interface{}, error) {
	return jwks.Keys[0].Materialize()
}

// Fetches the Auth0 JWKS from our Auth0 domain.
func loadCerts() (*jwk.Set, error) {
	return jwk.Fetch("https://colinking.auth0.com/.well-known/jwks.json")
}

// Get the Auth0 ID from a request context
func GetAuthId(ctx context.Context) (string, error) {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return "", err
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return "", errors.Errorf("invalid 'sub' key in JWT")
	}

	return sub, nil
}
