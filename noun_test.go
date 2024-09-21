package neng

import (
	"testing"

	"github.com/Zedran/neng/internal/tests"
)

// Tests plural. Fails if incorrect plural form of a noun is returned.
func TestPlural(t *testing.T) {
	var cases map[string]string
	if err := tests.ReadData("TestPlural.json", &cases); err != nil {
		t.Fatalf("Failed loading test data: %v", err)
	}

	gen, err := DefaultGenerator()
	if err != nil {
		t.Fatalf("Failed: DefaultGenerator returned an error: %v", err)
	}

	for input, expected := range cases {
		word, err := gen.Find(input, WC_NOUN)
		if err != nil {
			t.Logf("Test case '%s' does not exist in the word database. Assume it is regular and proceed.", input)

			word, err = NewWordFromParams(input, 0, nil)
			if err != nil {
				t.Errorf("Failed for '%s' - error from NewWordFromParams: %v", input, err)
			}
		}

		output := plural(word)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, expected, output)
		}
	}
}
