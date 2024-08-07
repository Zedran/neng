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
		expected *word
	}

	cases := []testCase{
		{true, "word 0", &word{nil, 0, "word"}},                                  // Regular
		{true, "word 1 f1,f2,f3", &word{&[]string{"f1", "f2", "f3"}, 1, "word"}}, // Irregular
		{true, "word 1 f1a_b,f2", &word{&[]string{"f1a b", "f2"}, 1, "word"}},    // Multi-word irregular
		{true, "word 2", &word{nil, 2, "word"}},                                  // Plural-only
		{true, "word 3", &word{nil, 3, "word"}},                                  // Suffixed
		{true, "word 4", &word{nil, 4, "word"}},                                  // Uncomparable
		{true, "word 5", &word{nil, 5, "word"}},                                  // Uncountable
		{false, "", nil},                                                         // Error: line of zero-length
		{false, "word", nil},                                                     // Error: no type field
		{false, "word ", nil},                                                    // Error: type field empty
		{false, "word 55", nil},                                                  // Error: type field too long
		{false, "word 1", nil},                                                   // Error: irregular without forms field
		{false, "word 1 ,", nil},                                                 // Error: zero-length irregular forms
	}

	for _, c := range cases {
		out, err := NewWord([]byte(c.line))

		if c.good {
			switch true {
			case err != nil:
				t.Errorf("Failed for case %v: error returned: '%s'", c, err.Error())
			case out.word != c.expected.word:
				t.Errorf("Failed for case %v: expected word '%s', got '%s'", c, c.expected.word, out.word)
			case out.t != c.expected.t:
				t.Errorf("Failed for case %v: expected type '%d', got '%d'", c, c.expected.t, out.t)
			case out.t == wt_irregular:
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
