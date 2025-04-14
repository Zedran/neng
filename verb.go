// neng -- Non-Extravagant Name Generator
// Copyright (C) 2024  Wojciech Głąb (github.com/Zedran)
//
// This file is part of neng.
//
// neng is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3 only.
//
// neng is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with neng.  If not, see <https://www.gnu.org/licenses/>.

package neng

import (
	"slices"
	"strings"
)

// gerund returns a gerund form of a verb.
func gerund(verb string) string {
	if len(verb) <= 2 {
		if verb == "up" {
			return doubleFinal(verb, "ing")
		}
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

	if verb == "quiz" {
		return doubleFinal(verb, "ing")
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

// handleCVC performs past tense and gerund transformations for verbs ending
// with consonant-vowel-consonant sequence. Takes a number of arguments:
//
//   - tenseEnding: '-ing' or '-ed'
//   - seq: vowel-consonant sequence
//   - tenseExceptions: tense-specific words whose endings are doubled,
//     regardless of transformation rules based on syllable count
//     and verb endings
func handleCVC(verb, tenseEnding, seq string, tenseExceptions []string) string {
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
		// Double the final consonant for exceptions that are common
		// for past forms and gerund
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
			// Do not double the final consonant of bisyllabic verbs
			// with specific endings
			return verb + tenseEnding
		}
	}

	if sylCount > 2 {
		// Do not double the final consonant of verbs consisting
		// of more than 2 syllables
		return verb + tenseEnding
	}

	// Double the final consonant of any other verb
	return doubleFinal(verb, tenseEnding)
}

// handleIt transforms verbs ending with '-it'.
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
		if verb != "limit" && verb != "profit" {
			return doubleFinal(verb, tenseEnding)
		}
	}

	return verb + tenseEnding
}

// handleR transforms verbs ending with '-r'.
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

// handleVVL transforms verbs ending with vowel-vowel-l sequence.
func handleVVL(verb, tenseEnding string) string {
	if strings.HasSuffix(verb, "uel") || verb == "victual" || verb == "vitriol" {
		return doubleFinal(verb, tenseEnding)
	}

	return verb + tenseEnding
}

// pastParticiple returns Past Participle form of a verb.
func pastParticiple(word *Word) string {
	if word.ft == FT_IRREGULAR {
		return (*word.irr)[1]
	}

	if word.word == "be" {
		return "been"
	}

	return pastRegular(word.word)
}

// pastRegular appends past tense suffix to a regular verb.
func pastRegular(verb string) string {
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

	if verb == "quiz" || verb == "up" {
		return doubleFinal(verb, "ed")
	}

	seq := getSequence(verb)

	if strings.HasSuffix(seq, "cvc") {
		return handleCVC(verb, "ed", seq, nil)
	}

	return verb + "ed"
}

// pastSimple returns Past Simple form of a verb.
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

// presentSimple returns Present Simple form of a verb.
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

	switch verb[len(verb)-1] {
	case 'y':
		if strings.HasSuffix(getSequence(verb), "v") {
			return verb[:len(verb)-1] + "ies"
		}
	case 's', 'x':
		return verb + "es"
	case 'o':
		if strings.HasSuffix(getSequence(verb), "cv") {
			return verb + "es"
		}
	case 'z':
		if verb == "quiz" {
			return doubleFinal(verb, "es")
		}
		return verb + "es"
	}

	if strings.HasSuffix(verb, "ch") || strings.HasSuffix(verb, "sh") {
		return verb + "es"
	}

	return verb + "s"
}
