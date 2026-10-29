package client

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jinjie/outsystems-cli/internal/config"
)

// TestDownloadKeyIntegration runs a real TLS server to verify the full path:
// request construction, Bearer header on the wire, octet-stream streaming, and
// Content-Disposition filename parsing — without mocking the transport.
func TestDownloadKeyIntegration(t *testing.T) {
	const payload = "THE-BINARY-BYTES"
	var gotAuth, gotPath string

	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", `attachment; filename="Server.oap"`)
		w.WriteHeader(200)
		io.WriteString(w, payload)
	}))
	defer srv.Close()

	// Point the client at the test server: its hostname is the server host and
	// we reuse the server's TLS-trusting client.
	c := New(config.Profile{Hostname: srv.Listener.Addr().String(), APIKey: "sekret"})
	c.httpClient = srv.Client()

	dl, err := c.DownloadKey(context.Background(), "key42")
	if err != nil {
		t.Fatalf("DownloadKey error: %v", err)
	}
	defer dl.Close()

	if gotAuth != "Bearer sekret" {
		t.Errorf("server saw Authorization %q, want Bearer sekret", gotAuth)
	}
	if gotPath != BasePath+"/downloads/key42/" {
		t.Errorf("server saw path %q, want %s/downloads/key42/", gotPath, BasePath)
	}
	if dl.Filename != "Server.oap" {
		t.Errorf("filename = %q, want Server.oap", dl.Filename)
	}
	data, err := io.ReadAll(dl.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if string(data) != payload {
		t.Errorf("body = %q, want %q", data, payload)
	}
}
