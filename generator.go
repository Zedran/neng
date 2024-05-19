package neng

import (
	"slices"
	"strings"
)

/* Generates random phrases or words. */
type Generator struct {
	adjectives []string
	adverbs    []string
	nouns      []string
	verbs      []string
	nounsIrr   [][]string
	verbsIrr   [][]string
	caser      *caser
}

/*
Generates a single random adjective and transforms it according to mods.
Returns an error if an undefined Mod is received.
*/
func (gen *Generator) Adjective(mods ...Mod) (string, error) {
	return gen.Transform(randItem(gen.adjectives), mods...)
}

/*
Generates a single random adverb and transforms it according to mods.
Returns an error if an undefined Mod is received.
*/
func (gen *Generator) Adverb(mods ...Mod) (string, error) {
	return gen.Transform(randItem(gen.adverbs), mods...)
}

/*
Transforms a word according to specified mods. Not all mods are compatible with every
part of speech. Compatibility is not checked. Returns an error if an undefined Mod is received.
*/
func (gen *Generator) Transform(word string, mods ...Mod) (string, error) {
	var (
		caseTransformation func(string) string
		pluralMod          bool
		verbMod            bool
	)

	// Ensure MOD_PLURAL is processed first
	slices.Sort(mods)

	for _, m := range mods {
		switch m {
		case MOD_PLURAL:
			pluralMod = true
		case MOD_GERUND:
			verbMod = true
			word = gerund(word)
		case MOD_PRESENT_SIMPLE:
			verbMod = true
			word = presentSimple(word, pluralMod)
		case MOD_PAST_SIMPLE:
			verbMod = true
			word = pastSimple(word, gen.verbsIrr)
		case MOD_PAST_PARTICIPLE:
			verbMod = true
			word = pastParticiple(word, gen.verbsIrr)
		case MOD_CASE_LOWER:
			caseTransformation = gen.caser.toLower
		case MOD_CASE_TITLE:
			caseTransformation = gen.caser.toTitle
		case MOD_CASE_UPPER:
			caseTransformation = gen.caser.toUpper
		default:
			return "", errUndefinedMod
		}
	}

	if pluralMod && !verbMod {
		word = plural(word, gen.nounsIrr)
	}

	if caseTransformation != nil {
		word = caseTransformation(word)
	}

	return word, nil
}

/*
Generates a single random noun and transforms it according to mods.
Returns an error if an undefined Mod is received.
*/
func (gen *Generator) Noun(mods ...Mod) (string, error) {
	return gen.Transform(randItem(gen.nouns), mods...)
}

/*
Generates a phrase given the pattern.

Syntax:

	Insertion:
		%% - inserts '%' sign
		%a - inserts a random adjective
		%m - inserts a random adverb
		%n - inserts a random noun
		%v - inserts a random verb

	Transformation:
		%2 - transforms a verb into its Past Simple form (2nd form)
		%3 - transforms a verb into its Past Participle form (3rd form)
		%N - transforms a verb into its Present Simple form (now)
		%g - transforms a verb into gerund
		%p - transform a noun or a verb (Present Simple) into its plural form
		%l - transform a word to lower case
		%t - transform a word to Title Case
		%u - transform a word to UPPER CASE

Error is returned if:
  - provided pattern is empty
  - character other than the above is escaped with a '%' sign
  - a single '%' ends the pattern

Error is not returned if:
  - incompatible modifier is assigned to the word
  - duplicate modifier is assigned to the same word

Example phrase:

	"%tn %2v a %ua %un" may produce "Serenade perplexed a STRAY SUPERBUG"
*/
func (gen *Generator) Phrase(pattern string) (string, error) {
	if len(pattern) == 0 {
		return "", errEmptyPattern
	}

	var (
		// If true, the next character is interpreted as syntax character
		escaped bool

		// Container for modifiers for the current word
		mods []Mod

		// Built phrase
		phrase strings.Builder
	)

	for i, c := range pattern {
		if escaped {
			switch c {
			case '%':
				phrase.WriteRune(c)
				escaped = false
			case '2', '3', 'N', 'g', 'l', 'p', 't', 'u':
				mods = append(mods, flagToMod(c))
			case 'a', 'm', 'n', 'v':
				word, err := gen.getGenerator(c)(mods...)
				if err != nil {
					return "", err
				}
				phrase.WriteString(word)
				escaped = false
			default:
				return "", errUnknownCommand
			}
		} else if c == '%' {
			if i == len(pattern)-1 {
				return "", errEscapedStrTerm
			}

			escaped = true
			mods = make([]Mod, 0)
		} else {
			phrase.WriteRune(c)
		}
	}

	return phrase.String(), nil
}

/*
Generates a single random verb and transforms it according to mods.
Returns an error if an undefined Mod is received.
*/
func (gen *Generator) Verb(mods ...Mod) (string, error) {
	return gen.Transform(randItem(gen.verbs), mods...)
}

/*
A helper method that was created to make the loop in Generator.Phrase easier to understand.
It accepts an insertion command character and returns the corresponding generator method.
nil is never returned as this method is only called when a valid insertion command
is encountered.
*/
func (gen *Generator) getGenerator(flag rune) func(...Mod) (string, error) {
	switch flag {
	case 'a':
		return gen.Adjective
	case 'm':
		return gen.Adverb
	case 'n':
		return gen.Noun
	case 'v':
		return gen.Verb
	default:
		return nil
	}
}

/* Returns a new Generator with default word lists. */
func DefaultGenerator() (*Generator, error) {
	a, err := loadWords("res/adj")
	if err != nil {
		return nil, err
	}

	m, err := loadWords("res/adv")
	if err != nil {
		return nil, err
	}

	n, err := loadWords("res/noun")
	if err != nil {
		return nil, err
	}

	v, err := loadWords("res/verb")
	if err != nil {
		return nil, err
	}

	return NewGenerator(a, m, n, v)
}

/* Initializes a new Generator with provided lists. Returns error if any of the lists is empty. */
func NewGenerator(adj, adv, noun, verb []string) (*Generator, error) {
	if len(adj) == 0 || len(noun) == 0 || len(verb) == 0 {
		return nil, errEmptyLists
	}

	in, err := loadIrregularWords("res/noun.irr")
	if err != nil {
		return nil, err
	}

	iv, err := loadIrregularWords("res/verb.irr")
	if err != nil {
		return nil, err
	}

	return &Generator{
		adjectives: adj,
		adverbs:    adv,
		nouns:      noun,
		verbs:      verb,
		nounsIrr:   in,
		verbsIrr:   iv,
		caser:      newCaser(),
	}, nil
}
