package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// bytesBody wraps a byte slice as an io.Reader for client.Do, returning nil for
// an empty body so no Content-Type header is set. Used by requests that need a
// verb/body combination not covered by the typed client helpers (e.g. DELETE
// with a JSON body).
func bytesBody(b []byte) io.Reader {
	if len(b) == 0 {
		return nil
	}
	return bytes.NewReader(b)
}

// runAndPrint resolves the client, runs the given API call, and writes the
// resulting JSON body using the configured output writer. It is the common
// path for read commands and for write commands that return a payload.
//
// The call closure receives the resolved client via resolveClient so each
// command stays a one-liner.
func runAndPrint(do func() ([]byte, error)) error {
	writer, err := newWriter()
	if err != nil {
		return err
	}
	raw, err := do()
	if err != nil {
		return err
	}
	return writer.WriteJSON(raw)
}

// readBodyInput resolves a --data flag value into raw request bytes.
//
// Supported forms (pipe-friendly):
//   - ""          : no body (returns nil, nil)
//   - "@path"     : read the file at path
//   - "-"         : read from stdin
//   - "<json>"    : use the literal string as the body
func readBodyInput(data string) ([]byte, error) {
	switch {
	case data == "":
		return nil, nil
	case data == "-":
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			return nil, fmt.Errorf("read body from stdin: %w", err)
		}
		return b, nil
	case strings.HasPrefix(data, "@"):
		path := strings.TrimPrefix(data, "@")
		b, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("read body file %s: %w", path, err)
		}
		return b, nil
	default:
		return []byte(data), nil
	}
}

// boolQuery adds a boolean query parameter only when the flag was changed from
// its default, so the API falls back to its own defaults otherwise. Callers
// pass the value and whether it was explicitly set.
func addBoolQuery(q map[string][]string, key string, value, set bool) {
	if !set {
		return
	}
	if value {
		q[key] = []string{"true"}
	} else {
		q[key] = []string{"false"}
	}
}
