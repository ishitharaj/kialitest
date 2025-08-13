package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// Simple test: missing auth context causes graph.CheckError to panic and be handled as 500
func TestGraphNamespaces_NoAuth_Returns500(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/graph/namespaces", func(w http.ResponseWriter, r *http.Request) {
		GraphNamespaces(w, r)
	}).Methods(http.MethodGet)

	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, err := ts.Client().Get(ts.URL + "/api/graph/namespaces")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", resp.StatusCode)
	}
}
