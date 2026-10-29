package client

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
)

func binResp(status int, body string, headers map[string]string) *http.Response {
	h := make(http.Header)
	for k, v := range headers {
		h.Set(k, v)
	}
	return &http.Response{
		StatusCode:    status,
		Status:        http.StatusText(status),
		Body:          io.NopCloser(strings.NewReader(body)),
		Header:        h,
		ContentLength: int64(len(body)),
	}
}

func TestDownloadKeyStreamsBinaryAndParsesFilename(t *testing.T) {
	var captured *http.Request
	c := newTestClient(func(r *http.Request) (*http.Response, error) {
		captured = r
		return binResp(200, "BINARYDATA", map[string]string{
			"Content-Type":        "application/octet-stream",
			"Content-Disposition": `attachment; filename="MyApp.oap"`,
		}), nil
	})

	dl, err := c.DownloadKey(context.Background(), "dl-key-123")
	if err != nil {
		t.Fatalf("DownloadKey error: %v", err)
	}
	defer dl.Close()

	// URL + base path.
	if got, want := captured.URL.Path, BasePath+"/downloads/dl-key-123/"; got != want {
		t.Errorf("path = %q, want %q", got, want)
	}
	// Bearer attached for the profile host.
	if got, want := captured.Header.Get("Authorization"), "Bearer tok-123"; got != want {
		t.Errorf("Authorization = %q, want %q", got, want)
	}
	if dl.Filename != "MyApp.oap" {
		t.Errorf("filename = %q, want MyApp.oap", dl.Filename)
	}
	if dl.ContentType != "application/octet-stream" {
		t.Errorf("content-type = %q", dl.ContentType)
	}
	data, _ := io.ReadAll(dl.Body)
	if string(data) != "BINARYDATA" {
		t.Errorf("body = %q, want BINARYDATA", data)
	}
}

func TestDownloadURLSameHostSendsAuth(t *testing.T) {
	var captured *http.Request
	c := newTestClient(func(r *http.Request) (*http.Response, error) {
		captured = r
		return binResp(200, "X", nil), nil
	})

	dl, err := c.DownloadURL(context.Background(), "https://lifetime.example.com/downloads/abc/")
	if err != nil {
		t.Fatalf("DownloadURL error: %v", err)
	}
	dl.Close()
	if got := captured.Header.Get("Authorization"); got != "Bearer tok-123" {
		t.Errorf("same-host should send auth, got %q", got)
	}
}

func TestDownloadURLDifferentHostOmitsAuth(t *testing.T) {
	var captured *http.Request
	c := newTestClient(func(r *http.Request) (*http.Response, error) {
		captured = r
		return binResp(200, "X", nil), nil
	})

	// Pre-signed link to object storage: must NOT leak the LifeTime token.
	dl, err := c.DownloadURL(context.Background(), "https://cdn.example.net/signed/abc?sig=xyz")
	if err != nil {
		t.Fatalf("DownloadURL error: %v", err)
	}
	dl.Close()
	if got := captured.Header.Get("Authorization"); got != "" {
		t.Errorf("cross-host should NOT send auth, got %q", got)
	}
}

func TestDownloadNon2xxReturnsAPIError(t *testing.T) {
	c := newTestClient(func(r *http.Request) (*http.Response, error) {
		return binResp(404, "no such key", nil), nil
	})
	_, err := c.DownloadKey(context.Background(), "missing")
	if err == nil {
		t.Fatal("expected error")
	}
	if _, ok := err.(*APIError); !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
}

func TestFilenameFromContentDisposition(t *testing.T) {
	cases := map[string]string{
		`attachment; filename="a.oap"`: "a.oap",
		`attachment; filename=b.zip`:   "b.zip",
		`inline`:                       "",
		``:                             "",
		`attachment; filename*=UTF-8''my%20app.oap`: "my app.oap",
		`attachment; filename="../../etc/passwd"`:   "passwd", // base() strips path traversal
	}
	for header, want := range cases {
		if got := filenameFromContentDisposition(header); got != want {
			t.Errorf("header %q: got %q, want %q", header, got, want)
		}
	}
}
