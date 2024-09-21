package neng

import (
	"testing"

	"github.com/Zedran/neng/internal/tests"
)

// Tests whether countSyllables returns a correct number of syllables.
func TestCountSyllables(t *testing.T) {
	var cases map[string]int
	if err := tests.ReadData("TestCountSyllables.json", &cases); err != nil {
		t.Fatalf("Failed loading test data: %v", err)
	}

	for input, expected := range cases {
		seq := getSequence(input)
		output := countSyllables(input, seq)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%d', got '%d'", input, expected, output)
		}
	}
}

// Tests getSequence.
func TestGetSequence(t *testing.T) {
	var cases map[string]string
	if err := tests.ReadData("TestGetSequence.json", &cases); err != nil {
		t.Fatalf("Failed loading test data: %v", err)
	}

	for input, expected := range cases {
		output := getSequence(input)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, expected, output)
		}
	}
}
