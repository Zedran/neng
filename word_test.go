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
		{true, "1word,f2,f3", &Word{&[]string{"f2", "f3"}, 1, "word"}},       // Irregular, two forms
		{true, "1word,f1a b,f2", &Word{&[]string{"f1a b", "f2"}, 1, "word"}}, // Multi-word irregular
		{true, "2word", &Word{nil, 2, "word"}},                               // Plural-only
		{true, "3word", &Word{nil, 3, "word"}},                               // Suffixed
		{true, "4word", &Word{nil, 4, "word"}},                               // Uncomparable
		{true, "5word", &Word{nil, 5, "word"}},                               // Uncountable
		{false, "6word", &Word{nil, 5, "word"}},                              // Type value out of defined range for wordType
		{false, "", nil},                                                     // Error: empty line
		{false, "0", nil},                                                    // Error: type field only, regular
		{false, "1", nil},                                                    // Error: type field only, irregular
		{false, "word", nil},                                                 // Error: no type field at the beginning of the line
		{false, "1,f1,f2", nil},                                              // Error: no word
		{false, ",f1", nil},                                                  // Error: no type, no word, just an irregular form
		{false, "1word", nil},                                                // Error: irregular without forms field
		{false, "1word,", nil},                                               // Error: one zero-length irregular form
		{false, "1word,,", nil},                                              // Error: two zero-length irregular forms
	}

	for _, c := range cases {
		out, err := NewWord(c.line)

		if c.good {
			switch true {
			case err != nil:
				t.Errorf("Failed for case %v: error returned: '%s'", c, err.Error())
			case out.word != c.expected.word:
				t.Errorf("Failed for case %v: expected word '%s', got '%s'", c, c.expected.word, out.word)
			case out.t != c.expected.t:
				t.Errorf("Failed for case %v: expected type '%d', got '%d'", c, c.expected.t, out.t)
			case out.t == WT_IRREGULAR:
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
