package neng

import "strings"

/* Constitutes a single word list entry. */
type Word struct {
	// A slice of irregular forms or nil
	irr *[]string

	// Form type
	ft FormType

	// Word from the list
	word string
}

/* Returns the unmodified word. */
func (w *Word) Word() string {
	return w.word
}

/* Parses a single word list line into a new word struct. Returns an error if malformed line is encountered. */
func NewWord(line string) (*Word, error) {
	if len(line) < 2 {
		// Line must contain at least two characters:
		// - a single digit denoting a type
		// - a word at least one character in length
		return nil, errBadWordList
	}

	w := Word{
		// Read the FormType by subtracting ASCII zero from the first byte
		ft: FormType(line[0] - 48),
	}

	if w.ft < FT_REGULAR || w.ft > FT_UNCOUNTABLE {
		// Type must be within range of defined wordType values.
		// Coincidentally, this expression returns an error
		// if a comma begins the line.
		return nil, errBadWordList
	}

	// Find the first comma
	c1 := strings.IndexByte(line, ',')

	if w.ft != FT_IRREGULAR {
		if c1 != -1 {
			// Only irregular words can have irregular forms
			return nil, errBadWordList
		}
		// Assign value to word field - from index 1 to the end of the line
		w.word = line[1:]
		return &w, nil
	}

	if c1 <= 1 || c1 == len(line)-1 {
		// -1 - No comma, there must be at least one in irregular word's line
		//  0 - Checked above already, comma cannot begin the line
		//  1 - If the first comma is found at index 1, the word
		//      has the length of zero
		// or - Comma cannot end the line
		return nil, errBadWordList
	}

	// Assign value to word field - from index 1 to the first comma
	w.word = line[1:c1]

	c1++
	// Find a second comma.
	// Substring operation - counts from the first comma,
	// not from the beginning of a line
	c2 := strings.IndexByte(line[c1:], ',')

	if c2 == -1 {
		// If there is no second comma, the word has only one irregular form
		// The condition above ensures it is not zero-length by this point
		// (comma does not end the line)
		w.irr = &[]string{line[c1:]}

		return &w, nil
	}

	c2 += c1

	if c2 == len(line)-1 || strings.IndexByte(line[c2+1:], ',') != -1 {
		// Comma at the end of the line means the second word is zero-length
		// or - more commas mean more irregular forms - two at most are allowed
		return nil, errBadWordList
	}

	// Assign both irregular forms to a word
	w.irr = &[]string{line[c1:c2], line[c2+1:]}

	return &w, nil
}

/*
Returns Word struct built from the specified parameters, or error, if the following conditions are not met:
  - word must be at least 1 character long
  - ft must be in range of the defined values
  - for irregular words, the length of irr must be 1 or 2 and the elements cannot be empty strings
  - for non-irregular words, irr must be empty
*/
func NewWordFromParams(word string, ft FormType, irr []string) (*Word, error) {
	if len(word) == 0 {
		return nil, errEmptyWord
	}

	if ft < FT_REGULAR || ft > FT_UNCOUNTABLE {
		return nil, errUndefinedFormType
	}

	var pIrr *[]string

	if ft == FT_IRREGULAR {
		if len(irr) != 1 && len(irr) != 2 {
			return nil, errMalformedIrr
		}

		for _, e := range irr {
			if len(e) == 0 {
				return nil, errMalformedIrr
			}
		}

		pIrr = &irr
	} else if len(irr) > 0 {
		return nil, errNonIrregular
	}

	return &Word{irr: pIrr, ft: ft, word: word}, nil
}
