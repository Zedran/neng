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
