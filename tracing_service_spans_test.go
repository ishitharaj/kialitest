package handlers

import (
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

func TestServiceSpans_BadTags_Returns400(t *testing.T) {
	cfg := config.NewConfig()
	config.Set(cfg)
	k8s := kubetest.NewFakeK8sClient()
	business.SetupBusinessLayer(t, k8s, *cfg)

	r := mux.NewRouter()
	r.HandleFunc("/api/namespaces/{namespace}/services/{service}/spans", func(w http.ResponseWriter, r *http.Request) {
		ctx := authentication.SetAuthInfoContext(r.Context(), &api.AuthInfo{Token: "t"})
		ServiceSpans(w, r.WithContext(ctx))
	}).Methods(http.MethodGet)

	ts := httptest.NewServer(r)
	defer ts.Close()

	// malformed JSON in tags param should cause 400
	resp, err := ts.Client().Get(ts.URL + "/api/namespaces/ns/services/svc/spans?tags={bad}")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}
