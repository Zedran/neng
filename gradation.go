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

// comparative returns a comparative form of an adjective or an adverb
// (good -> better).
func comparative(word *Word) string {
	switch word.ft {
	case FT_IRREGULAR:
		return (*word.irr)[0]
	case FT_SUFFIXED:
		return sufGrad(word.word, "er")
	default:
		return "more " + word.word
	}
}

// sufGrad returns comparative or superlative form of those adjectives to which
// suffix is appended during gradation process.
func sufGrad(a, suf string) string {
	switch a[len(a)-1] {
	case 'y':
		if strings.HasSuffix(a, "ey") {
			return a[:len(a)-2] + "i" + suf
		}
		return a[:len(a)-1] + "i" + suf
	case 'b', 'd', 'g', 'm', 'n', 'p', 't':
		if strings.HasSuffix(getSequence(a), "cvc") {
			return doubleFinal(a, suf)
		}
	case 'e':
		return a[:len(a)-1] + suf
	}

	return a + suf
}

// superlative returns a superlative form of an adjective or an adverb
// (good -> best).
func superlative(word *Word) string {
	switch word.ft {
	case FT_IRREGULAR:
		return (*word.irr)[1]
	case FT_SUFFIXED:
		return sufGrad(word.word, "est")
	default:
		return "most " + word.word
	}
}
