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

/* Returns a representation of vowel-consonant sequence in a word ('word' == 'cvcc'). Required for gerund and Past Simple formation. */
func getSequence(s string) string {
	var (
		seq strings.Builder
		sr = []rune(s)
	)

	for i := range s {
		if isVowel(sr[i]) {
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

/* Returns true if r is a vowel. */
func isVowel(r rune) bool {
	return strings.ContainsRune("aeiou", r)
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
