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

// Tests whether countSyllables returns a correct number of syllables.
func TestCountSyllables(t *testing.T) {
	var cases map[string]int
	if err := tests.ReadData("TestCountSyllables.json", &cases); err != nil {
		t.Fatalf("Failed loading test data: %v", err)
	}

	for input, expected := range cases {
		seq := getSequence(input)
		output := countSyllables(input, seq)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%d', got '%d'", input, expected, output)
		}
	}
}

// Tests getSequence.
func TestGetSequence(t *testing.T) {
	var cases map[string]string
	if err := tests.ReadData("TestGetSequence.json", &cases); err != nil {
		t.Fatalf("Failed loading test data: %v", err)
	}

	for input, expected := range cases {
		output := getSequence(input)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, expected, output)
		}
	}
}
