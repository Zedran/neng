package neng

import "errors"

var (
	// Error returned from Generator.Phrase, if pattern is empty
	errEmptyPattern   error = errors.New("provided pattern is empty")

	// Error returned from Generator.Phrase, if pattern ends with '%'
	errEscapedStrTerm error = errors.New("escape character at pattern termination")

	// Error returned from Generator.Phrase, if pattern contains escaped character that are not '%', 'a' or 'n'
	errUnknownCommand error = errors.New("unknown command specified")
)
