package neng

import "strings"

/* Returns comparative form of an adjective or an adverb (good -> better). */
func comparative(a string, adjIrr [][]string, adjSuf []string) string {
	aLine := findIrregular(a, adjIrr)
	if aLine != nil {
		return aLine[1]
	}

	if contains(adjSuf, a) {
		return sufGrad(a, "er")
	}

	return "more " + a
}

/* Returns comparative or superlative form of those adjectives to which suffix is appended during gradation process. */
func sufGrad(a, suf string) string {
	switch a[len(a)-1] {
	case 'y':
		if strings.HasSuffix(a, "ey") {
			return a[:len(a)-2] + "i" + suf
		}
		return a[:len(a)-1] + "i" + suf

	case 'b', 'd', 'g', 'm', 'n', 'p', 't':
		seq := getSequence(a)
		if strings.HasSuffix(seq, "cvc") {
			return doubleFinal(a, suf)
		}
	case 'e':
		return a[:len(a)-1] + suf
	}

	return a + suf
}

/* Returns superlative form of an adjective or an adverb (good -> best). */
func superlative(a string, adjIrr [][]string, adjSuf []string) string {
	aLine := findIrregular(a, adjIrr)
	if aLine != nil {
		return aLine[2]
	}

	if contains(adjSuf, a) {
		return sufGrad(a, "est")
	}

	return "most " + a
}
