package neng

import "strings"

/* Returns plural form of a noun. */
func plural(noun string, nounsIrr [][]string) string {
	nounLine := findIrregular(noun, nounsIrr)
	if nounLine != nil {
		return nounLine[1]
	}

	if endsWithAny(noun, []string{"fish", "ics", "gs", "craft", "hertz"}) {
		return noun
	}

	if endsWithAny(noun, []string{"sh", "ch", "fix", "virus"}) {
		return noun + "es"
	}

	if endsWithAny(noun, []string{"ff", "io", "tto"}) {
		return noun + "s"
	}

	switch noun[len(noun)-2:] {
	case "um":
		return noun[:len(noun)-2] + "a"
	case "us":
		return noun[:len(noun)-2] + "i"
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
	case 's', 'o', 'x':
		return noun + "es"
	case 'f':
		return noun[:len(noun)-1] + "ves"
	case 'z':
		if strings.HasSuffix(seq, "vc") {
			return doubleFinal(noun, "es")
		}
		return noun + "es"
	}

	return noun + "s"
}
