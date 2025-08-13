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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd/api"
)

func TestPodLogsDisabledFeature(t *testing.T) {
	cfg := config.NewConfig()
	cfg.KialiFeatureFlags.DisabledFeatures = []string{string(config.FeatureLogView)}
	config.Set(cfg)
	// create namespace and pod to satisfy lookups, though logs are forbidden early
	k8s := kubetest.NewFakeK8sClient(
		&corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "ns"}},
		&corev1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "p", Namespace: "ns"}},
	)
	business.SetupBusinessLayer(t, k8s, *cfg)

	r := mux.NewRouter()
	r.HandleFunc("/api/namespaces/{namespace}/pods/{pod}/logs", func(w http.ResponseWriter, r *http.Request) {
		ctx := authentication.SetAuthInfoContext(r.Context(), &api.AuthInfo{Token: "t"})
		PodLogs(w, r.WithContext(ctx))
	})
	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, err := ts.Client().Get(ts.URL + "/api/namespaces/ns/pods/p/logs")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusForbidden {
		t.Fatalf("expected 403, got %d", resp.StatusCode)
	}
}
