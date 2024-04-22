package neng

import "errors"

var (
	// Error returned from Generator.Phrase, if pattern is empty
	errEmptyPattern   error = errors.New("provided pattern is empty")

	// Error returned from Generator.Phrase, if pattern ends with '%'
	errEscapedStrTerm error = errors.New("escape character at pattern termination")

	// Error returned from Generator.Transform, if unknown modifier value is received (e.g. Mod(123) is passed)
	errUndefinedMod   error = errors.New("undefined modifier")

	// Error returned from Generator.Phrase, if pattern contains undefined escaped character
	errUnknownCommand error = errors.New("unknown command specified")
)
