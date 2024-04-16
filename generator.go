package neng

import "strings"

/* Generates random phrases or words. */
type Generator struct {
	adjectives []string
	nouns      []string
}

/* Generates a single random adjective. */
func (gen *Generator) Adjective() string {
	return randItem(gen.adjectives)
}

/* Generates a single random noun. */
func (gen *Generator) Noun() string {
	return randItem(gen.nouns)
}

/* 
Generates a phrase given the pattern.

Syntax:
	%% - inserts '%' sign
    %a - inserts a random adjective
	%n - inserts a random noun

Error is returned if:
	* provided pattern is empty
	* any other character is escaped with a '%' sign
	* a single '%' ends the pattern

Example phrase:
    "a pretty %a %n" may produce "a pretty revocable snowfall"    
*/
func (gen *Generator) Phrase(pattern string) (string, error) {
	if len(pattern) == 0 {
		return "", errEmptyPattern
	}

	var (
		// If true, the next character is interpreted as syntax character
		escaped   bool

		phrase    strings.Builder
	)

	for i, c := range pattern {
		if escaped {
			switch c {
			case '%':
				phrase.WriteRune(c)
			case 'a':
				phrase.WriteString(randItem(gen.adjectives))
			case 'n':
				phrase.WriteString(randItem(gen.nouns))
			default:
				return "", errUnknownCommand
			}

			escaped = false
		} else if c == '%' {
			if i == len(pattern) - 1 {
				return "", errEscapedStrTerm
			}

			escaped = true
		} else {
			phrase.WriteRune(c)
		}
	}

	return phrase.String(), nil
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

	return &Generator{adjectives: a, nouns: n}, nil
}
