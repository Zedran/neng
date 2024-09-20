package neng

import "errors"

var (
	// errBadIterLimit is returned by NewGenerator if a non-positive iterLimit
	// is specified.
	errBadIterLimit error = errors.New("iteration limit equal or lower than 0")

	// errBadWordList is returned by NewWord if any line in the word list is
	// incorrectly formatted or by NewGeneratorFromWord if any of the provided
	// slices contains a nil pointer.
	errBadWordList error = errors.New("word list contains invalid element(s)")

	// errEmptyLists is returned by NewGenerator and NewGeneratorFromWord
	// if any of the user-provided lists is empty or nil.
	errEmptyLists error = errors.New("empty list provided")

	// errEmptyPattern is returned by Generator.Phrase if pattern is empty.
	errEmptyPattern error = errors.New("provided pattern is empty")

	// errEmptyWord is returned from NewWordFromParams if the 'word' parameter
	// is an empty string.
	errEmptyWord error = errors.New("provided word is an empty string")

	// errEscapedStrTerm is returned by Generator.Phrase
	// if pattern ends with '%'.
	errEscapedStrTerm error = errors.New("escape character at pattern termination")

	// errIncompatible is returned by Generator.TransformWord, if given
	// WordClass is incompatible with requested transformations.
	errIncompatible error = errors.New("WordClass not compatible with the provided Mod(s)")

	// errIterLimit is returned by Generator.Adjective, Generator.Adverb
	// or Generator.Noun if iteration limit is reached while trying to
	// generate a valid word to perform the requested gradation
	// or pluralization.
	errIterLimit error = errors.New("iteration limit reached while trying to draw a comparable or countable word")

	// errMalformedIrr is returned from NewWordFromParams, if ft == FT_IRREGULAR
	// and irr has incorrect length or any of its elements is an empty string.
	errMalformedIrr error = errors.New("irregular forms slice is empty, too long, or contains an empty string")

	// errNonComparable is returned by Generator.TransformWord,
	// if non-comparable adjective or adverb is received along
	// with gradation modifier.
	errNonComparable error = errors.New("gradation requested, but the provided word is non-comparable")

	// errNonIrregular is returned if non-nil slice is passed as irr parameter
	// to NewWordsFromPar, but ft != FT_IRREGULAR.
	errNonIrregular error = errors.New("attempt to assign irregular forms, but Word is not irregular")

	// errNotFound is returned by Generator.Find if the specified word
	// is not found in the word database.
	errNotFound error = errors.New("no matches found")

	// errSpecStrTerm is returned by Generator.Phrase if a pattern ends
	// with transformation specifier (e.g "%t2").
	errSpecStrTerm error = errors.New("transformation specifier ends the pattern")

	// errUncountable is returned by Generator.TransformWord if an uncountable
	// noun is received along with MOD_PLURAL.
	errUncountable error = errors.New("pluralization requested, but the provided word is uncountable")

	// errUndefinedFormType is returned from NewWordFromParams if an undefined
	// FormType is passed as ft parameter, e.g. FormType(123).
	errUndefinedFormType error = errors.New("undefined FormType")

	// errUndefinedMod is returned by Generator.TransformWord if an undefined
	// modifier value is received, e.g. Mod(65536).
	errUndefinedMod error = errors.New("undefined modifier")

	// errUndefinedWordClass is returned by Generator.Find if an undefined
	// WordClass value is received, e.g. WordClass(123).
	errUndefinedWordClass error = errors.New("undefined WordClass")

	// errUnknownCommand is returned by Generator.Phrase if a pattern contains
	// an undefined escaped character.
	errUnknownCommand error = errors.New("unknown command specified")
)
