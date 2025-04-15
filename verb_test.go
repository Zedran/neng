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

// Tests gerund. Fails if improper gerund form of a verb is returned.
func TestGerund(t *testing.T) {
	var cases map[string]string
	if err := tests.ReadData("TestGerund.json", &cases); err != nil {
		t.Fatalf("Failed loading test data: %v", err)
	}

	for input, expected := range cases {
		output := gerund(input)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, expected, output)
		}
	}
}

// Tests pastParticiple. Fails if improper Past Participle form of a verb
// is returned. Handling of regular verbs is only symbolically checked,
// as it is the main goal of TestPastSimpleRegular.
func TestPastParticiple(t *testing.T) {
	var cases map[string]string
	if err := tests.ReadData("TestPastParticiple.json", &cases); err != nil {
		t.Fatalf("Failed loading test data: %v", err)
	}

	gen, err := DefaultGenerator(nil)
	if err != nil {
		t.Fatalf("Failed: DefaultGenerator returned an error: %v", err)
	}

	for input, expected := range cases {
		word, err := gen.Find(input, WC_VERB)
		if err != nil {
			t.Logf("Test case '%s' does not exist in the word database. Assume it is regular and proceed.", input)

			word, err = NewWordFromParams(input, 0, nil)
			if err != nil {
				t.Errorf("Failed for '%s' - error from NewWordFromParams: %v", input, err)
			}
		}

		output := pastParticiple(word)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, expected, output)
		}
	}
}

// Tests pastRegular. Fails if incorrect past tense form of a regular verb
// is returned.
func TestPastRegular(t *testing.T) {
	var cases map[string]string
	if err := tests.ReadData("TestPastRegular.json", &cases); err != nil {
		t.Fatalf("Failed loading test data: %v", err)
	}

	for input, expected := range cases {
		output := pastRegular(input)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, expected, output)
		}
	}
}

// Tests pastSimple. Fails if improper Past Simple form of a verb is returned.
// Handling of regular verbs is only symbolically checked, as it is the main
// goal of TestPastSimpleRegular
func TestPastSimple(t *testing.T) {
	type testCase struct {
		Input    string `json:"input"`
		Expected string `json:"expected"`
		Plural   bool   `json:"plural"`
	}

	var cases []testCase
	if err := tests.ReadData("TestPastSimple.json", &cases); err != nil {
		t.Fatalf("Failed loading test data: %v", err)
	}

	gen, err := DefaultGenerator(nil)
	if err != nil {
		t.Fatalf("Failed: DefaultGenerator returned an error: %v", err)
	}

	for _, c := range cases {
		word, err := gen.Find(c.Input, WC_VERB)
		if err != nil {
			t.Logf("Test case '%s' does not exist in the word database. Assume it is regular and proceed.", c.Input)

			word, err = NewWordFromParams(c.Input, 0, nil)
			if err != nil {
				t.Errorf("Failed for '%s' - error from NewWordFromParams: %v", c.Input, err)
			}
		}

		output := pastSimple(word, c.Plural)

		if output != c.Expected {
			t.Errorf("Failed for '%s' (plural = %v): expected '%s', got '%s'", c.Input, c.Plural, c.Expected, output)
		}
	}
}

// Tests presentSimple. Fails if improper Present Simple form of a verb
// is returned.
func TestPresentSimple(t *testing.T) {
	type testCase struct {
		Input    string `json:"input"`
		Expected string `json:"expected"`
		Plural   bool   `json:"plural"`
	}

	var cases []testCase
	if err := tests.ReadData("TestPresentSimple.json", &cases); err != nil {
		t.Fatalf("Failed loading test data: %v", err)
	}

	for _, c := range cases {
		output := presentSimple(c.Input, c.Plural)

		if output != c.Expected {
			number := "sing."

			if c.Plural {
				number = "pl."
			}
			t.Errorf("Failed for '%s': expected '%s' (%s), got '%s'", c.Input, c.Expected, number, output)
		}
	}
}
