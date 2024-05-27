package neng

import "testing"

/* Tests whether countSyllables returns a proper number of syllables for sample words. */
func TestCountSyllables(t *testing.T) {
	cases := map[string]int{
		"agree":       2,
		"be":          1,
		"beat":        1,
		"buoy":        1,
		"care":        1,
		"carry":       2,
		"caveat":      3,
		"chikadee":    3,
		"clear":       1,
		"create":      2,
		"commemorate": 4,
		"commission":  3,
		"commit":      2,
		"country":     2,
		"covenant":    3,
		"decoy":       2,
		"eat":         1,
		"entresol":    3,
		"evacuate":    4,
		"exfoliate":   4,
		"ford":        1,
		"go":          1,
		"heave":       1,
		"lee":         1,
		"panic":       2,
		"receipt":     2,
		"reposit":     3,
		"salaam":      2,
		"spree":       1,
		"stop":        1,
		"study":       2,
		"torpedo":     3,
		"vex":         1,
		"":            0,
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
		"do",
		"freeze",
		"forgive",
		"give",
	}

	irregular, err := loadIrregularWords("res/verb.irr")
	if err != nil {
		t.Fatalf("loadIrregularWords failed: %s", err.Error())
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
