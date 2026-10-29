package client

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/jinjie/outsystems-cli/internal/config"
)

// roundTripFunc lets a test intercept the outgoing request and return a canned
// response, so we can assert on URL, headers, and body without a live server.
type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func newTestClient(rt roundTripFunc) *Client {
	c := New(config.Profile{Hostname: "lifetime.example.com", APIKey: "tok-123"})
	c.httpClient = &http.Client{Transport: rt}
	return c
}

func respond(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Status:     http.StatusText(status),
		Body:       io.NopCloser(strings.NewReader(body)),
		Header:     make(http.Header),
	}
}

func TestGetBuildsURLAndBearerHeader(t *testing.T) {
	var captured *http.Request
	c := newTestClient(func(r *http.Request) (*http.Response, error) {
		captured = r
		return respond(200, `{"ok":true}`), nil
	})

	q := url.Values{}
	q.Set("IncludeModules", "true")
	body, err := c.Get(context.Background(), "/applications/", q)
	if err != nil {
		t.Fatalf("Get returned error: %v", err)
	}
	if string(body) != `{"ok":true}` {
		t.Fatalf("unexpected body: %s", body)
	}

	if got, want := captured.URL.Scheme, "https"; got != want {
		t.Errorf("scheme = %q, want %q", got, want)
	}
	if got, want := captured.URL.Host, "lifetime.example.com"; got != want {
		t.Errorf("host = %q, want %q", got, want)
	}
	if got, want := captured.URL.Path, BasePath+"/applications/"; got != want {
		t.Errorf("path = %q, want %q", got, want)
	}
	if got, want := captured.URL.Query().Get("IncludeModules"), "true"; got != want {
		t.Errorf("query IncludeModules = %q, want %q", got, want)
	}
	if got, want := captured.Header.Get("Authorization"), "Bearer tok-123"; got != want {
		t.Errorf("Authorization = %q, want %q", got, want)
	}
}

func TestPostSendsBodyAndContentType(t *testing.T) {
	var captured *http.Request
	var capturedBody []byte
	c := newTestClient(func(r *http.Request) (*http.Response, error) {
		captured = r
		if r.Body != nil {
			capturedBody, _ = io.ReadAll(r.Body)
		}
		return respond(200, "new-key"), nil
	})

	payload := []byte(`{"Name":"X"}`)
	if _, err := c.Post(context.Background(), "/deployments/", nil, payload); err != nil {
		t.Fatalf("Post returned error: %v", err)
	}
	if got, want := captured.Method, http.MethodPost; got != want {
		t.Errorf("method = %q, want %q", got, want)
	}
	if got, want := captured.Header.Get("Content-Type"), "application/json"; got != want {
		t.Errorf("Content-Type = %q, want %q", got, want)
	}
	if string(capturedBody) != string(payload) {
		t.Errorf("body = %q, want %q", capturedBody, payload)
	}
}

func TestNon2xxReturnsAPIError(t *testing.T) {
	c := newTestClient(func(r *http.Request) (*http.Response, error) {
		return respond(404, `{"Errors":["not found"]}`), nil
	})
	_, err := c.Get(context.Background(), "/applications/missing/", nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.StatusCode != 404 {
		t.Errorf("status code = %d, want 404", apiErr.StatusCode)
	}
	if !strings.Contains(apiErr.Error(), "not found") {
		t.Errorf("error should contain server body, got: %v", apiErr.Error())
	}
}

func TestEmptyBodyOmitsContentType(t *testing.T) {
	var captured *http.Request
	c := newTestClient(func(r *http.Request) (*http.Response, error) {
		captured = r
		return respond(200, ""), nil
	})
	if _, err := c.Post(context.Background(), "/deployments/key/start/", nil, nil); err != nil {
		t.Fatalf("Post returned error: %v", err)
	}
	if got := captured.Header.Get("Content-Type"); got != "" {
		t.Errorf("Content-Type should be empty for payload-less request, got %q", got)
	}
}
