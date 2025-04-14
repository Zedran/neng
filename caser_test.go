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

import "testing"

// Tests case transformations handled by caser.
func TestCaser(t *testing.T) {
	type testCase struct {
		input    string
		expected string
		function func(string) string
	}

	caser := newCaser()

	cases := []testCase{
		{"LOWER", "lower", caser.toLower},
		{"LoWeR", "lower", caser.toLower},
		{"a man", "A man", caser.toSentence},
		{"A MAN", "A man", caser.toSentence},
		{"title", "Title", caser.toSentence},
		{"title", "Title", caser.toTitle},
		{"tItLe", "Title", caser.toTitle},
		{"a man", "A Man", caser.toTitle},
		{"upper", "UPPER", caser.toUpper},
		{"uPpEr", "UPPER", caser.toUpper},
	}

	for _, c := range cases {
		output := c.function(c.input)

		if output != c.expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", c.input, c.expected, output)
		}
	}
}
