package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testRequestStatusCode(method string, url string, body []byte, code int, t *testing.T) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	_, _ = http.NewRequest(method, url, bytes.NewBuffer(body))
	if w.Code != code {
		t.Errorf("Status should be %d, got %d", code, w.Code)
	}
	return w
}
