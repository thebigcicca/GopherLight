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

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
	assert.Equal(t, "*", recorder.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "", recorder.Header().Get("Access-Control-Allow-Credentials"))
	assert.Equal(t, "Content-Type,Authorization", recorder.Header().Get("Access-Control-Allow-Headers"))
	assert.Equal(t, "", recorder.Header().Get("Access-Control-Expose-Headers"))
	assert.Equal(t, "600", recorder.Header().Get("Access-Control-Max-Age"))
	assert.Equal(t, "GET,HEAD,DELETE,OPTIONS,PATCH,POST", recorder.Header().Get("Access-Control-Allow-Methods"))

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
	assert.Equal(t, http.StatusForbidden, recorder.Result().StatusCode)
}
