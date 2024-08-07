package neng

import "bytes"

// Indicates a specific transformation constraint of a word
type wordType uint8

const (
	// A regular word
	// adj, adv - forms comparative and superlative by adding 'more' to it
	// noun     - can be both singular and plural
	// verb     - a regular verb
	wt_regular wordType = iota

	// An irregular word, has its own special forms for:
	// adj, adv - comparative, superlative
	// noun     - plural
	// verb     - Past Simple, Past Participle
	wt_irregular

	// A plural-only noun (e.g. scissors)
	wt_plural_only

	// Adjective or adverb graded by appending '-er' and '-est' suffixes
	wt_suffixed

	// Uncomparable adjective or adverb
	wt_uncomparable

	// Uncountable noun
	wt_uncountable
)

/* Constitutes a single word list entry. */
type word struct {
	// A slice of irregular forms or nil
	irr *[]string

	// Word type
	t wordType

	// Word from the list
	word string
}

/* Parses a single word list line into a new word struct. Returns an error if malformed line is encountered. */
func NewWord(line []byte) (*word, error) {
	var w word

	// Split a line into separate, space separated fields
	s := bytes.Split(line, []byte(" "))

	if len(s) < 2 || len(s[1]) != 1 {
		// If the line contains less than 2 fields
		// or the second field is not one byte long,
		// the input is malformed
		return nil, errBadWordList
	}

	// Assign the word to the struct field
	w.word = string(s[0])

	// Retrieve WordType by subtracting ASCII zero
	// from the second field
	w.t = wordType(s[1][0] - 48)

	if w.t != wt_irregular {
		// Regular words do not require further processing
		w.irr = nil
		return &w, nil
	}

	if len(s) < 3 {
		// Irregular words require the third field to be present
		return nil, errBadWordList
	}

	// Split multi-word entry (replace underscore with space - "a_b" -> "a b")
	multiWord := bytes.Replace(s[2], []byte("_"), []byte(" "), -1)

	// Split the third, comma-separated field
	irr := bytes.Split(multiWord, []byte(","))

	if len(irr) < 1 {
		// At least one irregular form is required
		return nil, errBadWordList
	}

	slice := make([]string, len(irr))

	for i := range slice {
		if len(irr[i]) == 0 {
			return nil, errBadWordList
		}

		slice[i] = string(irr[i])
	}

	w.irr = &slice

	return &w, nil
}
