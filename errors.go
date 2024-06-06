package neng

import "errors"

var (
	// Error returned by NewGenerator if nonpositive iterLimit is specified
	errBadIterLimit error = errors.New("iteration limit equal or lower than 0")

	// Error returned by NewGenerator if one or more of the user-provided lists are nil
	errEmptyLists error = errors.New("one or more of the provided word lists are nil")

	// Error returned by Generator.Phrase, if pattern is empty
	errEmptyPattern error = errors.New("provided pattern is empty")

	// Error returned by Generator.Phrase, if pattern ends with '%'
	errEscapedStrTerm error = errors.New("escape character at pattern termination")

	// Error returned by Generator.Adjective or Generator.Adverb if iteration limit is reached while searching for comparable word
	errIterLimit error = errors.New("iteration limit reached while trying to draw a valid comparative adjective or adverb")

	// Error returned by Generator.Transform, if non-comparable adjective or adverb is received along with gradation modifier
	errNonComparable error = errors.New("gradation requested, but the provided word is non-comparable")

	// Error returned by Generator.Transform, if uncountable noun is received along with pluralization modifier
	errUncountable error = errors.New("pluralization requested, but the provided word is uncountable")

	// Error returned by Generator.Transform, if undefined modifier value is received, e.g. Mod(123)
	errUndefinedMod error = errors.New("undefined modifier")

	// Error returned by Generator.Phrase, if pattern contains undefined escaped character
	errUnknownCommand error = errors.New("unknown command specified")
)
