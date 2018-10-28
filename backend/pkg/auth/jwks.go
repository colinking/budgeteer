package auth

import (
	"github.com/lestrrat/go-jwx/jwk"
)

type JWKS struct {
	*jwk.Set
}

func New() *JWKS {
	jwks, err := loadCerts()

	if err != nil {
		panic(err)
	}

	return &JWKS {
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
