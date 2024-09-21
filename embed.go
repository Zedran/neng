package neng

import (
	"embed"
	"strings"
)

//go:embed embed/*
var efs embed.FS

// readEFS loads a word list from the embedded path.
// Returns an error if the file is not found.
func readEFS(path string) ([]string, error) {
	stream, err := efs.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(stream), "\n"), nil
}
