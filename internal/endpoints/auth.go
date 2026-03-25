package endpoints

import (
	"context"
	"net/http"
	"strings"

	oidc "github.com/coreos/go-oidc/v3/oidc"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/render"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"error": "request does not contain an authorization header"})
			return
		}

		tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)
		provider, err := oidc.NewProvider(r.Context(), "http://localhost:8000/realms/golang")
		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, map[string]string{"error": "error to connect to the provider"})
			return
		}

		verifier := provider.Verifier(&oidc.Config{ClientID: "emailn"})
		_, err = verifier.Verify(r.Context(), tokenStr)
		if err != nil {
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"error": "invalid token"})
			return
		}

		token, _ := jwtgo.Parse(tokenStr, nil)
		clamis := token.Claims.(jwtgo.MapClaims)
		email := clamis["email"]
		ctx := context.WithValue(r.Context(), "email", email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
