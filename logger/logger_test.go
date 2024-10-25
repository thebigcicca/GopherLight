package logger

import (
	"bytes"
	"errors"
	"log"
	"testing"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetFlags(0)
	f()
	log.SetOutput(nil)
	return buf.String()
}

func TestLogCriticalError(t *testing.T) {
	output := captureOutput(func() {
		LogCriticalError("Critical error occurred")
	})

	expected := "[CRITICAL] Critical error occurred\n"
	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestCheckCriticalError(t *testing.T) {
	output := captureOutput(func() {
		CheckCriticalError(errors.New("Critical connection error"), "Database connection")
	})

	expected := "[CRITICAL] Database connection: Critical connection error\n"
	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestLogDebug(t *testing.T) {
	output := captureOutput(func() {
		LogDebug("Debugging application")
	})

	expected := "[DEBUG] Debugging application\n"
	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}
