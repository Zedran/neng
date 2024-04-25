package neng

import "testing"

/* Tests whether Generator.Phrase correctly parses pattern syntax and generates phrases. */
func TestPhrase(t *testing.T) {
	gen := Generator{
		adjectives: []string{"revocable"},
		nouns:      []string{"snowfall"},
		verbs:      []string{"stash"},
		caser:      newCaser(),
	}

	cases := map[string]string{
		"a pretty %a %n": "a pretty revocable snowfall",
		"%a %n of %n":    "revocable snowfall of snowfall",
		"a n":            "a n",
		"%%a":            "%a",
		"a %2v %n":       "a stashed snowfall",
		"%gv":            "stashing",
		"%Nv":            "stashes",
		"%3v":            "stashed",
		"%nn":            "snowfalln",
		"%%":             "%",
		"%tNv":           "Stashes",
		"%ta %un of %ln": "Revocable SNOWFALL of snowfall",
		"%ttua":          "REVOCABLE",
	}

	for input, expected := range cases {
		output, err := gen.Phrase(input)
		if err != nil {
			t.Errorf("Failed for case '%s': error returned: '%s'", input, err.Error())
		} else if output != expected {
			t.Errorf("Failed for case '%s': got '%s', expected '%s'", input, output, expected)
		}
	}

	errCases := []string{
		"",     // pattern is empty
		"abc%", // escape character at pattern termination
		"%q",   // unknown command
	}

	for _, bc := range errCases {
		output, err := gen.Phrase(bc)
		if err == nil {
			t.Errorf("Failed for errCase '%s': no error returned. Output: '%s'", bc, output)
		}
	}
}
