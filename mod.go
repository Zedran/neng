package neng

import "strings"

// Modification parameter for a generated word
type Mod uint8

const (
	// Add Past Simple suffix to a verb or substitute its irregular form
	MOD_PAST_SIMPLE Mod = iota

	// Add Past Simple suffix to a regular verb or substitute Past Participle form to an irregular one
	MOD_PAST_PARTICIPLE

	// Add Present Simple suffix to a verb (-s, -es)
	MOD_PRESENT_SIMPLE

	// Create gerund form of a verb (-ing)
	MOD_GERUND

	// Transform a word to lower case
	MOD_CASE_LOWER

	// Transform a word to Title Case
	MOD_CASE_TITLE

	// Transform a word to UPPER CASE
	MOD_CASE_UPPER
)

/* Returns gerund form of a verb. */
func gerund(verb string) string {
	if len(verb) <= 2 {
		return verb + "ing"
	}

	if verb == "quit" {
		return "quitting"
	}

	wi := getWordInfo(verb)

	if strings.HasSuffix(verb, "e") {
		if wi.sequence[len(wi.sequence)-2] == 'c' && verb[len(verb)-2] != 'y' {
			// Remove final 'e' if previous letter is consonant other than 'y'
			return verb[:len(verb)-1] + "ing"
		}

		switch verb[len(verb)-2] {
		case 'u': // ue
			return verb[:len(verb)-1] + "ing"
		case 'i': // ie
			return verb[:len(verb)-2] + "ying"
		}
	}

	if !endsWithAny(verb, []string{"h", "w", "x", "y"}) {
		if strings.HasSuffix(wi.sequence, "cvc") {
			// Double the consonant if the sequence of final letters is 'consonant-vowel-consonant'
			if strings.HasSuffix(verb, "c") {
				// If final letter is 'c', add 'k' instead
				return verb + "king"
			}

			if containsString([]string{
				"abet", "beget", "beset", "curvet", "forget", "inset", "offset", "overrun",
				"recommit", "regret", "reset", "sublet", "typeset", "underrun", "upset",
			}, verb) {
				// Seemingly singular exceptions to the rules below
				return verb + string(verb[len(verb)-1]) + "ing"
			}

			if wi.sylCount == 2 && endsWithAny(verb, []string{"en", "er", "et", "on", "or"}) {
				return verb + "ing"
			}

			if wi.sylCount > 2 {
				return verb + "ing"
			}

			// Double any other letter
			return verb + string(verb[len(verb)-1]) + "ing"
		}
	}

	return verb + "ing"
}

/* Returns Past Participle form of a verb. */
func pastParticiple(verb string, verbsIrr [][]string) string {
	verbLine := findIrregular(verb, verbsIrr)
	if verbLine == nil {
		return pastSimpleRegular(verb)
	}

	return strings.Replace(verb, verbLine[0], verbLine[2], 1)
}

/* Returns Past Simple form of a verb. */
func pastSimple(verb string, verbsIrr [][]string) string {
	verbLine := findIrregular(verb, verbsIrr)
	if verbLine == nil {
		return pastSimpleRegular(verb)
	}

	return strings.Replace(verb, verbLine[0], verbLine[1], 1)
}

/* Returns Present Simple form of a verb. */
func presentSimple(verb string) string {
	switch verb {
	case "be":
		return "is"
	case "have":
		return "has"
	}

	seq := getSequence(verb)

	if strings.HasSuffix(verb, "o") && strings.HasSuffix(seq, "cv") {
		return verb + "es"
	} else if strings.HasSuffix(verb, "y") && strings.HasSuffix(seq, "v") {
		return verb[:len(verb)-1] + "ies"
	} else if endsWithAny(verb, []string{"ch", "s", "sh", "x", "z"}) {
		return verb + "es"
	}

	return verb + "s"
}

/* Appends Past Simple suffix to a regular verb. */
func pastSimpleRegular(verb string) string {
	wi := getWordInfo(verb)

	if strings.HasSuffix(verb, "y") {
		if strings.HasSuffix(wi.sequence, "v") {
			return verb[:len(verb)-1] + "ied"
		}
		return verb + "ed"
	}

	if strings.HasSuffix(wi.sequence, "v") {
		if strings.HasSuffix(verb, "o") {
			return verb + "ed"
		}
		return verb + "d"
	}

	if !endsWithAny(verb, []string{"h", "w", "x"}) {
		if strings.HasSuffix(wi.sequence, "cvc") {
			// Double the consonant if the sequence of final letters is 'consonant-vowel-consonant'
			if strings.HasSuffix(verb, "c") {
				// If final letter is 'c', add 'k' instead
				return verb + "ked"
			}

			if containsString([]string{"abet", "curvet", "recommit", "regret"}, verb) {
				// Seemingly singular exceptions to the rules below
				return verb + string(verb[len(verb)-1]) + "ed"
			}

			if wi.sylCount == 2 && endsWithAny(verb, []string{"en", "er", "et", "on", "or"}) {
				return verb + "ed"
			}

			if wi.sylCount > 2 {
				// Do not double if the word is more than 2 syllables long
				return verb + "ed"
			}

			// Double any other letter
			return verb + string(verb[len(verb)-1]) + "ed"
		}
	}

	return verb + "ed"
}
