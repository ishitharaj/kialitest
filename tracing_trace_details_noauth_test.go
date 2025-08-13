package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// Simple test: missing auth context should trigger 500 from getBusiness
func TestTraceDetails_NoAuth_Returns500(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/traces/{traceID}", func(w http.ResponseWriter, r *http.Request) {
		TraceDetails(w, r)
	}).Methods(http.MethodGet)

	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, err := ts.Client().Get(ts.URL + "/api/traces/abc123")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", resp.StatusCode)
	}
}
