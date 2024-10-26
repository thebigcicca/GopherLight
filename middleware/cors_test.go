package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCORSMiddlewareDefaultOptions(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)

	handler := CORSMiddleware(DefaultCORSOptions)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))

	handler.ServeHTTP(recorder, request)

	assert.Equal(t, recorder.Result().StatusCode, http.StatusOK)
	assert.Equal(t, recorder.Header().Get("Access-Control-Allow-Origin"), "*")
	assert.Equal(t, recorder.Header().Get("Access-Control-Allow-Credentials"), "true")
	assert.Equal(t, recorder.Header().Get("Access-Control-Allow-Headers"), "Content-Type,Authorization")
	assert.Equal(t, recorder.Header().Get("Access-Control-Expose-Headers"), "")
	assert.Equal(t, recorder.Header().Get("Access-Control-Max-Age"), "600")
	assert.Equal(t, recorder.Header().Get("Access-Control-Allow-Methods"), "GET,HEAD,DELETE,OPTIONS,PATCH,POST")
}

func TestCORSMiddlewareOriginNotAllowed(t *testing.T) {
	opts := DefaultCORSOptions
	opts.AllowOrigin = "http://example.com"
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)
	handler := CORSMiddleware(opts)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))

	handler.ServeHTTP(recorder, request)
	assert.Equal(t, recorder.Result().StatusCode, http.StatusForbidden)
}
