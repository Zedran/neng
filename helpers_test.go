package neng

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Default test data directory
const test_data_directory string = "test_data"

/* Extracts [][]string from a test data file. */
func loadTest2DSliceString(fname string) ([][]string, error) {
	stream, err := os.ReadFile(filepath.Join(test_data_directory, fname))
	if err != nil {
		return nil, err
	}

	var cases [][]string
	if err := json.Unmarshal(stream, &cases); err != nil {
		return nil, err
	}

	return cases, nil
}

/* Unmarshals JSON file to an arbitrary data structure. */
func loadTestData(fname string, v any) error {
	stream, err := os.ReadFile(filepath.Join(test_data_directory, fname))
	if err != nil {
		return err
	}

	return json.Unmarshal(stream, v)
}

/* Extracts map[string]int from a test data file. */
func loadTestMapStringInt(fname string) (map[string]int, error) {
	stream, err := os.ReadFile(filepath.Join(test_data_directory, fname))
	if err != nil {
		return nil, err
	}

	var cases map[string]int
	if err := json.Unmarshal(stream, &cases); err != nil {
		return nil, err
	}

	return cases, nil
}

/* Extracts map[string]string from a test data file. */
func loadTestMapStringString(fname string) (map[string]string, error) {
	stream, err := os.ReadFile(filepath.Join(test_data_directory, fname))
	if err != nil {
		return nil, err
	}

	var cases map[string]string
	if err := json.Unmarshal(stream, &cases); err != nil {
		return nil, err
	}

	return cases, nil
}
