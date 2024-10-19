package req

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestQueryParam(t *testing.T) {
	req := httptest.NewRequest("GET", "/?key=value", nil)
	r := NewRequest(req)

	expected := "value"
	result := r.QueryParam("key")
	if result != expected {
		t.Fatalf("Expected query param value '%s', got '%s'", expected, result)
	}
}

func TestRequestHeader(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer token123")
	r := NewRequest(req)

	expected := "Bearer token123"
	result := r.Header("Authorization")
	if result != expected {
		t.Fatalf("Expected header value '%s', got '%s'", expected, result)
	}
}

func TestResponseSend(t *testing.T) {
	w := httptest.NewRecorder()
	res := NewResponse(w)

	res.Send("Hello, World!")

	expectedBody := "Hello, World!"
	if w.Body.String() != expectedBody {
		t.Fatalf("Expected body '%s', got '%s'", expectedBody, w.Body.String())
	}

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestResponseStatus(t *testing.T) {
	w := httptest.NewRecorder()
	res := NewResponse(w)

	res.Status(http.StatusCreated).Send("Created")

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}

	expectedBody := "Created"
	if w.Body.String() != expectedBody {
		t.Fatalf("Expected body '%s', got '%s'", expectedBody, w.Body.String())
	}
}

func TestResponseJSON(t *testing.T) {
	w := httptest.NewRecorder()
	res := NewResponse(w)

	res.JSON(map[string]string{"message": "Hello, JSON"})

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	expectedBody := map[string]string{"message": "Hello, JSON"}
	var responseBody map[string]string

	if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
		t.Fatalf("Error unmarshalling response body: %v", err)
	}

	if responseBody["message"] != expectedBody["message"] {
		t.Fatalf("Expected message '%s', got '%s'", expectedBody["message"], responseBody["message"])
	}
}
