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
	"testing"

	"github.com/Zedran/neng/internal/tests"
)

// Tests plural. Fails if incorrect plural form of a noun is returned.
func TestPlural(t *testing.T) {
	var cases map[string]string
	if err := tests.ReadData("TestPlural.json", &cases); err != nil {
		t.Fatalf("Failed loading test data: %v", err)
	}

	gen, err := DefaultGenerator(nil)
	if err != nil {
		t.Fatalf("Failed: DefaultGenerator returned an error: %v", err)
	}

	for input, expected := range cases {
		word, err := gen.Find(input, WC_NOUN)
		if err != nil {
			t.Logf("Test case '%s' does not exist in the word database. Assume it is regular and proceed.", input)

			word, err = NewWordFromParams(input, 0, nil)
			if err != nil {
				t.Errorf("Failed for '%s' - error from NewWordFromParams: %v", input, err)
			}
		}

		output := plural(word)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, expected, output)
		}
	}
}

func TestPossessive(t *testing.T) {
	type testCase struct {
		Input    string `json:"input"`
		Plural   bool   `json:"plural"`
		Expected string `json:"expected"`
	}

	var cases []testCase

	if err := tests.ReadData("TestPossessive.json", &cases); err != nil {
		t.Fatalf("Failed loading test data: %v", err)
	}

	gen, err := DefaultGenerator(nil)
	if err != nil {
		t.Fatalf("Failed: DefaultGenerator returned an error: %v", err)
	}

	for _, c := range cases {
		word, err := gen.Find(c.Input, WC_NOUN)
		if err != nil {
			word, err = NewWordFromParams(c.Input, 0, nil)
			if err != nil {
				t.Errorf("Failed for '%s' - error from NewWordFromParams: %v", c.Input, err)
			}
		}

		output := possessive(word, c.Plural)

		if output != c.Expected {
			t.Errorf("Failed for '%s' (plural = %v): expected '%s', got '%s'", c.Input, c.Plural, c.Expected, output)
		}
	}
}
