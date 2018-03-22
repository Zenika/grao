package auth0

import (
	"encoding/json"
	"net/http"
	"github.com/auth0-community/go-auth0"
	"github.com/Zenika/rao/rao-back/log"
	"gopkg.in/square/go-jose.v2"
)

type Response struct {
	Message string `json:"message"`
}

type Auth0 struct {
	jwksUri string
	apiIssuer string
	apiAudience string
}

func (auth Auth0) UserAuthenticatedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if "OPTIONS" == r.Method {
			next.ServeHTTP(w, r)
		}
		client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: auth.jwksUri}, nil)
		audience := []string{auth.apiAudience}
		configuration := auth0.NewConfiguration(client, audience, auth.apiIssuer, jose.RS256)
		validator := auth0.NewValidator(configuration, nil)
		_, err := validator.ValidateRequest(r)
		if err != nil {
			log.Error(err, log.ERROR)
			log.Debug(r.Method + ": Token is not valid or missing token")
			response := Response{
				Message: "Missing or invalid token.",
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func New(jwksUri string, apiIssuer string, apiAudience string) *Auth0 {
	return &Auth0 {
		jwksUri:   jwksUri,
		apiIssuer: apiIssuer,
		apiAudience:   apiAudience,
	}
}
