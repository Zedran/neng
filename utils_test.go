package neng

import "testing"

/* Tests whether countSyllables returns a proper number of syllables for sample words. */
func TestCountSyllables(t *testing.T) {
	cases, err := loadTestMapStringInt("TestCountSyllables.json")
	if err != nil {
		t.Fatalf("Failed loading test data: %s", err.Error())
	}

	for input, expected := range cases {
		seq := getSequence(input)
		output := countSyllables(input, seq)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%d', got '%d'", input, expected, output)
		}
	}
}

/* Tests getSequence function. */
func TestGetSequence(t *testing.T) {
	cases, err := loadTestMapStringString("TestGetSequence.json")
	if err != nil {
		t.Fatalf("Failed loading test data: %s", err.Error())
	}

	for input, expected := range cases {
		output := getSequence(input)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, expected, output)
		}
	}
}

/* Tests whether findIrregular returns a correct verb line. */
func TestFindIrregular(t *testing.T) {
	cases, err := loadTest2DSliceString("TestFindIrregular.json")
	if err != nil {
		t.Fatalf("Failed loading test data: %s", err.Error())
	}

	irregular, err := loadIrregularWords("res/verb.irr")
	if err != nil {
		t.Fatalf("loadIrregularWords failed: %s", err.Error())
	}

	for _, input := range cases {
		output := findIrregular(input[0], irregular)

		if output == nil {
			t.Errorf("Failed for '%s': nil returned", input)
		} else if output[0] != input[0] {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, input, output)
		}
	}

	regularVerb := "panic"

	if output := findIrregular(regularVerb, irregular); output != nil {
		t.Errorf("Failed for '%s': nil not returned", regularVerb)
	}
}
