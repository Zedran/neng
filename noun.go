package neng

import "strings"

/* Returns plural form of a noun. */
func plural(noun string, nounsIrr [][]string) string {
	nounLine := findIrregular(noun, nounsIrr)
	if nounLine != nil {
		return nounLine[1]
	}

	if endsWithAny(noun, []string{"craft", "fish", "hertz", "ics", "mania", "ness"}) {
		return noun
	}

	if endsWithAny(noun, []string{"ff", "tto"}) {
		return noun + "s"
	}

	if endsWithAny(noun, []string{"ch", "fix", "sh", "virus"}) {
		return noun + "es"
	}

	switch noun[len(noun)-2:] {
	case "ex", "ix":
		return noun[:len(noun)-2] + "ices"
	case "fe":
		return noun[:len(noun)-2] + "ves"
	case "is":
		return noun[:len(noun)-2] + "es"
	case "um":
		return noun[:len(noun)-2] + "a"
	case "us":
		return noun[:len(noun)-2] + "i"
	}

	seq := getSequence(noun)

	switch noun[len(noun)-1] {
	case 'f':
		return noun[:len(noun)-1] + "ves"
	case 'o', 's', 'x':
		return noun + "es"
	case 'y':
		if strings.HasSuffix(seq, "v") {
			return noun[:len(noun)-1] + "ies"
		}
	case 'z':
		if strings.HasSuffix(seq, "vc") {
			return doubleFinal(noun, "es")
		}
		return noun + "es"
	}

	return noun + "s"
}
