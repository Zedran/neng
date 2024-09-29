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
	switch word[1] {
	case 'n':
		if len(word) >= 3 && strings.HasPrefix(word, "uni") {
			if !strings.ContainsRune("nmdr", rune(word[3])) || word == "unimodal" || word == "uninominal" {
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
