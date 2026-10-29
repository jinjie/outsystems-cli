// Package output centralizes how command results are rendered and written.
//
// The CLI emits JSON by default. Successful JSON is pretty-printed for human
// reading but remains valid for piping into tools like jq. On success nothing
// other than the payload is written to stdout, keeping the CLI pipe-friendly.
// When --output is set, the payload is written to the file and stdout stays
// quiet.
package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Format enumerates the supported output formats. Only JSON is implemented for
// now; the type exists so new formats can be added without changing call sites.
type Format string

const (
	// FormatJSON renders results as indented JSON.
	FormatJSON Format = "json"
)

// ParseFormat validates a format string.
func ParseFormat(s string) (Format, error) {
	switch Format(s) {
	case FormatJSON:
		return FormatJSON, nil
	default:
		return "", fmt.Errorf("unsupported format %q (supported: json)", s)
	}
}

// Writer renders payloads according to a chosen format and destination.
type Writer struct {
	Format Format
	// File, when non-empty, is the path the payload is written to instead of
	// stdout. On success, stdout stays quiet.
	File string
	// Stdout is the stream used when File is empty. Defaults to os.Stdout.
	Stdout io.Writer
}

// NewWriter builds a Writer with sensible defaults.
func NewWriter(format Format, file string) *Writer {
	return &Writer{Format: format, File: file, Stdout: os.Stdout}
}

// WriteJSON renders raw JSON bytes from the API. It re-indents valid JSON for
// readability; if the body is not valid JSON (some endpoints return plain
// text, e.g. a newly created key), it is written through unchanged.
func (w *Writer) WriteJSON(raw []byte) error {
	formatted := reindentJSON(raw)
	return w.write(formatted)
}

// WriteValue marshals an arbitrary Go value to JSON and writes it. Useful for
// results the CLI constructs itself rather than proxying from the API.
func (w *Writer) WriteValue(v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal output: %w", err)
	}
	return w.write(append(data, '\n'))
}

// reindentJSON pretty-prints valid JSON and appends a trailing newline. If the
// input is not valid JSON, it is returned as-is (with a trailing newline added
// when missing) so plain-text responses still render cleanly.
func reindentJSON(raw []byte) []byte {
	trimmed := bytes.TrimSpace(raw)
	if len(trimmed) == 0 {
		return []byte("")
	}
	var buf bytes.Buffer
	if err := json.Indent(&buf, trimmed, "", "  "); err != nil {
		// Not JSON: return the original trimmed content with a newline.
		return append(trimmed, '\n')
	}
	buf.WriteByte('\n')
	return buf.Bytes()
}

// write sends the payload to the file (if set) or stdout.
func (w *Writer) write(data []byte) error {
	if w.File != "" {
		if err := os.WriteFile(w.File, data, 0o644); err != nil {
			return fmt.Errorf("write output file %s: %w", w.File, err)
		}
		return nil
	}
	out := w.Stdout
	if out == nil {
		out = os.Stdout
	}
	if _, err := out.Write(data); err != nil {
		return fmt.Errorf("write output: %w", err)
	}
	return nil
}
