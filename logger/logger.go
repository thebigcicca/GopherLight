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

func LogCriticalError(message string) {
	log.Printf("[CRITICAL] %s", message)
}

func LogFatalError(message string) {
	log.Fatalf("[FATAL] %s", message)
}

func LogDebug(message string) {
	log.Printf("[DEBUG] %s", message)
}

func LogRequest(r *http.Request, status int, duration time.Duration) {
	log.Printf("[REQUEST] Method: %s | Path: %s | Status: %d | Duration: %v | User-Agent: %s",
		r.Method, r.URL.Path, status, duration, r.UserAgent())
}

func CheckCriticalError(err error, context string) {
	if err != nil {
		log.Printf("[CRITICAL] %s: %v", context, err)
	}
}

func CheckFatalError(err error, context string) {
	if err != nil {
		log.Fatalf("[FATAL] %s: %v", context, err)
	}
}
