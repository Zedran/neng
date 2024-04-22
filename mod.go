package neng

import "strings"

// Modification parameter for a generated word
type Mod uint16

const (
	// Create gerund form of a verb (-ing)
	MOD_GERUND         Mod = iota

	// Add Present Simple suffix to a verb (-s, -es)
	MOD_PRESENT_SIMPLE

	// Add Past Simple suffix to a verb or substitute its irregular form
	MOD_PAST_SIMPLE

	// Add Past Simple suffix to a regular verb or substitute Past Participle form to an irregular one
	MOD_PAST_PARTICIPLE
)

/* Returns gerund form of a verb. */
func gerund(verb string) string {
	if len(verb) == 2 {
		return verb + "ing"
	}
	
	if strings.HasSuffix(verb, "e") && !isVowel([]rune(verb)[len(verb) - 2]) {
		// Remove final 'e' if previous letter is consonant
		return verb[:len(verb) - 1] + "ing"
	}
	
	if !endsWithAny(verb, []string{"h", "w", "x", "y"}) {
		// Double the consonant if the sequence of final letters is 'consonant-vowel-consonant'
		if getSequence(verb[len(verb) - 3:]) == "cvc" {
			// If final letter is 'c', add 'k'
			if verb[len(verb) - 1] == 'c' {
				return verb + "king"
			}
			// Double any other letter
			return verb + string(verb[len(verb) - 1]) + "ing"
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

	if strings.HasSuffix(verb, "y") {
		return verb[:len(verb) - 1] + "ies"
	} else if endsWithAny(verb, []string{"ch", "s", "sh", "x"}) {
		return verb + "es"
	}

	return verb + "s"
}

/* Appends Past Simple suffix to a regular verb. */
func pastSimpleRegular(verb string) string {
	if isVowel([]rune(verb)[len(verb) - 1]) {
		return verb + "d"
	}

	if strings.HasSuffix(verb, "y") {
		return verb[:len(verb) - 1] + "ied"
	}
	
	if !endsWithAny(verb, []string{"h", "w", "x"})  {
		// Double the consonant if the sequence of final letters is 'consonant-vowel-consonant'
		if getSequence(verb[len(verb) - 3:]) == "cvc" {
			// If final letter is 'c', add 'k'
			if verb[len(verb) - 1] == 'c' {
				return verb + "ked"
			}
		}
		// Double any other letter
		return verb + string(verb[len(verb) - 1]) + "ed"
	}

	return verb + "ed"
}
