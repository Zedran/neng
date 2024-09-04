package neng

import "strings"

/* Returns plural form of a noun. */
func plural(word *Word) string {
	switch word.ft {
	case FT_IRREGULAR:
		return (*word.irr)[0]
	case FT_PLURAL_ONLY:
		return word.word
	}

	noun := word.word

	switch noun[len(noun)-1] {
	case 'e':
		if endsWithAny(noun, []string{"life", "knife", "wife"}) {
			return noun[:len(noun)-2] + "ves"
		}
	case 'y':
		if strings.HasSuffix(getSequence(noun), "v") {
			return noun[:len(noun)-1] + "ies"
		}
	case 's':
		if endsWithAny(noun, []string{"ics", "gs"}) {
			return noun
		}
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
		if strings.HasSuffix(getSequence(noun), "vc") {
			return doubleFinal(noun, "es")
		}
		return noun + "es"
	}

	switch noun[len(noun)-2:] {
	case "um":
		return noun[:len(noun)-2] + "a"
	case "sh":
		if strings.HasSuffix(noun, "fish") {
			return noun
		}
		return noun + "es"
	case "ch":
		return noun + "es"
	}

	if strings.HasSuffix(noun, "man") {
		return noun[:len(noun)-2] + "en"
	}

	if strings.HasSuffix(noun, "craft") {
		return noun
	}

	return noun + "s"
}
