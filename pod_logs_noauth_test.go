package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/kiali/kiali/config"
)

// Simple test: feature enabled but missing auth context should trigger 500 from getBusiness
func TestPodLogs_NoAuth_Returns500(t *testing.T) {
	cfg := config.NewConfig()
	// Ensure logs-tab feature is enabled (not disabled)
	cfg.KialiFeatureFlags.DisabledFeatures = []string{}
	config.Set(cfg)

	r := mux.NewRouter()
	r.HandleFunc("/api/namespaces/{namespace}/pods/{pod}/logs", func(w http.ResponseWriter, r *http.Request) {
		PodLogs(w, r)
	}).Methods(http.MethodGet)

	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, err := ts.Client().Get(ts.URL + "/api/namespaces/ns/pods/p/logs")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", resp.StatusCode)
	}
}
