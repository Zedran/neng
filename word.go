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
	if len(line) < 2 {
		// Line must contain at least two characters:
		// - a single digit denoting a type
		// - a word at least one character in length
		return nil, errBadWordList
	}

	w := word{
		// Read the wordType by subtracting ASCII zero from the first byte
		t: wordType(line[0] - 48),
	}

	if w.t < wt_regular || w.t > wt_uncountable {
		// Type must be within range of defined wordType values.
		// Coincidentally, this expression returns an error
		// if a comma begins the line.
		return nil, errBadWordList
	}

	if w.t != wt_irregular {
		// Words other than irregular require no further processing
		// Assign value to word field - from index 1 to the end of the line
		w.word = string(line[1:])
		return &w, nil
	}

	// Find the first comma
	c1 := bytes.IndexByte(line, ',')

	if c1 <= 1 || c1 == len(line)-1 {
		// -1 - No comma, there must be at least one in irregular word's line
		//  0 - Checked above already, comma cannot begin the line
		//  1 - If the first comma is found at index 1, the word
		//      has the length of zero
		// or - Comma cannot end the line
		return nil, errBadWordList
	}

	// Assign value to word field - from index 1 to the first comma
	w.word = string(line[1:c1])

	c1++
	// Find a second comma.
	// Substring operation - counts from the first comma,
	// not from the beginning of a line
	c2 := bytes.IndexByte(line[c1:], ',')

	if c2 == -1 {
		// If there is no second comma, the word has only one irregular form
		// The condition above ensures it is not zero-length by this point
		// (comma does not end the line)
		w.irr = &[]string{string(line[c1:])}

		return &w, nil
	}

	c2 += c1

	if c2 == len(line)-1 {
		// Comma at the end of the line means the second word is zero-length
		return nil, errBadWordList
	}

	// Assign both irregular forms to a word
	w.irr = &[]string{string(line[c1:c2]), string(line[c2+1:])}

	return &w, nil
}
