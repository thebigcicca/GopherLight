package middleware

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoggingMiddleware(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	loggingMiddleware := LoggingMiddleware(handler)

	req := httptest.NewRequest("GET", "https://httpbin.org/get", nil)
	w := httptest.NewRecorder()

	loggingMiddleware.ServeHTTP(w, req)

	logOutput := buf.String()

	if !strings.Contains(logOutput, "Started GET /get") {
		t.Errorf("Expected log containing 'Started GET /get', but got: %s", logOutput)
	}

	if !strings.Contains(logOutput, "Completed /get") {
		t.Errorf("Expected log containing 'Completed /get', but got: %s", logOutput)
	}

	if w.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %v", w.Code)
	}
}
