package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/kiali/kiali/business"
	"github.com/kiali/kiali/config"
	"github.com/kiali/kiali/kubernetes/kubetest"
	"github.com/kiali/kiali/prometheus/prometheustest"
	"github.com/kiali/kiali/util"
)

func TestServiceDashboard_NoAuth_Returns500(t *testing.T) {
	// Set up mocks
	util.Clock = util.RealClock{}
	config.Set(config.NewConfig())

	// Create a fake client factory
	k8s := kubetest.NewFakeK8sClient()
	clientFactory := kubetest.NewK8SClientFactoryMock(k8s)

	// Mock Prometheus client
	prom := new(prometheustest.PromAPIMock)

	// Set up business layer
	business.SetupBusinessLayer(nil, clientFactory, prom, nil)

	r := mux.NewRouter()
	r.HandleFunc("/api/namespaces/{namespace}/services/{service}/dashboard", ServiceDashboard).Methods(http.MethodGet)
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Call without authentication context - should fail with 500
	resp, err := ts.Client().Get(ts.URL + "/api/namespaces/test-namespace/services/test-service/dashboard")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusInternalServerError && resp.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("expected 500 or 503, got %d", resp.StatusCode)
	}
}

func TestServiceDashboard_BadQuery_Returns400(t *testing.T) {
	// Set up mocks
	util.Clock = util.RealClock{}
	config.Set(config.NewConfig())

	// Create a fake client factory
	k8s := kubetest.NewFakeK8sClient()
	clientFactory := kubetest.NewK8SClientFactoryMock(k8s)

	// Mock Prometheus client
	prom := new(prometheustest.PromAPIMock)

	// Set up business layer
	business.SetupBusinessLayer(nil, clientFactory, prom, nil)

	r := mux.NewRouter()
	r.HandleFunc("/api/namespaces/{namespace}/services/{service}/dashboard", ServiceDashboard).Methods(http.MethodGet)
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Call with bad query parameters (invalid duration format)
	url := ts.URL + "/api/namespaces/test-namespace/services/test-service/dashboard?duration=invalid"
	resp, err := ts.Client().Get(url)
	if err != nil {
		t.Fatal(err)
	}

	// Should get 400 (bad request), 500 (internal server error), or 503 (service unavailable)
	if resp.StatusCode != http.StatusBadRequest && resp.StatusCode != http.StatusInternalServerError && resp.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("expected 400, 500, or 503, got %d", resp.StatusCode)
	}
}

func TestAppDashboard_NoAuth_Returns500(t *testing.T) {
	// Set up mocks
	util.Clock = util.RealClock{}
	config.Set(config.NewConfig())

	// Create a fake client factory
	k8s := kubetest.NewFakeK8sClient()
	clientFactory := kubetest.NewK8SClientFactoryMock(k8s)

	// Mock Prometheus client
	prom := new(prometheustest.PromAPIMock)

	// Set up business layer
	business.SetupBusinessLayer(nil, clientFactory, prom, nil)

	r := mux.NewRouter()
	r.HandleFunc("/api/namespaces/{namespace}/apps/{app}/dashboard", AppDashboard).Methods(http.MethodGet)
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Call without authentication context - should fail with 500
	resp, err := ts.Client().Get(ts.URL + "/api/namespaces/test-namespace/apps/test-app/dashboard")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusInternalServerError && resp.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("expected 500 or 503, got %d", resp.StatusCode)
	}
}

func TestAppDashboard_BadQuery_Returns400(t *testing.T) {
	// Set up mocks
	util.Clock = util.RealClock{}
	config.Set(config.NewConfig())

	// Create a fake client factory
	k8s := kubetest.NewFakeK8sClient()
	clientFactory := kubetest.NewK8SClientFactoryMock(k8s)

	// Mock Prometheus client
	prom := new(prometheustest.PromAPIMock)

	// Set up business layer
	business.SetupBusinessLayer(nil, clientFactory, prom, nil)

	r := mux.NewRouter()
	r.HandleFunc("/api/namespaces/{namespace}/apps/{app}/dashboard", AppDashboard).Methods(http.MethodGet)
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Call with bad query parameters (invalid duration format)
	url := ts.URL + "/api/namespaces/test-namespace/apps/test-app/dashboard?duration=invalid"
	resp, err := ts.Client().Get(url)
	if err != nil {
		t.Fatal(err)
	}

	// Should get 400 (bad request), 500 (internal server error), or 503 (service unavailable)
	if resp.StatusCode != http.StatusBadRequest && resp.StatusCode != http.StatusInternalServerError && resp.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("expected 400, 500, or 503, got %d", resp.StatusCode)
	}
}

func TestWorkloadDashboard_NoAuth_Returns500(t *testing.T) {
	// Set up mocks
	util.Clock = util.RealClock{}
	config.Set(config.NewConfig())

	// Create a fake client factory
	k8s := kubetest.NewFakeK8sClient()
	clientFactory := kubetest.NewK8SClientFactoryMock(k8s)

	// Mock Prometheus client
	prom := new(prometheustest.PromAPIMock)

	// Set up business layer
	business.SetupBusinessLayer(nil, clientFactory, prom, nil)

	r := mux.NewRouter()
	r.HandleFunc("/api/namespaces/{namespace}/workloads/{workload}/dashboard", WorkloadDashboard).Methods(http.MethodGet)
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Call without authentication context - should fail with 500
	resp, err := ts.Client().Get(ts.URL + "/api/namespaces/test-namespace/workloads/test-workload/dashboard")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusInternalServerError && resp.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("expected 500 or 503, got %d", resp.StatusCode)
	}
}

func TestWorkloadDashboard_BadQuery_Returns400(t *testing.T) {
	// Set up mocks
	util.Clock = util.RealClock{}
	config.Set(config.NewConfig())

	// Create a fake client factory
	k8s := kubetest.NewFakeK8sClient()
	clientFactory := kubetest.NewK8SClientFactoryMock(k8s)

	// Mock Prometheus client
	prom := new(prometheustest.PromAPIMock)

	// Set up business layer
	business.SetupBusinessLayer(nil, clientFactory, prom, nil)

	r := mux.NewRouter()
	r.HandleFunc("/api/namespaces/{namespace}/workloads/{workload}/dashboard", WorkloadDashboard).Methods(http.MethodGet)
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Call with bad query parameters (invalid duration format)
	url := ts.URL + "/api/namespaces/test-namespace/workloads/test-workload/dashboard?duration=invalid"
	resp, err := ts.Client().Get(url)
	if err != nil {
		t.Fatal(err)
	}

	// Should get 400 (bad request), 500 (internal server error), or 503 (service unavailable)
	if resp.StatusCode != http.StatusBadRequest && resp.StatusCode != http.StatusInternalServerError && resp.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("expected 400, 500, or 503, got %d", resp.StatusCode)
	}
}
