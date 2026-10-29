package output

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// BinaryResult reports where a streamed binary was written.
type BinaryResult struct {
	// Path is the file written, or "-" when the payload went to stdout.
	Path string
	// Bytes is the number of bytes written.
	Bytes int64
}

// StreamBinary copies a binary stream to its destination, chosen in this order:
//
//  1. outputFile when non-empty and not "-" → that exact path.
//  2. outputFile == "-" → stdout (for piping, e.g. "... -o - > file.oap").
//  3. empty outputFile and a non-empty suggestedName → ./<suggestedName>.
//  4. empty outputFile and no suggestion → stdout.
//
// Writing to a file keeps stdout quiet; the summary is left to the caller to
// print on stderr. Files are not overwritten unless overwrite is true.
func StreamBinary(src io.Reader, outputFile, suggestedName string, overwrite bool) (BinaryResult, error) {
	// Destination: stdout.
	if outputFile == "-" || (outputFile == "" && suggestedName == "") {
		n, err := io.Copy(os.Stdout, src)
		if err != nil {
			return BinaryResult{}, fmt.Errorf("write to stdout: %w", err)
		}
		return BinaryResult{Path: "-", Bytes: n}, nil
	}

	// Destination: a file path.
	path := outputFile
	if path == "" {
		path = filepath.Base(suggestedName)
	}

	flags := os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	if !overwrite {
		// O_EXCL makes the open fail if the file already exists.
		flags |= os.O_EXCL
	}
	f, err := os.OpenFile(path, flags, 0o644)
	if err != nil {
		if os.IsExist(err) {
			return BinaryResult{}, fmt.Errorf("file %s already exists (use --force to overwrite)", path)
		}
		return BinaryResult{}, fmt.Errorf("create output file %s: %w", path, err)
	}
	defer f.Close()

	n, err := io.Copy(f, src)
	if err != nil {
		return BinaryResult{}, fmt.Errorf("write output file %s: %w", path, err)
	}
	return BinaryResult{Path: path, Bytes: n}, nil
}
