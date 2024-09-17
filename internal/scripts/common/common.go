// This internal package contains common functions used by scripts.
package common

import (
	"crypto/sha256"
	"fmt"
	"os"
	"slices"
	"strings"
)

// Reads a file at OS path and splits its content into lines.
func ReadFile(path string) ([]string, error) {
	stream, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(stream), "\n"), nil
}

// Writes compiled word list to a resource file. Optionally, sorts lines before writing.
// Returns SHA256 checksum of the data written and errors related to file handling.
func WriteFile(path string, lines []string, sort bool) (string, error) {
	if sort {
		slices.Sort(lines)
	}

	data := []byte(strings.Join(lines, "\n"))
	csum := fmt.Sprintf("%x  %s", sha256.Sum256(data), path)

	return csum, os.WriteFile(path, data, 0644)
}
