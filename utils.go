package neng

import (
	"math/rand"
	"strings"
)

/* Returns true if s ends with any element of suf slice. */
func endsWithAny(s string, suf []string) bool {
	for _, suffix := range suf {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}

	return false
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

/* For irregular verbs, returns slice with verb forms from verbsIrr. For regular verbs, returns nil. */
func findIrregular(verb string, verbsIrr [][]string) []string {
	for _, iv := range verbsIrr {
		if iv[0] == verb {
			return iv
		}
	}

	return nil
}

/*
Calls loadWords to read lines from efs, splits those lines into a slices of verb forms
and returns [lines][forms]string.
*/
func loadIrregularVerbs(path string) ([][]string, error) {
	lines, err := loadWords(path)
	if err != nil {
		return nil, err
	}

	verbsIrr := make([][]string, len(lines))

	for i, ln := range lines {
		verbsIrr[i] = strings.Split(ln, ",")
	}

	return verbsIrr, nil
}

/* Loads a word list from path. Returns error if the file is not found. */
func loadWords(path string) ([]string, error) {
	stream, err := efs.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(stream), "\n"), nil
}

/* Returns a random item from s. */
func randItem(s []string) string {
	return s[rand.Intn(len(s))]
}
