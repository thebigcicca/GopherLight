package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCSRFMiddleware(t *testing.T) {

	validToken := GenerateCSRFToken()
	isValidToken := func(token string) bool {
		return token == validToken
	}

	handler := CSRFMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}), isValidToken)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-CSRF-Token", validToken)
	respRec := httptest.NewRecorder()
	handler.ServeHTTP(respRec, req)

	if respRec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", respRec.Code)
	}

	invalidReq := httptest.NewRequest(http.MethodGet, "/", nil)
	invalidReq.Header.Set("X-CSRF-Token", "invalid-token")
	invalidRespRec := httptest.NewRecorder()
	handler.ServeHTTP(invalidRespRec, invalidReq)

	if invalidRespRec.Code != http.StatusForbidden {
		t.Errorf("Expected status 403, got %d", invalidRespRec.Code)
	}
}
