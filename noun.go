package neng

import "strings"

/* Returns plural form of a noun. */
func plural(word *Word) string {
	if word.ft == FT_PLURAL_ONLY {
		return word.word
	}

	if word.ft == FT_IRREGULAR {
		return (*word.irr)[0]
	}

	noun := word.word

	if endsWithAny(noun, []string{"fish", "ics", "gs", "craft"}) {
		return noun
	}

	if endsWithAny(noun, []string{"sh", "ch"}) {
		return noun + "es"
	}

	if strings.HasSuffix(noun, "man") {
		return noun[:len(noun)-2] + "en"
	}

	if endsWithAny(noun, []string{"life", "knife", "wife"}) {
		return noun[:len(noun)-2] + "ves"
	}

	switch noun[len(noun)-2:] {
	case "um":
		return noun[:len(noun)-2] + "a"
	}

	seq := getSequence(noun)

	switch noun[len(noun)-1] {
	case 'y':
		if strings.HasSuffix(seq, "v") {
			return noun[:len(noun)-1] + "ies"
		}
	case 's':
		if endsWithAny(noun, []string{"sis", "xis"}) {
			return noun[:len(noun)-2] + "es"
		}
		if endsWithAny(noun, []string{"cirrus", "cumulus", "nimbus", "stratus"}) {
			return noun[:len(noun)-2] + "i"
		}
		return noun + "es"
	case 'f':
		if endsWithAny(noun, []string{"leaf", "elf", "arf", "alf", "wolf", "loaf"}) {
			return noun[:len(noun)-1] + "ves"
		}
	case 'x':
		if endsWithAny(noun, []string{"dex", "dix", "fex", "pex", "rix", "tex"}) {
			return noun[:len(noun)-2] + "ices"
		}
		return noun + "es"
	case 'z':
		if strings.HasSuffix(seq, "vc") {
			return doubleFinal(noun, "es")
		}
		return noun + "es"
	}

	return noun + "s"
}
