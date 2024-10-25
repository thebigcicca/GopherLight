package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type JWTConfig struct {
	SecretKey      []byte
	SigningMethod  jwt.SigningMethod
	ErrorHandler   func(w http.ResponseWriter, message string, code int)
	TokenExtractor func(r *http.Request) (string, error)
}

func DefaultErrorHandler(w http.ResponseWriter, message string, code int) {
	http.Error(w, message, code)
}

func DefaultTokenExtractor(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header missing")
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return "", errors.New("invalid token format")
	}
	return tokenString, nil
}

func NewAuthMiddleware(config JWTConfig) func(http.HandlerFunc) http.HandlerFunc {
	if config.ErrorHandler == nil {
		config.ErrorHandler = DefaultErrorHandler
	}
	if config.TokenExtractor == nil {
		config.TokenExtractor = DefaultTokenExtractor
	}
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			tokenString, err := config.TokenExtractor(r)
			if err != nil {
				config.ErrorHandler(w, err.Error(), http.StatusUnauthorized)
				return
			}

			token, err := validateJWT(tokenString, config)
			if err != nil || !token.Valid {
				config.ErrorHandler(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			next(w, r)
		}
	}
}

func validateJWT(tokenString string, config JWTConfig) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return config.SecretKey, nil
	})
}
