// neng -- Non-Extravagant Name Generator
// Copyright (C) 2024  Wojciech Głąb (github.com/Zedran)
//
// This file is part of neng.
//
// neng is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3 only.
//
// neng is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with neng.  If not, see <https://www.gnu.org/licenses/>.

package neng

import (
	"strings"

	"github.com/Zedran/neng/symbols"
)

// Word represents a single word list entry.
type Word struct {
	// Form type
	ft FormType

	// A slice of irregular forms or nil
	irr *[]string

	// Word from the list
	word string
}

// Returns FormType of the Word.
func (w *Word) FT() FormType {
	return w.ft
}

// Returns an irregular form of the Word at index i of the underlying slice
// of irregular forms. Returns an error if called for a non-irregular word
// or if i is out of bounds of the slice.
func (w *Word) Irr(i int) (string, error) {
	if w.ft != FT_IRREGULAR {
		return "", symbols.ErrNonIrregular
	}

	if i < 0 || i >= len(*w.irr) {
		return "", symbols.ErrOutOfBounds
	}

	return (*w.irr)[i], nil
}

// Word returns base form of the word.
func (w *Word) Word() string {
	return w.word
}

// NewWord parses a single word list line into a new word struct.
// Returns an error if malformed line is encountered.
func NewWord(line string) (*Word, error) {
	if len(line) < 2 {
		// Line must contain at least two characters:
		//   - a single digit denoting a type
		//   - a word at least one character in length
		return nil, symbols.ErrBadWordList
	}

	w := Word{
		// Read the FormType by subtracting ASCII zero from the first byte
		ft: FormType(line[0] - 48),
	}

	if w.ft < FT_REGULAR || w.ft > FT_UNCOUNTABLE {
		// FormType must be within range of the defined values.
		// Coincidentally, this expression returns an error
		// if a comma begins the line.
		return nil, symbols.ErrBadWordList
	}

	// Find the first comma
	c1 := strings.IndexByte(line, ',')

	if w.ft != FT_IRREGULAR {
		if c1 != -1 {
			// Only irregular words can have irregular forms
			return nil, symbols.ErrBadWordList
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
		return nil, symbols.ErrBadWordList
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
		// The condition above ensures it is not zero-length at this point
		// (comma does not end the line)
		w.irr = &[]string{line[c1:]}

		return &w, nil
	}

	c2 += c1

	if c2 == len(line)-1 || strings.IndexByte(line[c2+1:], ',') != -1 {
		// Comma at the end of the line means the second word is zero-length
		// or - more commas mean more irregular forms - two at most are allowed
		return nil, symbols.ErrBadWordList
	}

	// Assign both irregular forms to a word
	w.irr = &[]string{line[c1:c2], line[c2+1:]}

	return &w, nil
}

// NewWordFromParams returns a Word struct built from the specified parameters,
// or error, if the following conditions are not met:
//
//   - word must be at least 1 character long
//   - ft must be in range of the defined FormType values
//   - for irregular words, the length of irr must be 1 or 2
//     and the elements cannot be empty strings
//   - for non-irregular words, irr must be empty
func NewWordFromParams(word string, ft FormType, irr []string) (*Word, error) {
	if len(word) == 0 {
		return nil, symbols.ErrEmptyWord
	}

	if ft < FT_REGULAR || ft > FT_UNCOUNTABLE {
		return nil, symbols.ErrUndefinedFormType
	}

	var pIrr *[]string

	if ft == FT_IRREGULAR {
		if len(irr) != 1 && len(irr) != 2 {
			return nil, symbols.ErrMalformedIrr
		}

		for _, e := range irr {
			if len(e) == 0 {
				return nil, symbols.ErrMalformedIrr
			}
		}

		pIrr = &irr
	} else if len(irr) > 0 {
		return nil, symbols.ErrNonIrregular
	}

	return &Word{ft: ft, irr: pIrr, word: word}, nil
}
