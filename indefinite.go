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

import "strings"

// indefinite returns the word prefixed with an indefinite article.
func indefinite(word string) string {
	switch word[0] {
	case 'a', 'i', 'e', 'o':
		return "an " + word
	case 'u':
		return indefiniteU(word)
	}
	return "a " + word
}

// indefiniteU prefixes the word beginning with 'u-' with an indefinite article.
func indefiniteU(word string) string {
	if len(word) == 1 {
		return "a " + word
	}

	switch word[1] {
	case 'n':
		if strings.HasPrefix(word, "uni") {
			if len(word) == 3 || !strings.ContainsRune("nmdr", rune(word[3])) || word == "unimodal" || word == "uninominal" {
				// unin-, unim-, unid-, unir- (for a single 'unironed')
				return "a " + word
			}
		}
	case 's':
		if !strings.HasPrefix(word, "ush") {
			return "a " + word
		}
	case 't':
		if !strings.HasPrefix(word, "utt") {
			return "a " + word
		}
	}
	return "an " + word
}
