package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/kiali/kiali/business"
	"github.com/kiali/kiali/business/authentication"
	"github.com/kiali/kiali/config"
	"github.com/kiali/kiali/kubernetes/kubetest"
	"k8s.io/client-go/tools/clientcmd/api"
)

func TestIstioConfigCreate_InvalidType(t *testing.T) {
	cfg := config.NewConfig()
	config.Set(cfg)
	k8s := kubetest.NewFakeK8sClient()
	business.SetupBusinessLayer(t, k8s, *cfg)

	r := mux.NewRouter()
	r.HandleFunc("/api/namespaces/{namespace}/istio/{object_type}", func(w http.ResponseWriter, r *http.Request) {
		ctx := authentication.SetAuthInfoContext(r.Context(), &api.AuthInfo{Token: "t"})
		IstioConfigCreate(w, r.WithContext(ctx))
	}).Methods(http.MethodPost)

	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, err := ts.Client().Post(ts.URL+"/api/namespaces/ns/istio/invalidtype", "application/json", bytes.NewBufferString("{}"))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}

func TestIstioConfigCreate_VirtualService_MinimalBody(t *testing.T) {
	cfg := config.NewConfig()
	config.Set(cfg)
	k8s := kubetest.NewFakeK8sClient()
	business.SetupBusinessLayer(t, k8s, *cfg)

	r := mux.NewRouter()
	r.HandleFunc("/api/namespaces/{namespace}/istio/{object_type}", func(w http.ResponseWriter, r *http.Request) {
		ctx := authentication.SetAuthInfoContext(r.Context(), &api.AuthInfo{Token: "t"})
		IstioConfigCreate(w, r.WithContext(ctx))
	}).Methods(http.MethodPost)

	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, err := ts.Client().Post(ts.URL+"/api/namespaces/ns/istio/virtualservices", "application/json", bytes.NewBufferString("{}"))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusBadRequest && resp.StatusCode != http.StatusInternalServerError && resp.StatusCode != http.StatusForbidden && resp.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("unexpected status: %d", resp.StatusCode)
	}
}
