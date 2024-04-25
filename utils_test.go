package neng

import "testing"

/* Tests getSequence function. */
func TestGetSequence(t *testing.T) {
	cases := map[string]string{
		"agree":      "vccvv",
		"care":       "cvcv",
		"carry":      "cvccv",
		"commission": "cvccvccvvc",
		"covenant":   "cvcvcvcc",
		"decoy":      "cvcvc",
		"ford":       "cvcc",
		"panic":      "cvcvc",
		"stop":       "ccvc",
		"study":      "ccvcv",
		"torpedo":    "cvccvcv",
		"vex":        "cvc",
		"":           "",
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
	cases := []string{
		"be",
		"do",
		"freeze",
		"forgive",
	}

	irregular := [][]string{
		{"be", "was", "been"},
		{"do", "did", "done"},
		{"forgive", "forgave", "forgiven"},
		{"freeze", "froze", "frozen"},
		{"give", "gave", "given"},
	}

	for _, input := range cases {
		output := findIrregular(input, irregular)

		if output == nil {
			t.Errorf("Failed for '%s': nil returned", input)
		} else if output[0] != input {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, input, output)
		}
	}

	regularVerb := "panic"

	if output := findIrregular(regularVerb, irregular); output != nil {
		t.Errorf("Failed for '%s': nil not returned", regularVerb)
	}
}
