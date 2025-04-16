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
	"fmt"
	"strings"
)

// cmpWord is a comparison function for slices.IsSortedFunc
// and slices.SortFunc calls.
func cmpWord(a, b Word) int {
	return strings.Compare(a.word, b.word)
}

// countSyllables returns a number of syllables in s given the consonant-vowel
// sequence. General accuracy of this function is not very high, especially
// for borrowed words (cafe). It targets specific groups of verbs.
func countSyllables(s, seq string) int {
	if len(s) == 0 {
		return 0
	}

	var (
		prevVowel bool
		count     int
	)

	for _, sc := range seq {
		switch sc {
		case 'v':
			if !prevVowel {
				// Diphthongs and long vowels are part of one syllable
				prevVowel = true
				count++
			}
		default:
			prevVowel = false
		}
	}

	if count > 1 && endsWithAny(s, []string{"eat", "eate", "iate", "uate"}) {
		// When apparent diphthongs are in fact individual vowels
		// and belong to separate syllables
		count++
	}

	if strings.HasSuffix(s, "e") && strings.HasSuffix(seq, "cv") {
		// A final 'e' preceded by a consonant (silent 'e')
		// does not constitute the next syllable
		count--
	}

	if count == 0 {
		return 1
	}

	return count
}

// doubleFinal returns the verb with its final consonant doubled
// and tenseEnding appended.
func doubleFinal(verb, tenseEnding string) string {
	return verb + string(verb[len(verb)-1]) + tenseEnding
}

// endsWithAny returns true if s ends with any element of the suf slice.
func endsWithAny(s string, suf []string) bool {
	for _, suffix := range suf {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}

	return false
}

// getSequence returns a vowel-consonant sequence of s ('word' == 'cvcc').
func getSequence(s string) string {
	var (
		seq    strings.Builder
		vowels string = "aeiou"
	)

	for i, c := range s {
		if i == len(s)-1 && c == 'y' && !strings.ContainsRune(vowels, rune(s[i-1])) {
			// A special case of final 'y' following a consonant
			seq.WriteByte('v')
		} else if strings.ContainsRune(vowels, c) {
			seq.WriteByte('v')
		} else {
			seq.WriteByte('c')
		}
	}

	return seq.String()
}

// parseLines converts lines into a slice of Word.
// Relays an error from NewWord (line formatting).
func parseLines(lines []string) ([]Word, error) {
	words := make([]Word, len(lines))

	for i, ln := range lines {
		w, err := NewWord(ln)
		if err != nil {
			return nil, fmt.Errorf("%d '%s' - incorrect format", i, ln)
		}

		words[i] = w
	}

	return words, nil
}
