package neng

import (
	"fmt"
	"slices"
	"strings"
)

// Iteration limit used by DefaultGenerator function. This value can be used as iterLimit parameter
// in non-default Generator constructors.
//
// Iteration limit a safeguard for Generator.Adjective, Generator.Adverb and Generator.Noun methods.
// In presence of MOD_COMPARATIVE, MOD_SUPERLATIVE or MOD_PLURAL, those methods generate a word
// candidate until they find a comparable / countable one or until iteration limit is reached.
//
// iterLimit value should be set with regards to the size of your word base
// and the number of non-comparable adjectives, adverbs and uncountable nouns in it.
//
// For example, to meet the default iterLimit of 1,000, the Generator would need to draw
// a non-comparable or uncountable word 1,000 times in a row. The embedded database contains
// approximately 10,000 adjectives, of which 2,700 are non-comparable, and 23,000 nouns,
// with 3,000 being uncountable. Given these numbers, it is unlikely that the iterLimit is reached.
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
func (gen *Generator) Adjective(mods Mod) (string, error) {
	return gen.generateModifier(WC_ADJECTIVE, mods)
}

/*
Generates a single random adverb and transforms it according to mods.

Returns an error if:
  - an undefined Mod is received (relays from Generator.Transform)
  - an incompatible Mod is received (relays from Generator.Transform)
  - iteration limit is reached while attempting to generate a comparable adverb
*/
func (gen *Generator) Adverb(mods Mod) (string, error) {
	return gen.generateModifier(WC_ADVERB, mods)
}

