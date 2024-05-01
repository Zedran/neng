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

/*
Handles past tense and gerund transformations for verbs ending with consonant-vowel-consonant sequence. Takes a number of arguments:

  - tenseEnding: '-ing' or '-ed'
  - wi: wordInfo created during earlier processing steps
  - wordExceptions: words whose endings are doubled, regardless of transformation rules based on syllable count and verb endings
*/
func handleCVC(verb, tenseEnding string, wi wordInfo, wordExceptions []string) string {
	if strings.HasSuffix(verb, "c") && verb != "sic" {
		if strings.HasSuffix(verb, "lyric") {
			return verb + tenseEnding
		}
		// If final letter is 'c', add 'k' before tenseEnding
		return verb + "k" + tenseEnding
	}

	if containsString(wordExceptions, verb) {
		// Double the final consonant if verb is an exception to the rules below
		return doubleFinal(verb, tenseEnding)
	}

	if wi.sylCount == 2 {
		if endsWithAny(verb, []string{"en", "er", "et", "om", "on", "or"}) {
			// Do not double the final consonant of bisyllabic verbs with specific endings
			return verb + tenseEnding
		}

		if containsString([]string{"augur", "murmur", "sulphur"}, verb) {
			// Do not double the final consonant of bisyllabic exceptions that end with -ur
			return verb + tenseEnding
		}
	}

	if wi.sylCount > 2 {
		// Do not double the final consonant of verbs consisting of more than 2 syllables
		return verb + tenseEnding
	}

	// Double the final consonant of any other verb
	return doubleFinal(verb, tenseEnding)
}

/* Handles transformation of verbs ending with '-it'. */
func handleIt(verb, tenseEnding string, wi wordInfo) string {
	if strings.HasSuffix(wi.sequence, "vvc") {
		if strings.HasSuffix(verb, "quit") {
			return doubleFinal(verb, tenseEnding)
		}
		return verb + tenseEnding
	}

	if wi.sylCount == 1 {
		return doubleFinal(verb, tenseEnding)
	}

	if endsWithAny(verb, []string{"fit", "mit", "wit"}) {
		if !containsString([]string{"limit", "profit"}, verb) {
			return doubleFinal(verb, tenseEnding)
		}
	}

	return verb + tenseEnding
}

/* Returns gerund form of a verb. */
func gerund(verb string) string {
	if len(verb) <= 2 {
		return verb + "ing"
	}

	wi := getWordInfo(verb)

	if strings.HasSuffix(verb, "it") {
		return handleIt(verb, "ing", wi)
	}

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

	if endsWithAny(verb, []string{"h", "w", "x", "s", "y"}) {
		if strings.HasSuffix(verb, "gas") {
			// Double the ending of 'gas' and its derivatives
			return doubleFinal(verb, "ing")
		}

		return verb + "ing"
	}

	if strings.HasSuffix(wi.sequence, "cvc") {
		return handleCVC(verb, "ing", wi, []string{
			"abet", "abhor", "anagram", "beget", "beset", "curvet", "forget", "inset", "offset", "overrun",
			"regret", "reset", "sublet", "typeset", "underrun", "upset",
		})
	}

	return verb + "ing"
}

/* Returns Past Participle form of a verb. */
func pastParticiple(verb string, verbsIrr [][]string) string {
	verbLine := findIrregular(verb, verbsIrr)
	if verbLine == nil {
		return pastRegular(verb)
	}

	return verbLine[2]
}

/* Returns Past Simple form of a verb. */
func pastSimple(verb string, verbsIrr [][]string) string {
	verbLine := findIrregular(verb, verbsIrr)
	if verbLine == nil {
		return pastRegular(verb)
	}

	return verbLine[1]
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

/* Appends past tense suffix to a regular verb. */
func pastRegular(verb string) string {
	wi := getWordInfo(verb)

	if strings.HasSuffix(verb, "i") {
		return verb + "ed"
	}

	if strings.HasSuffix(verb, "it") {
		return handleIt(verb, "ed", wi)
	}

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

	if endsWithAny(verb, []string{"h", "s", "w", "x"}) {
		if strings.HasSuffix(verb, "gas") {
			// Double the ending of 'gas' and its derivatives
			return doubleFinal(verb, "ed")
		}

		return verb + "ed"
	}

	if strings.HasSuffix(wi.sequence, "cvc") {
		return handleCVC(verb, "ed", wi, []string{"abet", "abhor", "anagram", "curvet", "regret"})
	}

	return verb + "ed"
}
