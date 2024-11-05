package api

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

var jwtSecret = []byte("not_so_secretive_secret_key")

// GenerateToken creates a JWT token with user-specific claims.
func GenerateToken(userID uint, isAdmin bool) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"is_admin": isAdmin,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// validateToken parses and validates a JWT token, returning the claims if valid.
func validateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// AuthMiddleware validates Authorization header tokens and
// adds user information to request context if token is valid.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := validateToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Add claims to the request context
		ctx := r.Context()
		ctx = context.WithValue(ctx, "user_id", claims["user_id"])
		ctx = context.WithValue(ctx, "is_admin", claims["is_admin"])
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// AdminOnlyMiddleware checks if the user has admin privileges.
func AdminOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isAdmin, ok := r.Context().Value("is_admin").(bool)
		if !ok || !isAdmin {
			http.Error(w, "Admin privileges required", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
