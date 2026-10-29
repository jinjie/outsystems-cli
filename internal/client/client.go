// Package client is a thin HTTP client for the OutSystems LifeTime REST API v2.
//
// It builds request URLs from a profile hostname, attaches the Bearer token,
// and returns raw JSON response bodies. Decoding/printing is left to callers
// so that responses can be streamed straight to stdout for piping.
package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/jinjie/outsystems-cli/internal/config"
)

// BasePath is the fixed REST base path for the LifeTime v2 API. It is the same
// for every environment, so profiles only store the hostname.
const BasePath = "/lifetimeapi/rest/v2"

// Client talks to a single LifeTime environment.
type Client struct {
	hostname   string
	apiKey     string
	httpClient *http.Client
}

// New builds a Client from a resolved profile.
func New(p config.Profile) *Client {
	return &Client{
		hostname:   p.Hostname,
		apiKey:     p.APIKey,
		httpClient: &http.Client{Timeout: 120 * time.Second},
	}
}

// APIError represents a non-2xx response from the API. The raw body is kept so
// callers can surface the server's error payload.
type APIError struct {
	StatusCode int
	Status     string
	Body       []byte
}

func (e *APIError) Error() string {
	body := strings.TrimSpace(string(e.Body))
	if body == "" {
		return fmt.Sprintf("API request failed: %s", e.Status)
	}
	return fmt.Sprintf("API request failed: %s: %s", e.Status, body)
}

// buildURL constructs the absolute URL for an API path (relative to BasePath)
// with optional query parameters. The path should start with "/".
func (c *Client) buildURL(apiPath string, query url.Values) string {
	u := url.URL{
		Scheme: "https",
		Host:   c.hostname,
		Path:   BasePath + apiPath,
	}
	if len(query) > 0 {
		u.RawQuery = query.Encode()
	}
	return u.String()
}

// Do performs an HTTP request against the API and returns the response body on
// success (2xx). Non-2xx responses return an *APIError. A nil body is allowed
// for requests without a payload.
func (c *Client) Do(ctx context.Context, method, apiPath string, query url.Values, body io.Reader) ([]byte, error) {
	reqURL := c.buildURL(apiPath, query)

	req, err := http.NewRequestWithContext(ctx, method, reqURL, body)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request %s %s: %w", method, apiPath, err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			Body:       data,
		}
	}
	return data, nil
}

// Get performs a GET request.
func (c *Client) Get(ctx context.Context, apiPath string, query url.Values) ([]byte, error) {
	return c.Do(ctx, http.MethodGet, apiPath, query, nil)
}

// Post performs a POST request with an optional JSON body.
func (c *Client) Post(ctx context.Context, apiPath string, query url.Values, body []byte) ([]byte, error) {
	return c.Do(ctx, http.MethodPost, apiPath, query, bytesReader(body))
}

// Put performs a PUT request with an optional JSON body.
func (c *Client) Put(ctx context.Context, apiPath string, query url.Values, body []byte) ([]byte, error) {
	return c.Do(ctx, http.MethodPut, apiPath, query, bytesReader(body))
}

// Delete performs a DELETE request.
func (c *Client) Delete(ctx context.Context, apiPath string, query url.Values) ([]byte, error) {
	return c.Do(ctx, http.MethodDelete, apiPath, query, nil)
}

// Download is a streaming binary response. The caller owns Body and MUST close
// it. Filename is the server-suggested name parsed from the Content-Disposition
// header, or empty when the server did not provide one.
type Download struct {
	Body        io.ReadCloser
	Filename    string
	ContentType string
	// Size is the Content-Length when known, or -1 otherwise.
	Size int64
}

// Close releases the underlying response body.
func (d *Download) Close() error {
	if d.Body == nil {
		return nil
	}
	return d.Body.Close()
}

// DownloadKey streams the binary for a download key from the LifeTime
// /downloads/{DownloadKey}/ endpoint using the profile host and Bearer token.
// The returned Download must be closed by the caller.
func (c *Client) DownloadKey(ctx context.Context, downloadKey string) (*Download, error) {
	reqURL := c.buildURL("/downloads/"+url.PathEscape(downloadKey)+"/", nil)
	return c.downloadURL(ctx, reqURL, true)
}

// DownloadURL streams a binary from an absolute URL, such as the "url" field of
// a DownloadLink returned by the *-download commands. The Bearer token is sent
// only when the URL targets the profile host, so pre-signed links to other
// hosts (e.g. object storage) are not leaked the LifeTime credential.
// The returned Download must be closed by the caller.
func (c *Client) DownloadURL(ctx context.Context, rawURL string) (*Download, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("parse download url: %w", err)
	}
	sameHost := strings.EqualFold(parsed.Host, c.hostname)
	return c.downloadURL(ctx, rawURL, sameHost)
}

// downloadURL performs the streaming GET. When auth is true the Bearer token is
// attached. The response body is left open on success for the caller to stream.
func (c *Client) downloadURL(ctx context.Context, rawURL string, auth bool) (*Download, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return nil, fmt.Errorf("build download request: %w", err)
	}
	if auth {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("download request: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		// Read the (usually small) error body, then close.
		data, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			Body:       data,
		}
	}

	return &Download{
		Body:        resp.Body,
		Filename:    filenameFromContentDisposition(resp.Header.Get("Content-Disposition")),
		ContentType: resp.Header.Get("Content-Type"),
		Size:        resp.ContentLength,
	}, nil
}

// filenameFromContentDisposition extracts a filename from a Content-Disposition
// header value, returning "" when none is present or it cannot be parsed.
func filenameFromContentDisposition(header string) string {
	if header == "" {
		return ""
	}
	_, params, err := mime.ParseMediaType(header)
	if err != nil {
		return ""
	}
	// RFC 5987 extended form takes precedence when present.
	if v := params["filename*"]; v != "" {
		if name := decodeRFC5987(v); name != "" {
			return filepath.Base(name)
		}
	}
	if v := params["filename"]; v != "" {
		return filepath.Base(v)
	}
	return ""
}

// decodeRFC5987 decodes a value of the form "UTF-8”percent%20encoded". On any
// problem it returns the input unchanged.
func decodeRFC5987(v string) string {
	parts := strings.SplitN(v, "'", 3)
	if len(parts) != 3 {
		return v
	}
	decoded, err := url.QueryUnescape(parts[2])
	if err != nil {
		return parts[2]
	}
	return decoded
}

// bytesReader returns an io.Reader for the body, or nil when the body is empty
// so that no Content-Type header is set for payload-less requests.
func bytesReader(body []byte) io.Reader {
	if len(body) == 0 {
		return nil
	}
	return bytes.NewReader(body)
}
