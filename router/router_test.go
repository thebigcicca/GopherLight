package router

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/BrunoCiccarino/GopherLight/req"
)

func TestAppRouteGET(t *testing.T) {
	app := NewApp()

	app.Route("GET", "/test", func(req *req.Request, res *req.Response) {
		res.Send("GET Response")
	})

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	app.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	expectedBody := "GET Response"
	if w.Body.String() != expectedBody {
		t.Fatalf("Expected body '%s', got '%s'", expectedBody, w.Body.String())
	}
}

func TestAppRoutePOST(t *testing.T) {
	app := NewApp()

	app.Route("POST", "/test", func(req *req.Request, res *req.Response) {
		res.Send("POST Response")
	})

	req := httptest.NewRequest("POST", "/test", strings.NewReader("test body"))
	w := httptest.NewRecorder()

	app.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	expectedBody := "POST Response"
	if w.Body.String() != expectedBody {
		t.Fatalf("Expected body '%s', got '%s'", expectedBody, w.Body.String())
	}
}

func TestAppRoutePUT(t *testing.T) {
	app := NewApp()

	app.Route("PUT", "/test", func(req *req.Request, res *req.Response) {
		res.Send("PUT Response")
	})

	req := httptest.NewRequest("PUT", "/test", strings.NewReader("test body"))
	w := httptest.NewRecorder()

	app.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	expectedBody := "PUT Response"
	if w.Body.String() != expectedBody {
		t.Fatalf("Expected body '%s', got '%s'", expectedBody, w.Body.String())
	}
}

func TestAppRouteDELETE(t *testing.T) {
	app := NewApp()

	app.Route("DELETE", "/test", func(req *req.Request, res *req.Response) {
		res.Send("DELETE Response")
	})

	req := httptest.NewRequest("DELETE", "/test", nil)
	w := httptest.NewRecorder()

	app.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	expectedBody := "DELETE Response"
	if w.Body.String() != expectedBody {
		t.Fatalf("Expected body '%s', got '%s'", expectedBody, w.Body.String())
	}
}

func TestAppRoutePATCH(t *testing.T) {
	app := NewApp()

	app.Route("PATCH", "/test", func(req *req.Request, res *req.Response) {
		res.Send("PATCH Response")
	})

	req := httptest.NewRequest("PATCH", "/test", strings.NewReader("test body"))
	w := httptest.NewRecorder()

	app.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	expectedBody := "PATCH Response"
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

func TestAppRouteMethodNotAllowed(t *testing.T) {
	app := NewApp()

	app.Route("GET", "/test", func(req *req.Request, res *req.Response) {
		res.Send("GET Response")
	})

	req := httptest.NewRequest("POST", "/test", nil)
	w := httptest.NewRecorder()

	app.ServeHTTP(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("Expected status %d, got %d", http.StatusMethodNotAllowed, w.Code)
	}
}
