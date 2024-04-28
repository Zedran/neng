package neng

import "strings"

/* Generates random phrases or words. */
type Generator struct {
	adjectives []string
	nouns      []string
	verbs      []string
	verbsIrr   [][]string
	caser      *caser
}

/*
Generates a single random adjective and transforms it according to mods.
Returns error if undefined Mod value is passed, e.g. (Mod(123), where
maximum defined Mod value is 3).
*/
func (gen *Generator) Adjective(mods ...Mod) (string, error) {
	return gen.Transform(randItem(gen.adjectives), mods...)
}

/*
Transforms a word according to specified mods. Not all mods are compatible with every
part of speech. Compatibility is not checked. Returns error if undefined Mod value
is passed, e.g. (Mod(123), where maximum defined Mod value is 3).
*/
func (gen *Generator) Transform(word string, mods ...Mod) (string, error) {
	var caseTransformation func(string) string

	for _, m := range mods {
		switch m {
		case MOD_GERUND:
			word = gerund(word)
		case MOD_PRESENT_SIMPLE:
			word = presentSimple(word)
		case MOD_PAST_SIMPLE:
			word = pastSimple(word, gen.verbsIrr)
		case MOD_PAST_PARTICIPLE:
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

	if caseTransformation != nil {
		word = caseTransformation(word)
	}

	return word, nil
}

/*
Generates a single random noun and transforms it according to mods.
Returns error if undefined Mod value is passed, e.g. (Mod(123), where
maximum defined Mod value is 3).
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
		%n - inserts a random noun
		%v - inserts a random verb

	Transformation:
		%2 - transforms a verb into its Past Simple form (2nd form)
		%3 - transforms a verb into its Past Participle form (3rd form)
		%N - transforms a verb into its Present Simple form (now)
		%g - transforms a verb into gerund
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
		err error

		// If true, the next character is interpreted as syntax character
		escaped bool

		// Container for modifiers for the current word
		mods []Mod

		// Built phrase
		phrase strings.Builder

		// Out of scope container for the generated word
		word string
	)

	for i, c := range pattern {
		if escaped {
			switch c {
			case '%':
				phrase.WriteRune(c)
				escaped = false
				continue
			case '2':
				mods = append(mods, MOD_PAST_SIMPLE)
				continue
			case '3':
				mods = append(mods, MOD_PAST_PARTICIPLE)
				continue
			case 'N':
				mods = append(mods, MOD_PRESENT_SIMPLE)
				continue
			case 'a':
				word, err = gen.Adjective(mods...)
			case 'g':
				mods = append(mods, MOD_GERUND)
				continue
			case 'l':
				mods = append(mods, MOD_CASE_LOWER)
				continue
			case 'n':
				word, err = gen.Noun(mods...)
			case 't':
				mods = append(mods, MOD_CASE_TITLE)
				continue
			case 'u':
				mods = append(mods, MOD_CASE_UPPER)
				continue
			case 'v':
				word, err = gen.Verb(mods...)
			default:
				return "", errUnknownCommand
			}

			if err != nil {
				return "", err
			}

			phrase.WriteString(word)

			escaped = false
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
Returns error if undefined Mod value is passed, e.g. (Mod(123), where
maximum defined Mod value is 3).
*/
func (gen *Generator) Verb(mods ...Mod) (string, error) {
	return gen.Transform(randItem(gen.verbs), mods...)
}

/* Returns a new Generator with default word lists. */
func DefaultGenerator() (*Generator, error) {
	a, err := loadWords("res/adj")
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

	return NewGenerator(a, n, v)
}

/* Initializes a new Generator with provided lists. Returns error if any of the lists is empty. */
func NewGenerator(adj, noun, verb []string) (*Generator, error) {
	if len(adj) == 0 || len(noun) == 0 || len(verb) == 0 {
		return nil, errEmptyLists
	}

	iv, err := loadIrregularVerbs("res/verb.irr")
	if err != nil {
		return nil, err
	}

	return &Generator{
		adjectives: adj,
		nouns:      noun,
		verbs:      verb,
		verbsIrr:   iv,
		caser:      newCaser(),
	}, nil
}
