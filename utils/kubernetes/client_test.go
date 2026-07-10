package kubernetes

import "testing"

func TestProcessMultiContextConfigKeepsAllContexts(t *testing.T) {
	kubeconfig := []byte(`apiVersion: v1
kind: Config
clusters:
- name: docker-desktop
  cluster:
    server: https://127.0.0.1:54406
    certificate-authority-data: ZHVtbXktY2E=
- name: kind-second
  cluster:
    server: https://127.0.0.1:6443
    certificate-authority-data: ZHVtbXktY2E=
users:
- name: docker-desktop
  user:
    client-certificate-data: ZHVtbXktY2VydA==
    client-key-data: ZHVtbXkta2V5
- name: kind-second
  user:
    client-certificate-data: ZHVtbXktY2VydA==
    client-key-data: ZHVtbXkta2V5
contexts:
- name: docker-desktop
  context:
    cluster: docker-desktop
    user: docker-desktop
- name: kind-second
  context:
    cluster: kind-second
    user: kind-second
current-context: kind-second
`)

	parsed, _, err := ProcessMultiContextConfig(kubeconfig, "")
	if err != nil {
		t.Fatalf("ProcessMultiContextConfig returned error: %v", err)
	}

	if got := len(parsed.Contexts); got != 2 {
		t.Fatalf("expected 2 contexts, got %d", got)
	}

	if _, ok := parsed.Contexts["docker-desktop"]; !ok {
		t.Fatal("expected docker-desktop context to be preserved")
	}

	if _, ok := parsed.Contexts["kind-second"]; !ok {
		t.Fatal("expected kind-second context to be preserved")
	}
}
