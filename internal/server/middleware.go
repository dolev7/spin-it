package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dolev7/spin-it/pkg/auth"
)

// AuthMiddleware ensures the user is authenticated before accessing protected routes
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "Missing token"})
			return
		}

		// Extract token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := auth.ValidateJWT(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "Invalid token"})
			return
		}

		// Store email in request header for use in handlers
		r.Header.Set("X-User-Email", claims.Email)

		next.ServeHTTP(w, r)
	})
}
