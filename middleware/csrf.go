package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
)

func CSRFMiddleware(next http.HandlerFunc, isValidToken func(string) bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		csrfToken := r.Header.Get("X-CSRF-Token")
		if csrfToken == "" || !isValidToken(csrfToken) {
			http.Error(w, "Invalid CSRF token", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}

func GenerateCSRFToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
