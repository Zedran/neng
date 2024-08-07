package neng

import (
	"bytes"
	"math/rand/v2"
	"strings"
)

/*
Returns number of syllables in s given consonant-vowel sequence seq.
Accuracy of this function is uncertain, especially for borrowed words (cafe).
*/
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
		// When apparent diphthongs are in fact individual vowels and belong to separate syllables
		count++
	}

	if strings.HasSuffix(s, "e") && strings.HasSuffix(seq, "cv") {
		// A final 'e' preceded by a consonant (silent 'e') does not constitute the next syllable
		count--
	}

	if count == 0 {
		return 1
	}

	return count
}

/* Doubles the final consonant of a verb and appends tenseEnding to it. */
func doubleFinal(verb, tenseEnding string) string {
	return verb + string(verb[len(verb)-1]) + tenseEnding
}

/* Returns true if s ends with any element of suf slice. */
func endsWithAny(s string, suf []string) bool {
	for _, suffix := range suf {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}

	return false
}

/* For irregular words, returns slice with word forms from wordsIrr. For regular words, returns nil. */
func findIrregular(word string, wordsIrr [][]string) []string {
	for _, iw := range wordsIrr {
		if iw[0] == word {
			return iw
		}
	}

	return nil
}

/* Returns a representation of vowel-consonant sequence in s ('word' == 'cvcc'). */
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

/* Loads a word list from the embedded path. Returns error if the file is not found. */
func loadLines(path string) ([][]byte, error) {
	stream, err := efs.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return bytes.Split(stream, []byte("\n")), nil
}

/* Parses the loaded word list into a slice of word struct pointers. Relays error from NewWord (line formatting). */
func parseLines(lines [][]byte) ([]*word, error) {
	words := make([]*word, len(lines))

	for i, ln := range lines {
		w, err := NewWord(ln)
		if err != nil {
			return nil, err
		}

		words[i] = w
	}

	return words, nil
}

/* Returns a random word from s. */
func randItem(s []*word) *word {
	return s[rand.IntN(len(s))]
}
