package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTimeoutMiddleware(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
	})

	timeoutMiddleware := TimeoutMiddleware(1 * time.Second)(handler)

	req := httptest.NewRequest("GET", "https://httpbin.org/delay/2", nil)
	w := httptest.NewRecorder()

	timeoutMiddleware.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusGatewayTimeout {
		t.Errorf("handler returned wrong status: expected %v, got %v", http.StatusGatewayTimeout, status)
	}
}
