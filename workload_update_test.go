package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// Simple test: missing auth context should trigger 500 from getBusiness
func TestWorkloadUpdate_NoAuth_Returns500(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/namespaces/{namespace}/workloads/{workload}", func(w http.ResponseWriter, r *http.Request) {
		WorkloadUpdate(w, r)
	}).Methods(http.MethodPatch)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest(http.MethodPatch, ts.URL+"/api/namespaces/ns/workloads/wl", bytes.NewBufferString("[]"))
	if err != nil {
		t.Fatal(err)
	}
	resp, err := ts.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", resp.StatusCode)
	}
}
