package neng

import (
	"slices"
	"strings"
)

/* Returns gerund form of a verb. */
func gerund(verb string) string {
	if slices.Contains([]string{"quiz", "up"}, verb) {
		return doubleFinal(verb, "ing")
	}

	if len(verb) <= 2 {
		return verb + "ing"
	}

	switch verb[len(verb)-1] {
	case 'e':
		switch verb[len(verb)-2] {
		case 'u': // ue
			return verb[:len(verb)-1] + "ing"
		case 'i': // ie
			return verb[:len(verb)-2] + "ying"
		}

		if strings.HasSuffix(getSequence(verb), "cv") && !strings.HasSuffix(verb, "ye") && verb != "ante" {
			// Remove final 'e' if previous letter is consonant other than 'y' and the verb is not 'ante'
			return verb[:len(verb)-1] + "ing"
		}
	case 'y', 'h', 'w', 'x':
		return verb + "ing"
	case 'r':
		return handleR(verb, "ing")
	case 'l':
		if strings.HasSuffix(getSequence(verb), "vvc") {
			return handleVVL(verb, "ing")
		}
	case 's':
		if strings.HasSuffix(verb, "gas") {
			// Double the ending of 'gas' and its derivatives
			return doubleFinal(verb, "ing")
		}
		return verb + "ing"
	}

	if strings.HasSuffix(verb, "it") {
		return handleIt(verb, "ing")
	}

	seq := getSequence(verb)

	if strings.HasSuffix(seq, "cvc") {
		return handleCVC(verb, "ing", seq, []string{
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
func handleCVC(verb, tenseEnding string, seq string, tenseExceptions []string) string {
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
		"abet", "anagram", "curvet", "regret", "revet", "underpin", "unpin",
	}

	if slices.Contains(commonSingleExceptions, verb) {
		return verb + tenseEnding
	}

	if slices.Contains(commonDoubledExceptions, verb) {
		// Double the final consonant for exceptions that are common for past forms and gerund
		return doubleFinal(verb, tenseEnding)
	}

	// nil-safe, omitting check
	if slices.Contains(tenseExceptions, verb) {
		// Double the final consonant for exceptions specific to the caller
		return doubleFinal(verb, tenseEnding)
	}

	if strings.HasSuffix(verb, "l") {
		return doubleFinal(verb, tenseEnding)
	}

	sylCount := countSyllables(verb, seq)

	if sylCount == 2 {
		if endsWithAny(verb, []string{"en", "et", "in", "om", "on"}) {
			// Do not double the final consonant of bisyllabic verbs with specific endings
			return verb + tenseEnding
		}
	}

	if sylCount > 2 {
		// Do not double the final consonant of verbs consisting of more than 2 syllables
		return verb + tenseEnding
	}

	// Double the final consonant of any other verb
	return doubleFinal(verb, tenseEnding)
}

/* Handles transformation of verbs ending with '-it'. */
func handleIt(verb, tenseEnding string) string {
	seq := getSequence(verb)

	if strings.HasSuffix(seq, "vvc") {
		if strings.HasSuffix(verb, "quit") {
			// The case of 'acquit' and 'quit'
			return doubleFinal(verb, tenseEnding)
		}
		return verb + tenseEnding
	}

	if countSyllables(verb, seq) == 1 {
		return doubleFinal(verb, tenseEnding)
	}

	if endsWithAny(verb, []string{"fit", "mit", "wit"}) {
		if !slices.Contains([]string{"limit", "profit"}, verb) {
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
		"refer", "scar", "slur", "spar", "spur", "star", "stir", "tar",
		"transfer", "unbar", "war",
	}

	if slices.Contains(doubled, verb) {
		return doubleFinal(verb, tenseEnding)
	}

	return verb + tenseEnding
}

/* Handles transformation of verbs ending with vowel-vowel-l sequence. */
func handleVVL(verb, tenseEnding string) string {
	if strings.HasSuffix(verb, "uel") || slices.Contains([]string{"victual", "vitriol"}, verb) {
		return doubleFinal(verb, tenseEnding)
	}

	return verb + tenseEnding
}

/* Returns Past Participle form of a verb. */
func pastParticiple(word *Word) string {
	if word.ft == FT_IRREGULAR {
		return (*word.irr)[1]
	}

	if word.word == "be" {
		return "been"
	}

	return pastRegular(word.word)
}

/* Appends past tense suffix to a regular verb. */
func pastRegular(verb string) string {
	if slices.Contains([]string{"quiz", "up"}, verb) {
		return doubleFinal(verb, "ed")
	}

	switch verb[len(verb)-1] {
	case 'e':
		return verb + "d"
	case 'r':
		return handleR(verb, "ed")
	case 'h', 'w', 'o', 'x', 'a', 'i', 'u':
		return verb + "ed"
	case 'l':
		if strings.HasSuffix(getSequence(verb), "vvc") {
			return handleVVL(verb, "ed")
		}
	case 'y':
		if strings.HasSuffix(getSequence(verb), "v") {
			return verb[:len(verb)-1] + "ied"
		}
		return verb + "ed"
	case 's':
		if strings.HasSuffix(verb, "gas") {
			// Double the ending of 'gas' and its derivatives
			return doubleFinal(verb, "ed")
		}
		return verb + "ed"
	}

	if strings.HasSuffix(verb, "it") {
		return handleIt(verb, "ed")
	}

	seq := getSequence(verb)

	if strings.HasSuffix(seq, "cvc") {
		return handleCVC(verb, "ed", seq, nil)
	}

	return verb + "ed"
}

/* Returns Past Simple form of a verb. */
func pastSimple(word *Word, plural bool) string {
	if word.ft == FT_IRREGULAR {
		return (*word.irr)[0]
	}

	if word.word == "be" {
		if plural {
			return "were"
		}

		return "was"
	}

	return pastRegular(word.word)
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

	switch verb[len(verb)-1] {
	case 'y':
		if strings.HasSuffix(seq, "v") {
			return verb[:len(verb)-1] + "ies"
		}
	case 's', 'x':
		return verb + "es"
	case 'o':
		if strings.HasSuffix(seq, "cv") {
			return verb + "es"
		}
	case 'z':
		if verb == "quiz" {
			return doubleFinal(verb, "es")
		}
		return verb + "es"
	}

	if endsWithAny(verb, []string{"ch", "sh"}) {
		return verb + "es"
	}

	return verb + "s"
}
