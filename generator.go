package neng

import "strings"

/* Generates random phrases or words. */
type Generator struct {
	adjectives []string
	nouns      []string
	verbs      []string

	verbsIrr   [][]string
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
   Transform a word according to specified mods. Not all mods are compatible with every 
   part of speech. Compatibility is not checked. Returns error if undefined Mod value 
   is passed, e.g. (Mod(123), where maximum defined Mod value is 3).
*/
func (gen *Generator) Transform(word string, mods ...Mod) (string, error) {
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
		default:
			return "", errUndefinedMod
		}
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

Error is returned if:
	* provided pattern is empty
	* character other than the above is escaped with a '%' sign
	* a single '%' ends the pattern

Error is not returned if:
	* incompatible modifier is assigned to the word
	* duplicate modifier is assigned to the same word

Example phrase:
	"a pretty %a %n" may produce "a pretty revocable snowfall"

Example transformation:
	"%n %2v a %a %n" may produce "serenade perplexed a stray superbug"
*/
func (gen *Generator) Phrase(pattern string) (string, error) {
	if len(pattern) == 0 {
		return "", errEmptyPattern
	}

	var (
		err     error

		// If true, the next character is interpreted as syntax character
		escaped bool
		
		// Container for modifiers for the current word
		mods    []Mod

		// Built phrase
		phrase  strings.Builder

		// Out of scope container for the generated word
		word    string
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
			case 'n':
				word, err = gen.Noun(mods...)
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
			if i == len(pattern) - 1 {
				return "", errEscapedStrTerm
			}

			escaped = true
			mods    = make([]Mod, 0)
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

/* Returns a new Generator. */
func NewGenerator() (*Generator, error) {
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

	iv, err := loadIrregularVerbs("res/verb.irr")
	if err != nil {
		return nil, err
	}

	return &Generator{
		adjectives:  a,
		nouns     :  n,
		verbs     :  v,
		verbsIrr  : iv,
	}, nil
}
