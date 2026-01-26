package hello

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", "", "Hello, world!"},
		{"with name", "?name=Gopher", "Hello, Gopher!"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/"+tc.input, nil)

			Hello(rec, req)

			if rec.Code != http.StatusOK {
				t.Errorf("expected 200, got %d", rec.Code)
			}

			contentType := rec.Header().Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf(`expected "application/json", got %q`, contentType)
			}

			var got Data
			err := json.Unmarshal(rec.Body.Bytes(), &got)
			if err != nil {
				t.Fatalf("failed to unmarshal: %v", err)
			}

			if got.Message != tc.expected {
				t.Fatalf(`expected %q, got %q`, tc.expected, got.Message)
			}
		})
	}
}
