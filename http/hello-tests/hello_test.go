package hello

import (
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	Hello(rec, req)

	if rec.Code != 200 {
		t.Errorf("expected 200, got %d", rec.Code)
	}

	msg := "Hello, world!"
	if rec.Body.String() != msg {
		t.Errorf("expected %q, got %q", msg, rec.Body.String())
	}
}
