// neng -- Non-Extravagant Name Generator
// Copyright (C) 2024  Wojciech Głąb (github.com/Zedran)
//
// This file is part of neng.
//
// neng is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3 only.
//
// neng is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with neng.  If not, see <https://www.gnu.org/licenses/>.

package neng

import (
	"slices"
	"testing"
)

// Ensures that NewWord correctly parses the received line and returns an error
// if malformed input is encountered.
func TestNewWord(t *testing.T) {
	type testCase struct {
		good     bool
		line     string
		expected Word
	}

	cases := []testCase{
		{true, "0word", Word{FT_REGULAR, nil, "word"}},                                 // Regular
		{true, "1word,f2", Word{FT_IRREGULAR, &[]string{"f2"}, "word"}},                // Irregular, one form
		{true, "1word,f", Word{FT_IRREGULAR, &[]string{"f"}, "word"}},                  // One-letter irregular form
		{true, "1word,f,f", Word{FT_IRREGULAR, &[]string{"f", "f"}, "word"}},           // One-letter irregular forms
		{true, "1word,f2,f3", Word{FT_IRREGULAR, &[]string{"f2", "f3"}, "word"}},       // Irregular, two forms
		{true, "1word,f1a b,f2", Word{FT_IRREGULAR, &[]string{"f1a b", "f2"}, "word"}}, // Multi-word irregular
		{true, "2word", Word{FT_PLURAL_ONLY, nil, "word"}},                             // Plural-only
		{true, "3word", Word{FT_SUFFIXED, nil, "word"}},                                // Suffixed
		{true, "4word", Word{FT_NON_COMPARABLE, nil, "word"}},                          // Non-comparable
		{true, "5word", Word{FT_UNCOUNTABLE, nil, "word"}},                             // Uncountable
		{false, "6word", Word{}},                                                       // Error: Type value out of defined range for FormType
		{false, "0word,f", Word{}},                                                     // Error: Non-irregular with one irregular forms
		{false, "0word,f,f", Word{}},                                                   // Error: Non-irregular with two irregular forms
		{false, "0word,", Word{}},                                                      // Error: Non-irregular with comma at the end of the line
		{false, "", Word{}},                                                            // Error: empty line
		{false, "0", Word{}},                                                           // Error: FormType field only, regular
		{false, "1", Word{}},                                                           // Error: FormType field only, irregular
		{false, "word", Word{}},                                                        // Error: no FormType field at the beginning of the line
		{false, "1,f1,f2", Word{}},                                                     // Error: no word
		{false, ",f1", Word{}},                                                         // Error: no FormType, no word, just an irregular form
		{false, "1word", Word{}},                                                       // Error: irregular without irregular forms
		{false, "1word,", Word{}},                                                      // Error: one zero-length irregular form
		{false, "1word,,", Word{}},                                                     // Error: two zero-length irregular forms
		{false, "1word,f2,f3,f4", Word{}},                                              // Error: too many irregular forms
	}

	for _, c := range cases {
		out, err := NewWord(c.line)

		if c.good {
			switch true {
			case err != nil:
				t.Errorf("Failed for case %v: error returned: %v", c, err)
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

// Tests whether NewWordFromParams enforces the designed limitations.
func TestNewWordFromParams(t *testing.T) {
	type testCase struct {
		good     bool
		word     string
		ft       FormType
		irr      []string
		expected Word
	}

	w := "word"

	cases := []testCase{
		{true, w, FT_REGULAR, nil, Word{FT_REGULAR, nil, w}},                                                     // Regular
		{true, w, FT_IRREGULAR, []string{"f1"}, Word{FT_IRREGULAR, &[]string{"f1"}, w}},                          // Irregular, one form
		{true, w, FT_IRREGULAR, []string{"f1", "f2"}, Word{FT_IRREGULAR, &[]string{"f1", "f2"}, w}},              // Irregular, two forms
		{true, w, FT_PLURAL_ONLY, nil, Word{FT_PLURAL_ONLY, nil, w}},                                             // Plural-only
		{true, w, FT_SUFFIXED, nil, Word{FT_SUFFIXED, nil, w}},                                                   // Suffixed
		{true, w, FT_NON_COMPARABLE, nil, Word{FT_NON_COMPARABLE, nil, w}},                                       // Non-comparable
		{true, w, FT_UNCOUNTABLE, nil, Word{FT_UNCOUNTABLE, nil, w}},                                             // Uncountable
		{false, w, FT_IRREGULAR, []string{"f1", "f2", "f3"}, Word{}},                                             // Error: too many forms
		{false, w, FT_SUFFIXED, []string{"f1"}, Word{}},                                                          // Error: irregular forms for non-irregular
		{false, w, FT_IRREGULAR, []string{}, Word{}},                                                             // Error: empty slice for irregular
		{false, w, FT_IRREGULAR, nil, Word{}},                                                                    // Error: nil slice for irregular
		{false, "", FT_REGULAR, nil, Word{}},                                                                     // Error: empty word
		{false, w, 255, nil, Word{}},                                                                             // Error: undefined FormType
		{false, w, FT_IRREGULAR, []string{""}, Word{}},                                                           // Error: first irregular form empty
		{false, w, FT_IRREGULAR, []string{w, ""}, Word{}},                                                        // Error: second irregular form empty
	}

	for i, c := range cases {
		out, err := NewWordFromParams(c.word, c.ft, c.irr)

		if c.good {
			switch true {
			case err != nil:
				t.Errorf("Failed for case %d: error returned: %v", i, err)
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

// Ensures that the Word.Irr method provides safety when working with
// irregular forms of the Word.
func TestWord_Irr(t *testing.T) {
	type testCase struct {
		good     bool
		word     Word
		index    int
		expected string
	}

	regular, _ := NewWord("0word")
	irregular, _ := NewWord("1good,better,best")

	cases := []testCase{
		{true, irregular, 0, "better"},
		{true, irregular, 1, "best"},
		{false, irregular, -1, ""},
		{false, irregular, 2, ""},
		{false, regular, 0, ""},
	}

	for i, c := range cases {
		output, err := c.word.Irr(c.index)

		if c.good {
			if err != nil {
				t.Errorf("Failed for case %d - error returned: %v", i, err)
			} else if output != c.expected {
				t.Errorf("Failed for case %d - expected %s, got %s", i, c.expected, output)
			}
		} else if err == nil {
			t.Errorf("Failed for case %d - error not returned. Got: %s", i, output)
		}
	}
}
