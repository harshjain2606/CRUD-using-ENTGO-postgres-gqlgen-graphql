package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// Define a private type for context key to avoid collisions
type contextKey string

const UserIDKey contextKey = "userID"

// Secret key (must match the one used to sign tokens)
var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

// errorResponse sends a JSON error message
func errorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// AuthMiddleware validates JWT and stores userID in context
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		// Allow unauthenticated requests (for public routes)
		if authHeader == "" {
			next.ServeHTTP(w, r)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			errorResponse(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return JwtSecret, nil
		})

		if err != nil || !token.Valid {
			errorResponse(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Extract user ID from claims
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			errorResponse(w, "Invalid token payload", http.StatusUnauthorized)
			return
		}

		userID := int(userIDFloat)

		// Store in context with the SAME key as resolver
		ctx := context.WithValue(r.Context(), UserIDKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
