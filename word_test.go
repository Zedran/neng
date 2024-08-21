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
		{true, "0word", &Word{0, nil, "word"}},                               // Regular
		{true, "1word,f2", &Word{1, &[]string{"f2"}, "word"}},                // Irregular, one form
		{true, "1word,f", &Word{1, &[]string{"f"}, "word"}},                  // One-letter irregular form
		{true, "1word,f,f", &Word{1, &[]string{"f", "f"}, "word"}},           // One-letter irregular forms
		{true, "1word,f2,f3", &Word{1, &[]string{"f2", "f3"}, "word"}},       // Irregular, two forms
		{true, "1word,f1a b,f2", &Word{1, &[]string{"f1a b", "f2"}, "word"}}, // Multi-word irregular
		{true, "2word", &Word{2, nil, "word"}},                               // Plural-only
		{true, "3word", &Word{3, nil, "word"}},                               // Suffixed
		{true, "4word", &Word{4, nil, "word"}},                               // Uncomparable
		{true, "5word", &Word{5, nil, "word"}},                               // Uncountable
		{false, "6word", nil},                                                // Error: Type value out of defined range for FormType
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

/* Tests whether NewWordFromParams enforces the designed limitations. */
func TestNewWordFromParams(t *testing.T) {
	type testCase struct {
		good     bool
		word     string
		ft       FormType
		irr      []string
		expected *Word
	}

	w := "word"

	cases := []testCase{
		{true, w, FT_REGULAR, nil, &Word{FT_REGULAR, nil, w}},                                                     // Regular
		{true, w, FT_IRREGULAR, []string{"f1"}, &Word{FT_IRREGULAR, &[]string{"f1"}, w}},                          // Irregular, one form
		{true, w, FT_IRREGULAR, []string{"f1", "f2"}, &Word{FT_IRREGULAR, &[]string{"f1", "f2"}, w}},              // Irregular, two forms
		{true, w, FT_PLURAL_ONLY, nil, &Word{FT_PLURAL_ONLY, nil, w}},                                             // Plural-only
		{true, w, FT_SUFFIXED, nil, &Word{FT_SUFFIXED, nil, w}},                                                   // Suffixed
		{true, w, FT_UNCOMPARABLE, nil, &Word{FT_UNCOMPARABLE, nil, w}},                                           // Uncomparable
		{true, w, FT_UNCOUNTABLE, nil, &Word{FT_UNCOUNTABLE, nil, w}},                                             // Uncountable
		{false, w, FT_IRREGULAR, []string{"f1", "f2", "f3"}, &Word{FT_IRREGULAR, &[]string{"f1", "f2", "f3"}, w}}, // Error: too many forms
		{false, w, FT_SUFFIXED, []string{"f1"}, &Word{FT_SUFFIXED, nil, w}},                                       // Error: irregular forms for non-irregular
		{false, w, FT_IRREGULAR, []string{}, nil},                                                                 // Error: empty slice for irregular
		{false, w, FT_IRREGULAR, nil, nil},                                                                        // Error: nil slice for irregular
		{false, "", FT_REGULAR, nil, nil},                                                                         // Error: empty word
		{false, w, 255, nil, nil},                                                                                 // Error: undefined FormType
		{false, w, FT_IRREGULAR, []string{""}, nil},                                                               // Error: first irregular form empty
		{false, w, FT_IRREGULAR, []string{w, ""}, nil},                                                            // Error: second irregular form empty
	}

	for i, c := range cases {
		out, err := NewWordFromParams(c.word, c.ft, c.irr)

		if c.good {
			switch true {
			case err != nil:
				t.Errorf("Failed for case %d: error returned: '%s'", i, err.Error())
			case out.word != c.expected.word:
				t.Errorf("Failed for case %d: expected word '%s', got '%s'", i, c.expected.word, out.word)
			case out.ft != c.expected.ft:
				t.Errorf("Failed for case %d: expected FormType '%d', got '%d'", i, c.expected.ft, out.ft)
			case out.ft == FT_IRREGULAR:
				if out.irr == nil || !slices.Equal(*out.irr, *c.expected.irr) {
					t.Errorf("Failed for case %d: slices are not equal, expected %v, got %v", i, c.expected.irr, out.irr)
				}
			case out.ft != FT_IRREGULAR:
				if out.irr != nil {
					t.Errorf("Failed for case %d: irregular slice assigned to non-irregular", i)
				}
			}
		} else {
			if err == nil {
				t.Errorf("Failed for case %d: no error returned, got %v", i, out)
			}
		}
	}
}
