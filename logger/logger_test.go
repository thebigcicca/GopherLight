package logger

import (
	"bytes"
	"log"
	"testing"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetFlags(0)
	f() // Leave it here because otherwise the logger will try to write to a null pointer, resulting in an "invalid memory address or nil pointer dereference" error.
	log.SetOutput(nil)
	return buf.String()
}

func TestLogInfo(t *testing.T) {
	output := captureOutput(func() {
		LogInfo("Testing info message")
	})

	expected := "[INFO] Testing info message\n"
	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestLogWarning(t *testing.T) {
	output := captureOutput(func() {
		LogWarning("Testing warning message")
	})

	expected := "[WARNING] Testing warning message\n"
	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestLogError(t *testing.T) {
	output := captureOutput(func() {
		LogError("Testing error message")
	})

	expected := "[ERROR] Testing error message\n"
	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}
