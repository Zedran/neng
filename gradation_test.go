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

// Tests comparative. Fails if incorrect comparative form is returned.
// If the word does not exist in the database, the test attempts
// to transform it as FT_REGULAR.
func TestComparative(t *testing.T) {
	type testCase struct {
		Input    string    `json:"input"`
		WC       WordClass `json:"word_class"`
		Expected string    `json:"expected"`
	}

	var cases []testCase
	if err := tests.ReadData("TestComparative.json", &cases); err != nil {
		t.Fatalf("Failed loading test data: %v", err)
	}

	gen, err := DefaultGenerator(nil)
	if err != nil {
		t.Fatalf("Failed: DefaultGenerator returned an error: %v", err)
	}

	for _, c := range cases {
		word, err := gen.Find(c.Input, c.WC)
		if err != nil {
			t.Logf("Test case '%s' (WordClass %d) does not exist in the word database. Assume it is regular and proceed.", c.Input, c.WC)

			word, err = NewWordFromParams(c.Input, 0, nil)
			if err != nil {
				t.Errorf("Failed for '%s' - error from NewWordFromParams: %v", c.Input, err)
			}
		}

		output := comparative(word)

		if output != c.Expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", c.Input, c.Expected, output)
		}
	}
}

// Tests sufGrad. Fails if a malformed graded form is returned.
func TestSufGrad(t *testing.T) {
	var cases [][]string
	if err := tests.ReadData("TestSufGrad.json", &cases); err != nil {
		t.Fatalf("Failed loading test data: %v", err)
	}
	for _, c := range cases {
		cmp := sufGrad(c[0], "er")
		sup := sufGrad(c[0], "est")

		if cmp != c[1] || sup != c[2] {
			t.Errorf("Failed for '%s': expected '%s' - '%s', got '%s' - '%s'", c[0], c[1], c[2], cmp, sup)
		}
	}
}

// Tests superlative. Fails if a malformed superlative form is returned.
// If the word does not exist in the database, the test attempts
// to transform it as FT_REGULAR.
func TestSuperlative(t *testing.T) {
	type testCase struct {
		Input    string    `json:"input"`
		WC       WordClass `json:"word_class"`
		Expected string    `json:"expected"`
	}

	var cases []testCase
	if err := tests.ReadData("TestSuperlative.json", &cases); err != nil {
		t.Fatalf("Failed loading test data: %v", err)
	}

	gen, err := DefaultGenerator(nil)
	if err != nil {
		t.Fatalf("Failed: DefaultGenerator returned an error: %v", err)
	}

	for _, c := range cases {
		word, err := gen.Find(c.Input, c.WC)
		if err != nil {
			t.Logf("Test case '%s' (WordClass %d) does not exist in the word database. Assume it is regular and proceed.", c.Input, c.WC)

			word, err = NewWordFromParams(c.Input, 0, nil)
			if err != nil {
				t.Errorf("Failed for '%s' - error from NewWordFromParams: %v", c.Input, err)
			}
		}

		output := superlative(word)

		if output != c.Expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", c.Input, c.Expected, output)
		}
	}
}
