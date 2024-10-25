package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

var testSecretKey = []byte("test-secret-key")

func generateTestToken(secretKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
	})
	return token.SignedString(secretKey)
}

func TestAuthMiddleware(t *testing.T) {
	config := JWTConfig{
		SecretKey:     testSecretKey,
		SigningMethod: jwt.SigningMethodHS256,
	}

	token, err := generateTestToken(testSecretKey)
	assert.NoError(t, err)

	reqValid := httptest.NewRequest("GET", "/", nil)
	reqValid.Header.Set("Authorization", "Bearer "+token)
	recorderValid := httptest.NewRecorder()

	handler := NewAuthMiddleware(config)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	handler.ServeHTTP(recorderValid, reqValid)
	assert.Equal(t, http.StatusOK, recorderValid.Code)

	reqInvalid := httptest.NewRequest("GET", "/", nil)
	reqInvalid.Header.Set("Authorization", "Bearer invalid-token")
	recorderInvalid := httptest.NewRecorder()

	handler.ServeHTTP(recorderInvalid, reqInvalid)
	assert.Equal(t, http.StatusUnauthorized, recorderInvalid.Code)

	reqMissing := httptest.NewRequest("GET", "/", nil)
	recorderMissing := httptest.NewRecorder()

	handler.ServeHTTP(recorderMissing, reqMissing)
	assert.Equal(t, http.StatusUnauthorized, recorderMissing.Code)
}
