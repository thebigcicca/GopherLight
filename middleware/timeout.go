package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/BrunoCiccarino/GopherLight/router"
)

func TimeoutMiddleware(timeout time.Duration) router.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), timeout)
			defer cancel()

			r = r.WithContext(ctx)
			done := make(chan struct{})

			go func() {
				next(w, r)
				close(done)
			}()

			select {
			case <-done:
			case <-ctx.Done():
				http.Error(w, "Request Timeout", http.StatusGatewayTimeout)
			}
		}
	}
}
