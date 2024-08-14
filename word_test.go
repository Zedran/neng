package neng

import (
	"slices"
	"testing"
)

/* Ensures that NewWord function correctly parses the received line and returns error if malformed input is encountered. */
func TestNewWord(t *testing.T) {
	type testCase struct {
		good     bool
		line     string
		expected *Word
	}

	cases := []testCase{
		{true, "0word", &Word{nil, 0, "word"}},                               // Regular
		{true, "1word,f2", &Word{&[]string{"f2"}, 1, "word"}},                // Irregular, one form
		{true, "1word,f", &Word{&[]string{"f"}, 1, "word"}},                  // One-letter irregular form
		{true, "1word,f,f", &Word{&[]string{"f", "f"}, 1, "word"}},           // One-letter irregular forms
		{true, "1word,f2,f3", &Word{&[]string{"f2", "f3"}, 1, "word"}},       // Irregular, two forms
		{true, "1word,f1a b,f2", &Word{&[]string{"f1a b", "f2"}, 1, "word"}}, // Multi-word irregular
		{true, "2word", &Word{nil, 2, "word"}},                               // Plural-only
		{true, "3word", &Word{nil, 3, "word"}},                               // Suffixed
		{true, "4word", &Word{nil, 4, "word"}},                               // Uncomparable
		{true, "5word", &Word{nil, 5, "word"}},                               // Uncountable
		{false, "6word", &Word{nil, 5, "word"}},                              // Error: Type value out of defined range for FormType
		{false, "0word,f", nil},                                              // Error: Non-irregular with one irregular forms
		{false, "0word,f,f", nil},                                            // Error: Non-irregular with two irregular forms
		{false, "0word,", nil},                                               // Error: Non-irregular with comma at the end of the line
		{false, "", nil},                                                     // Error: empty line
		{false, "0", nil},                                                    // Error: FormType field only, regular
		{false, "1", nil},                                                    // Error: FormType field only, irregular
		{false, "word", nil},                                                 // Error: no FormType field at the beginning of the line
		{false, "1,f1,f2", nil},                                              // Error: no word
		{false, ",f1", nil},                                                  // Error: no FormType, no word, just an irregular form
		{false, "1word", nil},                                                // Error: irregular without irregular forms
		{false, "1word,", nil},                                               // Error: one zero-length irregular form
		{false, "1word,,", nil},                                              // Error: two zero-length irregular forms
		{false, "1word,f2,f3,f4", nil},                                       // Error: too many irregular forms
	}

	for _, c := range cases {
		out, err := NewWord(c.line)

		if c.good {
			switch true {
			case err != nil:
				t.Errorf("Failed for case %v: error returned: '%s'", c, err.Error())
			case out.word != c.expected.word:
				t.Errorf("Failed for case %v: expected word '%s', got '%s'", c, c.expected.word, out.word)
			case out.ft != c.expected.ft:
				t.Errorf("Failed for case %v: expected FormType '%d', got '%d'", c, c.expected.ft, out.ft)
			case out.ft == FT_IRREGULAR:
				if out.irr == nil || !slices.Equal(*out.irr, *c.expected.irr) {
					t.Errorf("Failed for case %v: slices are not equal, expected %v, got %v", c, c.expected.irr, out.irr)
				}
			}
		} else {
			if err == nil {
				t.Errorf("Failed for case %v: no error returned, got %v", c, out)
			}
		}
	}
}
