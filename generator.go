package neng

import (
	"fmt"
	"slices"
	"strings"
)

// Iteration limit used by the default Generator
const DEFAULT_ITER_LIMIT int = 1000

/* Generates random phrases or words. */
type Generator struct {
	// List of adjectives
	adj []*Word

	// List of adverbs
	adv []*Word

	// List of nouns
	noun []*Word

	// List of verbs
	verb []*Word

	// Case transformation handler
	caser *caser

	// A safeguard for Generator.generateModifier and Generator.Noun methods.
	// In presence of MOD_COMPARATIVE, MOD_SUPERLATIVE or MOD_PLURAL, these methods
	// attempt to generate a comparable adjective, adverb or countable noun until
	// one is found. iterLimit was implemented to ensure the looped generation
	// does not render the application unresponsive by becoming either too long
	// or infinite, depending on the underlying word database.
	iterLimit int
}

/*
Generates a single random adjective and transforms it according to mods.

Returns an error if:
- an undefined Mod is received (relays from Generator.Transform)
- an incompatible Mod is received (relays from Generator.Transform)
- iteration limit is reached while attempting to generate a comparable adjective
*/
func (gen *Generator) Adjective(mods ...Mod) (string, error) {
	return gen.generateModifier(WC_ADJECTIVE, mods...)
}

/*
Generates a single random adverb and transforms it according to mods.

Returns an error if:
- an undefined Mod is received (relays from Generator.Transform)
- an incompatible Mod is received (relays from Generator.Transform)
- iteration limit is reached while attempting to generate a comparable adverb
*/
func (gen *Generator) Adverb(mods ...Mod) (string, error) {
	return gen.generateModifier(WC_ADVERB, mods...)
}

/*
Searches the word list for the specified query. Returns an error if no word
is found or if an unknown WordClass is passed to the function.
*/
func (gen *Generator) Find(query string, wc WordClass) (*Word, error) {
	var list []*Word

	switch wc {
	case WC_ADJECTIVE:
		list = gen.adj
	case WC_ADVERB:
		list = gen.adv
	case WC_NOUN:
		list = gen.noun
	case WC_VERB:
		list = gen.verb
	default:
		return nil, errUndefinedWordClass
	}

	for _, w := range list {
		if w.word == query {
			return w, nil
		}
	}

	return nil, errNotFound
}

