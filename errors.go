package neng

import "errors"

var (
	// Error returned by NewGenerator if one or more of the user-provided lists are nil
	errEmptyLists error = errors.New("one or more of the provided word lists are nil")

	// Error returned by Generator.Phrase, if pattern is empty
	errEmptyPattern error = errors.New("provided pattern is empty")

	// Error returned by Generator.Phrase, if pattern ends with '%'
	errEscapedStrTerm error = errors.New("escape character at pattern termination")

	// Error returned by Generator.Transform, if unknown modifier value is received (e.g. Mod(123) is passed)
	errUndefinedMod error = errors.New("undefined modifier")

	// Error returned by Generator.Phrase, if pattern contains undefined escaped character
	errUnknownCommand error = errors.New("unknown command specified")
)
