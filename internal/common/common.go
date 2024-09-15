package common

import (
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
func WriteFile(path string, lines []string, sort bool) error {
	if sort {
		slices.Sort(lines)
	}
	return os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
}
