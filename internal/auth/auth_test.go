package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_Success(t *testing.T) {
	// Simula un header válido
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey test-key-123")

	key, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if key != "test-key-123" {
		t.Errorf("Expected 'test-key-123', got '%s'", key)
	}
}

func TestGetAPIKey_NoAuthHeader(t *testing.T) {
	// Simula un header faltante
	headers := http.Header{} // Vacío

	key, err := GetAPIKey(headers)
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("Expected ErrNoAuthHeaderIncluded, got %v", err)
	}
	if key != "" {
		t.Errorf("Expected empty string, got '%s'", key)
	}
}
