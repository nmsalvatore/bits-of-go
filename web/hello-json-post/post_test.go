package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestJsonPost(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"message": "Hello, world!"}`))
	rec := httptest.NewRecorder()

	Handler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec.Code)
	}

	contentType := rec.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected application/json, got %s", contentType)
	}

	var data Data
	err := json.Unmarshal(rec.Body.Bytes(), &data)
	if err != nil {
		t.Errorf("failed to unmarshal: %v", err)
	}

	if data.Message != "Hello, world!" {
		t.Errorf(`expected "Hello, world!", got %q`, data.Message)
	}
}
