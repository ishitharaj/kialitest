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

func TestIstioConfigUpdate_InvalidType(t *testing.T) {
	cfg := config.NewConfig()
	config.Set(cfg)
	k8s := kubetest.NewFakeK8sClient()
	business.SetupBusinessLayer(t, k8s, *cfg)

	r := mux.NewRouter()
	r.HandleFunc("/api/namespaces/{namespace}/istio/{object_type}/{object}", func(w http.ResponseWriter, r *http.Request) {
		ctx := authentication.SetAuthInfoContext(r.Context(), &api.AuthInfo{Token: "t"})
		IstioConfigUpdate(w, r.WithContext(ctx))
	}).Methods(http.MethodPatch)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest(http.MethodPatch, ts.URL+"/api/namespaces/ns/istio/invalidtype/foo", bytes.NewBufferString("[]"))
	if err != nil {
		t.Fatal(err)
	}
	resp, err := ts.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}
