package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// Simple test: missing auth context should trigger 500 from getBusiness
func TestServiceDetails_NoAuth_Returns500(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/namespaces/{namespace}/services/{service}", func(w http.ResponseWriter, r *http.Request) {
		ServiceDetails(w, r)
	}).Methods(http.MethodGet)

	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, err := ts.Client().Get(ts.URL + "/api/namespaces/ns/services/svc")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", resp.StatusCode)
	}
}
