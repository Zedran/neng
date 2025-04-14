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
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// caser handles case transformations.
type caser struct {
	lower cases.Caser
	title cases.Caser
	upper cases.Caser
}

// toLower transforms word to lower case.
func (c *caser) toLower(word string) string {
	return c.lower.String(word)
}

// toSentence transforms the first word in a space-separated sequence
// to title case and everything that follows to lower case.
func (c *caser) toSentence(words string) string {
	space := strings.Index(words, " ")
	if space == -1 {
		return c.toTitle(words)
	}
	return c.toTitle(words[:space]) + c.toLower(words[space:])
}

// toTitle transforms word to title case.
func (c *caser) toTitle(word string) string {
	return c.title.String(word)
}

// toUpper transforms word to upper case.
func (c *caser) toUpper(word string) string {
	return c.upper.String(word)
}

// newCaser returns a pointer to new caser struct.
func newCaser() *caser {
	return &caser{
		lower: cases.Lower(language.English),
		title: cases.Title(language.English),
		upper: cases.Upper(language.English),
	}
}
