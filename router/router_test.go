package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BrunoCiccarino/express-go/req"
)

func TestAppRouteValid(t *testing.T) {
	app := NewApp()

	app.Route("/test", func(req *req.Request, res *req.Response) {
		res.Send("Test Response")
	})

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	app.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	expectedBody := "Test Response"
	if w.Body.String() != expectedBody {
		t.Fatalf("Expected body '%s', got '%s'", expectedBody, w.Body.String())
	}
}

func TestAppRouteNotFound(t *testing.T) {
	app := NewApp()

	req := httptest.NewRequest("GET", "/unknown", nil)
	w := httptest.NewRecorder()

	app.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("Expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}
