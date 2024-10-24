package logger

import (
	"log"
	"net/http"
	"time"
)

func LogInfo(message string) {
	log.Printf("[INFO] %s", message)
}

func LogWarning(message string) {
	log.Printf("[WARNING] %s", message)
}

func LogError(message string) {
	log.Printf("[ERROR] %s", message)
}

func LogRequest(r *http.Request, status int, duration time.Duration) {
	log.Printf("[REQUEST] Method: %s | Path: %s | Status: %d | Duration: %v | User-Agent: %s",
		r.Method, r.URL.Path, status, duration, r.UserAgent())
}
