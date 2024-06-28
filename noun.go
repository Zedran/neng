package neng

import (
	"slices"
	"strings"
)

/* Returns plural form of a noun given slices of plural-only nouns and irregularly pluralized nouns. */
func plural(noun string, nounsPlO []string, nounsIrr [][]string) string {
	if slices.Contains(nounsPlO, noun) {
		return noun
	}

	nounLine := findIrregular(noun, nounsIrr)
	if nounLine != nil {
		return nounLine[1]
	}

	if endsWithAny(noun, []string{"fish", "ics", "gs", "craft", "hertz"}) {
		return noun
	}

	if endsWithAny(noun, []string{"sh", "ch", "fix"}) {
		return noun + "es"
	}

	if strings.HasSuffix(noun, "man") {
		return noun[:len(noun)-2] + "en"
	}

	switch noun[len(noun)-2:] {
	case "um":
		return noun[:len(noun)-2] + "a"
	case "is":
		return noun[:len(noun)-2] + "es"
	case "ex", "ix":
		return noun[:len(noun)-2] + "ices"
	case "fe":
		return noun[:len(noun)-2] + "ves"
	}

	seq := getSequence(noun)

	switch noun[len(noun)-1] {
	case 'y':
		if strings.HasSuffix(seq, "v") {
			return noun[:len(noun)-1] + "ies"
		}
	case 's', 'x':
		if endsWithAny(noun, []string{"cirrus", "cumulus", "nimbus", "stratus"}) {
			return noun[:len(noun)-2] + "i"
		}
		return noun + "es"
	case 'f':
		if !strings.HasSuffix(noun, "ff") {
			return noun[:len(noun)-1] + "ves"
		}
	case 'z':
		if strings.HasSuffix(seq, "vc") {
			return doubleFinal(noun, "es")
		}
		return noun + "es"
	}

	return noun + "s"
}