/*
Searches the word list for the specified word. Returns an error if no word
is found or if an unknown WordClass is passed to the function.

Assumes the following about the 'word' argument:
  - Word is lower case
  - Adjectives and adverbs are in their positive forms
  - Nouns are in their singular forms
  - Verbs are in their base forms
*/
func (gen *Generator) Find(word string, wc WordClass) (*Word, error) {
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

	n, found := slices.BinarySearchFunc(list, word, func(listItem *Word, word string) int {
		return strings.Compare((*listItem).word, word)
	})

	if found {
		return list[n], nil
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
func (gen *Generator) Noun(mods Mod) (string, error) {
	var excluded FormType

	if mods.Enabled(MOD_PLURAL) {
		excluded = FT_UNCOUNTABLE
	} else {
		excluded = FT_PLURAL_ONLY
	}

	for range gen.iterLimit {
		if n := gen.noun[randIndex(len(gen.noun))]; n.ft != excluded {
			return gen.TransformWord(n, WC_NOUN, mods)
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
		%p - transforms a noun or a verb (Present Simple) into its plural form
		%s - transforms an adjective or an adverb into superlative (best)
		%l - transforms a word to lower case
		%t - transforms a word to Title Case
		%u - transforms a word to UPPER CASE

Error is returned if:
  - provided pattern is empty
  - character other than the above is escaped with a '%' sign
  - a single '%' ends the pattern
  - transformation specifier ends the pattern ("%t2")
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

		// Collects Mod values for the current word
		mods Mod

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
				if i == len(pattern)-1 {
					return "", errSpecStrTerm
				}
				mods |= flagToMod(c)
			case 'a', 'm', 'n', 'v':
				word, err := gen.getGenerator(c)(mods)
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
			mods = MOD_NONE
		} else {
			phrase.WriteRune(c)
		}
	}

	return phrase.String(), nil
}

/*
Searches for the specified word and, if found, calls Generator.TransformWord to transform it.

Assumes the following about the 'word' argument:
  - Word is lower case
  - Adjectives and adverbs are in their positive forms
  - Nouns are in their singular forms
  - Verbs are in their base forms

Returns an error if:
  - word of the WordClass wc does not exist in the database
  - specified WordClass value of the word is unknown

Relays an error from Generator.TransformWord if:
  - WordClass of the word is not compatible with any Mod in mods
  - transformation into comparative or superlative form is requested for non-comparable adjective or adverb
  - transformation into plural form is requested for an uncountable noun
*/
func (gen *Generator) Transform(word string, wc WordClass, mods Mod) (string, error) {
	w, err := gen.Find(word, wc)
	if err != nil {
		return "", err
	}

	return gen.TransformWord(w, wc, mods)
}

/*
Transforms a word according to specified mods. Not all mods are compatible with every WordClass.

Assumes the following about the 'word' field of the 'word' argument:
  - Word.word is lower case
  - Adjectives and adverbs are in their positive forms
  - Nouns are in their singular forms
  - Verbs are in their base forms

Returns an error if:
  - WordClass of the word is not compatible with any Mod in mods
  - transformation into comparative or superlative form is requested for non-comparable adjective or adverb
  - transformation into plural form is requested for an uncountable noun
*/
func (gen *Generator) TransformWord(word *Word, wc WordClass, mods Mod) (string, error) {
	if mods == MOD_NONE {
		return word.word, nil
	}

	if mods.Undefined() {
		return "", errUndefinedMod
	}

	if !wc.CompatibleWith(mods) {
		return "", errIncompatible
	}

	switch wc {
	case WC_ADJECTIVE, WC_ADVERB:
		if word.ft == FT_NON_COMPARABLE && mods.Enabled(MOD_COMPARATIVE|MOD_SUPERLATIVE) {
			return "", errNonComparable
		}
	case WC_NOUN:
		if word.ft == FT_UNCOUNTABLE && mods.Enabled(MOD_PLURAL) {
			return "", errUncountable
		}
	}

	var w string

	switch wc {
	case WC_ADJECTIVE, WC_ADVERB:
		if mods.Enabled(MOD_COMPARATIVE) {
			w = comparative(word)
		} else if mods.Enabled(MOD_SUPERLATIVE) {
			w = superlative(word)
		}
	case WC_NOUN:
		if mods.Enabled(MOD_PLURAL) {
			w = plural(word)
		}
	case WC_VERB:
		if mods.Enabled(MOD_PAST_SIMPLE) {
			w = pastSimple(word, mods.Enabled(MOD_PLURAL))
		} else if mods.Enabled(MOD_PAST_PARTICIPLE) {
			w = pastParticiple(word)
		} else if mods.Enabled(MOD_PRESENT_SIMPLE) {
			w = presentSimple(word.word, mods.Enabled(MOD_PLURAL))
		} else if mods.Enabled(MOD_GERUND) {
			w = gerund(word.word)
		}
	}

	if len(w) == 0 {
		// If no mods other than case transformation
		// are requested, w remains empty
		w = word.word
	}

	switch true {
	case mods.Enabled(MOD_CASE_LOWER):
		w = gen.caser.toLower(w)
	case mods.Enabled(MOD_CASE_TITLE):
		w = gen.caser.toTitle(w)
	case mods.Enabled(MOD_CASE_UPPER):
		w = gen.caser.toUpper(w)
	}

	return w, nil
}

/*
Generates a single random verb and transforms it according to mods.
Returns an error if an undefined Mod is received.
*/
func (gen *Generator) Verb(mods Mod) (string, error) {
	return gen.TransformWord(gen.verb[randIndex(len(gen.verb))], WC_VERB, mods)
}

/*
A common method used to generate adjectives (noun modifiers) and adverbs (verb modifiers).

Returns an error if:
  - an undefined Mod is received (relays from Generator.TransformWord)
  - an incompatible Mod is received (relays from Generator.TransformWord)
  - Generator.iterLimit is reached while attempting to generate a comparable adjective or adverb
*/
func (gen *Generator) generateModifier(wc WordClass, mods Mod) (string, error) {
	var items []*Word

	if wc == WC_ADJECTIVE {
		items = gen.adj
	} else {
		items = gen.adv
	}

	if !mods.Enabled(MOD_COMPARATIVE | MOD_SUPERLATIVE) {
		return gen.TransformWord(items[randIndex(len(items))], wc, mods)
	}

	for range gen.iterLimit {
		if a := items[randIndex(len(items))]; a.ft != FT_NON_COMPARABLE {
			return gen.TransformWord(a, wc, mods)
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
func (gen *Generator) getGenerator(flag rune) func(Mod) (string, error) {
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

	return NewGenerator(a, m, n, v, DEFAULT_ITER_LIMIT, false)
}

/*
Initializes a new Generator with provided lists. Returns an error if any of the lists is empty
or if any of the elements is incorrectly formatted.

Line structure:

	<FormType><word>[,irr1][,irr2]

Base:
  - FormType         - a single digit
  - Word             - the word itself, at least one character long

If FormType == FT_IRREGULAR:
  - Irregular form 1 - separated from the word by a comma
  - Irregular form 2 - separated from the first irregular form by a comma

Like the word, irregular forms must be at least one character long. All words must be lower case.
Every slice must be sorted A-Z by word.

The safe parameter allows users to bypass word list checks.

If safe is false:
  - empty or nil slices do not trigger an error
  - slices are not sorted

Regardless of safe value:
  - malformed lines trigger an error
  - letter case is not checked

iterLimit is an adjustable safety mechanism to prevent inifinite loops during certain transformations.
For more information, refer to DEFAULT_ITER_LIMIT in the section 'Constants'.
*/
func NewGenerator(adj, adv, noun, verb []string, iterLimit int, safe bool) (*Generator, error) {
	wAdj, err := parseLines(adj)
	if err != nil {
		return nil, fmt.Errorf("adj:%w", err)
	}

	wAdv, err := parseLines(adv)
	if err != nil {
		return nil, fmt.Errorf("adv:%w", err)
	}

	wNoun, err := parseLines(noun)
	if err != nil {
		return nil, fmt.Errorf("noun:%w", err)
	}

	wVerb, err := parseLines(verb)
	if err != nil {
		return nil, fmt.Errorf("verb:%w", err)
	}

	return NewGeneratorFromWord(wAdj, wAdv, wNoun, wVerb, iterLimit, safe)
}

/*
Returns Generator created using the provided lists of Word structs and iterLimit. Returns an error
if any of the lists is empty or contains a nil pointer. If safe is false, empty / nil checks are omitted.
It is assumed that Word structs are created using one of the safe constructors, therefore their validity
is not verified. Those constructors do not check word case though - all words should be lower case.
Every slice must be sorted A-Z by Word.word field. If safe is true, the function ensures the correct order.
iterLimit is an adjustable safety mechanism to prevent inifinite loops during certain transformations.
For more information, refer to DEFAULT_ITER_LIMIT in the section 'Constants'.
*/
func NewGeneratorFromWord(adj, adv, noun, verb []*Word, iterLimit int, safe bool) (*Generator, error) {
	if iterLimit <= 0 {
		return nil, errBadIterLimit
	}

	if safe {
		for _, wordList := range [][]*Word{adj, adv, noun, verb} {
			if len(wordList) == 0 {
				return nil, errEmptyLists
			}

			if slices.Contains(wordList, nil) {
				return nil, errBadWordList
			}

			if !slices.IsSortedFunc(wordList, cmpWord) {
				slices.SortFunc(wordList, cmpWord)
			}
		}
	}

	gen := Generator{
		adj:       adj,
		adv:       adv,
		noun:      noun,
		verb:      verb,
		caser:     newCaser(),
		iterLimit: iterLimit,
	}

	return &gen, nil
}
