package neng

import (
	"slices"
	"strings"
)

// Iteration limit used by the default Generator
const DEFAULT_ITER_LIMIT int = 1000

/* Generates random phrases or words. */
type Generator struct {
	// Main list of adjectives
	adjectives []string

	// Main list of adverbs
	adverbs []string

	// Main list of nouns
	nouns []string

	// Main list of verbs
	verbs []string

	// Adjectives that are graded by adding suffix (-er, -est)
	adjSuf []string

	// Non-comparable adjectives
	adjNC []string

	// Uncountable nouns
	nounsUnc []string

	// Irregularly graded adjectives
	adjIrr [][]string

	// Nouns with irregular plural forms
	nounsIrr [][]string

	// Irregular verbs
	verbsIrr [][]string

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
Returns an error if an undefined Mod is received or if the iteration limit
is reached while attempting to generate a comparable adjective.
*/
func (gen *Generator) Adjective(mods ...Mod) (string, error) {
	return gen.generateModifier(gen.adjectives, mods...)
}

/*
Generates a single random adverb and transforms it according to mods.
Returns an error if an undefined Mod is received or if the iteration limit
is reached while attempting to generate a comparable adverb.
*/
func (gen *Generator) Adverb(mods ...Mod) (string, error) {
	return gen.generateModifier(gen.adverbs, mods...)
}

/*
Transforms a word according to specified mods. Not all mods are compatible with every
part of speech. Compatibility is not checked. Returns an error if an undefined Mod is received
or if user requested gradation or pluralization, but provided word does not have the requested form.
*/
func (gen *Generator) Transform(word string, mods ...Mod) (string, error) {
	if (contains(mods, MOD_COMPARATIVE) || contains(mods, MOD_SUPERLATIVE)) && contains(gen.adjNC, word) {
		return "", errNonComparable
	}

	if contains(mods, MOD_PLURAL) && contains(gen.nounsUnc, word) {
		return "", errNonComparable
	}

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
			word = pastSimple(word, gen.verbsIrr, pluralMod)
		case MOD_PAST_PARTICIPLE:
			verbMod = true
			word = pastParticiple(word, gen.verbsIrr)
		case MOD_COMPARATIVE:
			word = comparative(word, gen.adjIrr, gen.adjSuf)
		case MOD_SUPERLATIVE:
			word = superlative(word, gen.adjIrr, gen.adjSuf)
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
Generates a single random noun and transforms it according to mods. Returns error
if timeout (Generator.iterLimit) is reached while attempting to generate a countable noun
or if an undefined Mod is received.
*/
func (gen *Generator) Noun(mods ...Mod) (string, error) {
	n := randItem(gen.nouns)

	if contains(mods, MOD_PLURAL) {
		i := 0
		for contains(gen.nounsUnc, n) {
			if i == DEFAULT_ITER_LIMIT {
				return "", errIterLimit
			}

			n = randItem(gen.nouns)
			i++
		}
	}

	return gen.Transform(n, mods...)
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
Generates a single random verb and transforms it according to mods.
Returns an error if an undefined Mod is received.
*/
func (gen *Generator) Verb(mods ...Mod) (string, error) {
	return gen.Transform(randItem(gen.verbs), mods...)
}

/*
A common method used to generate adjectives (noun modifiers) and adverbs (verb modifiers).
Returns error if Generator.iterLimit is reached while attempting to generate a comparable
adjective or adverb. Relays errUndefinedMod from Generator.Transform.
*/
func (gen *Generator) generateModifier(items []string, mods ...Mod) (string, error) {
	a := randItem(items)

	if contains(mods, MOD_COMPARATIVE) || contains(mods, MOD_SUPERLATIVE) {
		i := 0
		for contains(gen.adjNC, a) {
			if i == DEFAULT_ITER_LIMIT {
				return "", errIterLimit
			}

			a = randItem(items)
			i++
		}
	}

	return gen.Transform(a, mods...)
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

	return NewGenerator(a, m, n, v, DEFAULT_ITER_LIMIT)
}

/*
Initializes a new Generator with provided lists. Returns an error if any of the lists is empty.

iterLimit is a safeguard for Generator.Adjective, Generator.Adverb and Generator.Noun methods.
In presence of MOD_COMPARATIVE, MOD_SUPERLATIVE or MOD_PLURAL, those methods generate a word
candidate until they find a comparable / countable one or until iteration limit is reached.

iterLimit value should be set with regards to the size of your word base
and the number of non-comparable adjectives, adverbs and countable nouns in it.
*/
func NewGenerator(adj, adv, noun, verb []string, iterLimit int) (*Generator, error) {
	if iterLimit <= 0 {
		return nil, errBadIterLimit
	}

	if len(adj) == 0 || len(noun) == 0 || len(verb) == 0 {
		return nil, errEmptyLists
	}

	ia, err := loadIrregularWords("res/adj.irr")
	if err != nil {
		return nil, err
	}

	sa, err := loadWords("res/adj.suf")
	if err != nil {
		return nil, err
	}

	nc, err := loadWords("res/adj.ncmp")
	if err != nil {
		return nil, err
	}

	un, err := loadWords("res/noun.unc")
	if err != nil {
		return nil, err
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
		adjIrr:     ia,
		adjSuf:     sa,
		adjNC:      nc,
		nounsUnc:   un,
		nounsIrr:   in,
		verbsIrr:   iv,
		caser:      newCaser(),
		iterLimit:  iterLimit,
	}, nil
}
