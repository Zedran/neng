package neng

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Default test data directory
const test_data_directory string = "test_data"

/* A struct used for storing test data related to verb transformations that are influenced by MOD_PLURAL. */
type testCasePlural struct {
	Input    string `json:"input"`
	Expected string `json:"expected"`
	Plural   bool   `json:"plural"`
}

/* Extracts []testCasePlural from a test data file. */
func loadSliceTestCasePlural(fname string) ([]testCasePlural, error) {
	stream, err := os.ReadFile(filepath.Join(test_data_directory, fname))
	if err != nil {
		return nil, err
	}

	var cases []testCasePlural
	if err := json.Unmarshal(stream, &cases); err != nil {
		return nil, err
	}

	return cases, nil
}

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
