package neng

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Unmarshals JSON file to an arbitrary data structure.
func loadTestData(fname string, v any) error {
	stream, err := os.ReadFile(filepath.Join("test_data", fname))
	if err != nil {
		return err
	}

	return json.Unmarshal(stream, v)
}
