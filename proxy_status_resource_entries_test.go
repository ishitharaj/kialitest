package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// Simple test: missing auth context should trigger 500 from getBusiness
func TestConfigDumpResourceEntries_NoAuth_Returns500(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/namespaces/{namespace}/pods/{pod}/config_dump/{resource}", func(w http.ResponseWriter, r *http.Request) {
		ConfigDumpResourceEntries(w, r)
	}).Methods(http.MethodGet)

	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, err := ts.Client().Get(ts.URL + "/api/namespaces/ns/pods/pod/config_dump/routes")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", resp.StatusCode)
	}
}