/*
Generates a single random noun and transforms it according to mods.

Returns an error if:
- an undefined Mod is received (relays from Generator.Transform)
- an incompatible Mod is received (relays from Generator.Transform)
- iteration limit is reached while attempting to generate a countable noun
*/
func (gen *Generator) Noun(mods ...Mod) (string, error) {
	var excluded FormType

	if slices.Contains(mods, MOD_PLURAL) {
		excluded = FT_UNCOUNTABLE
	} else {
		excluded = FT_PLURAL_ONLY
	}

	for range gen.iterLimit {
		if n := gen.noun[randIndex(len(gen.noun))]; n.ft != excluded {
			return gen.TransformWord(n, WC_NOUN, mods...)
		}
	}

	return "", errIterLimit
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
		%c - transforms an adjective or an adverb into comparative (better)
		%g - transforms a verb into gerund
		%p - transform a noun or a verb (Present Simple) into its plural form
		%s - transforms an adjective or an adverb into superlative (best)
		%l - transform a word to lower case
		%t - transform a word to Title Case
		%u - transform a word to UPPER CASE

Error is returned if:
  - provided pattern is empty
  - character other than the above is escaped with a '%' sign
  - a single '%' ends the pattern
  - incompatible modifier is assigned to the word

Error is not returned if:
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
			case '2', '3', 'N', 'c', 'g', 'l', 'p', 's', 't', 'u':
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
Searches for the specified word and, if found, calls Generator.TransformWord to transform it.

Assumes the following about the 'word' argument:
- Word is lower case (irrelevant if case transformation is requested)
- Adjectives and adverbs are in their positive forms
- Nouns are in their singular forms
- Verbs are in their infinitive forms

Returns an error if:
- word of the WordClass wc does not exist in the database
- specified WordClass value of the word is unknown

Relays an error from Generator.Transform if:
- WordClass of the word is not compatible with any Mod in mods
- transformation into comparative or superlative form is requested for non-comparable adjective or adverb
- transformation into plural form is requested for an uncountable noun
*/
func (gen *Generator) Transform(word string, wc WordClass, mods ...Mod) (string, error) {
	w, err := gen.Find(word, wc)
	if err != nil {
		return "", err
	}

	return gen.TransformWord(w, wc, mods...)
}

/*
Transforms a word according to specified mods. Not all mods are compatible with every WordClass.

Assumes the following about the 'word' field of the 'word' argument:
- Word is lower case (irrelevant if case transformation is requested)
- Adjectives and adverbs are in their positive forms
- Nouns are in their singular forms
- Verbs are in their infinitive forms

Returns an error if:
- WordClass of the word is not compatible with any Mod in mods
- transformation into comparative or superlative form is requested for non-comparable adjective or adverb
- transformation into plural form is requested for an uncountable noun
*/
func (gen *Generator) TransformWord(word *Word, wc WordClass, mods ...Mod) (string, error) {
	if !wc.CompatibleWith(mods...) {
		return "", errIncompatible
	}

	switch wc {
	case WC_ADJECTIVE, WC_ADVERB:
		if (slices.Contains(mods, MOD_COMPARATIVE) || slices.Contains(mods, MOD_SUPERLATIVE)) && word.ft == FT_UNCOMPARABLE {
			return "", errNonComparable
		}
	case WC_NOUN:
		if slices.Contains(mods, MOD_PLURAL) && word.ft == FT_UNCOUNTABLE {
			return "", errUncountable
		}
	}

	var (
		caseTransformation func(string) string
		pluralMod          bool
		w                  string
	)

	// Ensure MOD_PLURAL is processed first
	slices.Sort(mods)

	for _, m := range mods {
		switch m {
		case MOD_PLURAL:
			if wc == WC_NOUN {
				w = plural(word)
				continue
			}
			pluralMod = true
		case MOD_GERUND:
			w = gerund(word.word)
		case MOD_PRESENT_SIMPLE:
			w = presentSimple(word.word, pluralMod)
		case MOD_PAST_SIMPLE:
			w = pastSimple(word, pluralMod)
		case MOD_PAST_PARTICIPLE:
			w = pastParticiple(word)
		case MOD_COMPARATIVE:
			w = comparative(word)
		case MOD_SUPERLATIVE:
			w = superlative(word)
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

	if len(w) == 0 {
		// If no mods other than case transformation
		// are requested, w remains empty
		w = word.word
	}

	if caseTransformation != nil {
		w = caseTransformation(w)
	}

	return w, nil
}

/*
Generates a single random verb and transforms it according to mods.
Returns an error if an undefined Mod is received.
*/
func (gen *Generator) Verb(mods ...Mod) (string, error) {
	return gen.TransformWord(gen.verb[randIndex(len(gen.verb))], WC_VERB, mods...)
}

/*
A common method used to generate adjectives (noun modifiers) and adverbs (verb modifiers).

Returns an error if:
- an undefined Mod is received (relays from Generator.TransformWord)
- an incompatible Mod is received (relays from Generator.TransformWord)
- Generator.iterLimit is reached while attempting to generate a comparable adjective or adverb
*/
func (gen *Generator) generateModifier(wc WordClass, mods ...Mod) (string, error) {
	var items []*Word

	if wc == WC_ADJECTIVE {
		items = gen.adj
	} else {
		items = gen.adv
	}

	if !slices.Contains(mods, MOD_COMPARATIVE) && !slices.Contains(mods, MOD_SUPERLATIVE) {
		return gen.TransformWord(items[randIndex(len(items))], wc, mods...)
	}

	for range gen.iterLimit {
		if a := items[randIndex(len(items))]; a.ft != FT_UNCOMPARABLE {
			return gen.TransformWord(a, wc, mods...)
		}
	}

	return "", errIterLimit
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
	a, err := loadLines("embed/adj")
	if err != nil {
		return nil, err
	}

	m, err := loadLines("embed/adv")
	if err != nil {
		return nil, err
	}

	n, err := loadLines("embed/noun")
	if err != nil {
		return nil, err
	}

	v, err := loadLines("embed/verb")
	if err != nil {
		return nil, err
	}

	return NewGenerator(a, m, n, v, DEFAULT_ITER_LIMIT)
}

/*
Initializes a new Generator with provided lists. Returns an error if any of the lists is empty.

iterLimit is a safeguard for Generator.Adjective, Generator.Adverb and Generator.Noun methods.
In presence of MOD_COMPARATIVE, MOD_SUPERLATIVE or MOD_PLURAL, those methods generate a word
candidate until they find a comparable / countable one or until iteration limit is reached.

iterLimit value should be set with regards to the size of your word base
and the number of non-comparable adjectives, adverbs and uncountable nouns in it.

For example, to meet the default iterLimit of 1000, the Generator would need to draw
a non-comparable or uncountable word 1,000 times in a row. The embedded database contains
approximately 10,000 adjectives, of which 700 are non-comparable, and 24,000 nouns,
with 1,700 being uncountable. Given these numbers, it is unlikely that the iterLimit
will be reached.
*/
func NewGenerator(adj, adv, noun, verb []string, iterLimit int) (*Generator, error) {
	if iterLimit <= 0 {
		return nil, errBadIterLimit
	}

	if len(adj) == 0 || len(adv) == 0 || len(noun) == 0 || len(verb) == 0 {
		return nil, errEmptyLists
	}

	var (
		err error

		gen = Generator{
			caser:     newCaser(),
			iterLimit: iterLimit,
		}
	)

	gen.adj, err = parseLines(adj)
	if err != nil {
		return nil, fmt.Errorf("adj:%w", err)
	}

	gen.adv, err = parseLines(adv)
	if err != nil {
		return nil, fmt.Errorf("adv:%w", err)
	}

	gen.noun, err = parseLines(noun)
	if err != nil {
		return nil, fmt.Errorf("noun:%w", err)
	}

	gen.verb, err = parseLines(verb)
	if err != nil {
		return nil, fmt.Errorf("verb:%w", err)
	}

	return &gen, nil
}
