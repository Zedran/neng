package neng

import "strings"

/* Returns gerund form of a verb. */
func gerund(verb string) string {
	if len(verb) <= 2 {
		return verb + "ing"
	}

	if strings.HasSuffix(verb, "r") {
		return handleR(verb, "ing")
	}

	wi := getWordInfo(verb)

	if strings.HasSuffix(verb, "l") && strings.HasSuffix(wi.sequence, "vvc") {
		return handleVVL(verb, "ing")
	}

	if strings.HasSuffix(verb, "it") {
		return handleIt(verb, "ing", wi)
	}

	if strings.HasSuffix(verb, "e") && verb != "ante" {
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

	if endsWithAny(verb, []string{"h", "s", "w", "x", "y"}) {
		if strings.HasSuffix(verb, "gas") {
			// Double the ending of 'gas' and its derivatives
			return doubleFinal(verb, "ing")
		}

		return verb + "ing"
	}

	if strings.HasSuffix(wi.sequence, "cvc") {
		return handleCVC(verb, "ing", wi, []string{
			"beget", "begin", "beset", "forget", "inset", "offset", "overrun",
			"reset", "sublet", "typeset", "underrun", "upset",
		})
	}

	return verb + "ing"
}

/*
Handles past tense and gerund transformations for verbs ending with consonant-vowel-consonant sequence. Takes a number of arguments:

  - tenseEnding: '-ing' or '-ed'
  - wi: wordInfo created during earlier processing steps
  - tenseExceptions: tense-specific words whose endings are doubled, regardless of transformation rules based on syllable count and verb endings
*/
func handleCVC(verb, tenseEnding string, wi wordInfo, tenseExceptions []string) string {
	if strings.HasSuffix(verb, "c") && verb != "sic" {
		if strings.HasSuffix(verb, "lyric") {
			return verb + tenseEnding
		}
		// If final letter is 'c', add 'k' before tenseEnding
		return verb + "k" + tenseEnding
	}

	commonSingleExceptions := []string{
		"batik", "kayak", "orphan", "pyramid", "wedel",
	}

	commonDoubledExceptions := []string{
		"abet", "anagram", "curvet", "regret", "revet", "unpin",
	}

	if contains(commonSingleExceptions, verb) {
		return verb + tenseEnding
	}

	if contains(commonDoubledExceptions, verb) {
		// Double the final consonant for exceptions that are common for past forms and gerund
		return doubleFinal(verb, tenseEnding)
	}

	// nil-safe, omitting check
	if contains(tenseExceptions, verb) {
		// Double the final consonant for exceptions specific to the caller
		return doubleFinal(verb, tenseEnding)
	}

	if strings.HasSuffix(verb, "l") {
		return doubleFinal(verb, tenseEnding)
	}

	if wi.sylCount == 2 {
		if endsWithAny(verb, []string{"en", "et", "in", "om", "on"}) {
			// Do not double the final consonant of bisyllabic verbs with specific endings
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

/* Handles transformation of verbs ending with vowel-vowel-l sequence. */
func handleVVL(verb, tenseEnding string) string {
	if strings.HasSuffix(verb, "uel") || contains([]string{"victual", "vitriol"}, verb) {
		return doubleFinal(verb, tenseEnding)
	}

	return verb + tenseEnding
}

/* Handles transformation of verbs ending with '-it'. */
func handleIt(verb, tenseEnding string, wi wordInfo) string {
	if strings.HasSuffix(wi.sequence, "vvc") {
		if strings.HasSuffix(verb, "quit") {
			// The case of 'acquit' and 'quit'
			return doubleFinal(verb, tenseEnding)
		}
		return verb + tenseEnding
	}

	if wi.sylCount == 1 {
		return doubleFinal(verb, tenseEnding)
	}

	if endsWithAny(verb, []string{"fit", "mit", "wit"}) {
		if !contains([]string{"limit", "profit"}, verb) {
			return doubleFinal(verb, tenseEnding)
		}
	}

	return verb + tenseEnding
}

/* Handles transformation of verbs ending with '-r'. */
func handleR(verb, tenseEnding string) string {
	doubled := []string{
		"abhor", "bar", "bestir", "blur", "bur", "char", "concur",
		"confer", "debar", "demur", "deter", "disbar", "disinter",
		"incur", "jar", "mar", "occur", "par", "prefer", "recur",
		"refer", "scar", "slur", "spar", "spur", "star", "tar",
		"transfer", "unbar", "war",
	}

	if contains(doubled, verb) {
		return doubleFinal(verb, tenseEnding)
	}

	return verb + tenseEnding
}

/* Returns Past Participle form of a verb. */
func pastParticiple(verb string, verbsIrr [][]string) string {
	if verb == "be" {
		return "been"
	}

	verbLine := findIrregular(verb, verbsIrr)
	if verbLine != nil {
		return verbLine[2]
	}

	return pastRegular(verb)
}

/* Appends past tense suffix to a regular verb. */
func pastRegular(verb string) string {
	wi := getWordInfo(verb)

	if endsWithAny(verb, []string{"a", "i"}) {
		return verb + "ed"
	}

	if strings.HasSuffix(verb, "l") && strings.HasSuffix(wi.sequence, "vvc") {
		return handleVVL(verb, "ed")
	}

	if strings.HasSuffix(verb, "it") {
		return handleIt(verb, "ed", wi)
	}

	if strings.HasSuffix(verb, "r") {
		return handleR(verb, "ed")
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
		return handleCVC(verb, "ed", wi, nil)
	}

	return verb + "ed"
}

/* Returns Past Simple form of a verb. */
func pastSimple(verb string, verbsIrr [][]string, plural bool) string {
	if verb == "be" {
		if plural {
			return "were"
		}

		return "was"
	}

	verbLine := findIrregular(verb, verbsIrr)
	if verbLine != nil {
		return verbLine[1]
	}

	return pastRegular(verb)
}

/* Returns Present Simple form of a verb. */
func presentSimple(verb string, plural bool) string {
	if plural {
		if verb == "be" {
			return "are"
		}
		return verb
	}

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
