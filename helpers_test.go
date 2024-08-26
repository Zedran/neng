package neng

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Default test data directory
const test_data_directory string = "test_data"

/* Unmarshals JSON file to an arbitrary data structure. */
func loadTestData(fname string, v any) error {
	stream, err := os.ReadFile(filepath.Join(test_data_directory, fname))
	if err != nil {
		return err
	}

	return json.Unmarshal(stream, v)
}
