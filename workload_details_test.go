package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/kiali/kiali/util"
)

// Simple test: missing auth context should trigger 500 from getBusiness
func TestWorkloadDetails_NoAuth_Returns500(t *testing.T) {
	// ensure util.Clock is non-nil to avoid panic in baseExtract
	util.Clock = util.RealClock{}

	r := mux.NewRouter()
	r.HandleFunc("/api/namespaces/{namespace}/workloads/{workload}", func(w http.ResponseWriter, r *http.Request) {
		WorkloadDetails(w, r)
	}).Methods(http.MethodGet)

	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, err := ts.Client().Get(ts.URL + "/api/namespaces/ns/workloads/wl")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", resp.StatusCode)
	}
}
