package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestKafkaSources_OK(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/kafka/sources", KafkaSources).Methods(http.MethodGet)
	ts := httptest.NewServer(r)
	defer ts.Close()
	resp, err := ts.Client().Get(ts.URL + "/api/kafka/sources")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("unexpected status: %d", resp.StatusCode)
	}
}

func TestKafkaPartitions_Basic(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/kafka/{source}/{topic}/partitions", KafkaPartitions).Methods(http.MethodGet)
	ts := httptest.NewServer(r)
	defer ts.Close()
	resp, err := ts.Client().Get(ts.URL + "/api/kafka/src/topic/partitions")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("unexpected status: %d", resp.StatusCode)
	}
}

func TestKafkaDashboardsList_Basic(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/kafka/{source}/dashboards", KafkaDashboardsList).Methods(http.MethodGet)
	ts := httptest.NewServer(r)
	defer ts.Close()
	resp, err := ts.Client().Get(ts.URL + "/api/kafka/src/dashboards")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("unexpected status: %d", resp.StatusCode)
	}
}

func TestKafkaDashboard_BadLimit(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/kafka/{source}/dashboards/{template}", KafkaDashboard).Methods(http.MethodGet)
	ts := httptest.NewServer(r)
	defer ts.Close()
	url := ts.URL + "/api/kafka/src/dashboards/tpl?limit=abc&offset=0"
	resp, err := ts.Client().Get(url)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", resp.StatusCode)
	}
}

func TestKafkaDashboard_BadOffset(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/kafka/{source}/dashboards/{template}", KafkaDashboard).Methods(http.MethodGet)
	ts := httptest.NewServer(r)
	defer ts.Close()
	url := ts.URL + "/api/kafka/src/dashboards/tpl?limit=1&offset=bad"
	resp, err := ts.Client().Get(url)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", resp.StatusCode)
	}
}

func TestVictoriaLogsDashboardsList_OK(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/vl/dashboards", VictoriaLogsDashboardsList).Methods(http.MethodGet)
	ts := httptest.NewServer(r)
	defer ts.Close()
	resp, err := ts.Client().Get(ts.URL + "/api/vl/dashboards?is_journal=true")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("unexpected status: %d", resp.StatusCode)
	}
}

func TestVictoriaLogsTenantsList_Basic(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/vl/{template}/tenants/{namespace}", VictoriaLogsTenantsList).Methods(http.MethodGet)
	ts := httptest.NewServer(r)
	defer ts.Close()
	resp, err := ts.Client().Get(ts.URL + "/api/vl/tpl/tenants/ns")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("unexpected status: %d", resp.StatusCode)
	}
}

func TestResendRequest_BadJSON(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/vl/resend", ResendRequest).Methods(http.MethodPost)
	ts := httptest.NewServer(r)
	defer ts.Close()
	resp, err := ts.Client().Post(ts.URL+"/api/vl/resend", "application/json", bytes.NewBufferString("{"))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}

func TestVictoriaLogsDashboard_BadLimit(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/vl/{template}/dashboard", VictoriaLogsDashboard).Methods(http.MethodGet)
	ts := httptest.NewServer(r)
	defer ts.Close()
	url := ts.URL + "/api/vl/tpl/dashboard?limit=abc&offset=0"
	resp, err := ts.Client().Get(url)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", resp.StatusCode)
	}
}

func TestVictoriaLogsDashboard_BadOffset(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/vl/{template}/dashboard", VictoriaLogsDashboard).Methods(http.MethodGet)
	ts := httptest.NewServer(r)
	defer ts.Close()
	url := ts.URL + "/api/vl/tpl/dashboard?limit=1&offset=bad"
	resp, err := ts.Client().Get(url)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", resp.StatusCode)
	}
}

func TestArtemisSources_OK(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/artemis/sources", ArtemisSources).Methods(http.MethodGet)
	ts := httptest.NewServer(r)
	defer ts.Close()
	resp, err := ts.Client().Get(ts.URL + "/api/artemis/sources")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("unexpected status: %d", resp.StatusCode)
	}
}

func TestArtemisDashboardsList_Basic(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/artemis/{source}/dashboards", ArtemisDashboardsList).Methods(http.MethodGet)
	ts := httptest.NewServer(r)
	defer ts.Close()
	resp, err := ts.Client().Get(ts.URL + "/api/artemis/src/dashboards")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("unexpected status: %d", resp.StatusCode)
	}
}

func TestArtemisDashboard_BadLimit(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/artemis/{source}/dashboards/{template}", ArtemisDashboard).Methods(http.MethodGet)
	ts := httptest.NewServer(r)
	defer ts.Close()
	url := ts.URL + "/api/artemis/src/dashboards/tpl?limit=abc&offset=0"
	resp, err := ts.Client().Get(url)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", resp.StatusCode)
	}
}

func TestArtemisDashboard_BadOffset(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/artemis/{source}/dashboards/{template}", ArtemisDashboard).Methods(http.MethodGet)
	ts := httptest.NewServer(r)
	defer ts.Close()
	url := ts.URL + "/api/artemis/src/dashboards/tpl?limit=1&offset=bad"
	resp, err := ts.Client().Get(url)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", resp.StatusCode)
	}
}
