// Package tests contains common functions used by test
package tests

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const TEST_DATA_DIR string = "testdata"

// ReadData unmarshals JSON file containing test data into an arbitrary
// data structure.
func ReadData(fname string, v any) error {
	stream, err := os.ReadFile(filepath.Join(TEST_DATA_DIR, fname))
	if err != nil {
		return err
	}

	return json.Unmarshal(stream, v)
}
