package auth

import (
	"net/http"
	"github.com/auth0-community/go-auth0"
	"gopkg.in/square/go-jose.v2"
	"fmt"
)

type Auth0 struct {
	secret string
	audience string
	domain string
}

func (auth Auth0) UserAuthenticatedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secret := []byte(auth.secret)
		secretProvider := auth0.NewKeyProvider(secret)
		audience := []string{auth.audience}
		domain := auth.domain
		configuration := auth0.NewConfiguration(secretProvider, audience, domain, jose.HS256)
		validator := auth0.NewValidator(configuration)
		token, err := validator.ValidateRequest(r)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Token is not valid:", token)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func New(secret string, audience string, domain string) *Auth0 {
	return &Auth0 {
		secret: secret,
		audience: audience,
		domain: domain,
	}
}