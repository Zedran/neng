package neng

import "testing"

/* Tests plural function. Fails if incorrect plural form of a noun is returned. */
func TestPlural(t *testing.T) {
	cases, err := loadTestMapStringString("TestPlural.json")
	if err != nil {
		t.Fatalf("Failed loading test data: %s", err.Error())
	}

	pluralOnly, err := loadWords("res/noun.plo")
	if err != nil {
		t.Fatalf("loadIrregularWords failed: %s", err.Error())
	}

	irregular, err := loadIrregularWords("res/noun.irr")
	if err != nil {
		t.Fatalf("loadIrregularWords failed: %s", err.Error())
	}

	for input, expected := range cases {
		output := plural(input, pluralOnly, irregular)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, expected, output)
		}
	}
}
