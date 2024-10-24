package middleware

import (
	"net/http"
	"time"

	"github.com/BrunoCiccarino/GopherLight/logger"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logger.LogInfo("Started " + r.Method + " " + r.URL.Path)

		next(w, r)

		duration := time.Since(start)
		logger.LogRequest(r, http.StatusOK, duration)
		logger.LogInfo("Completed " + r.URL.Path + " in " + duration.String())
	}
}
