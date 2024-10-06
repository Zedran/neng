package symbols

import "errors"

var (
	// ErrBadIterLimit is returned by NewGenerator if a non-positive iterLimit
	// is specified.
	ErrBadIterLimit = errors.New("iteration limit equal or lower than 0")

	// ErrBadWordList is returned by NewWord if any line in the word list is
	// incorrectly formatted or by NewGeneratorFromWord if any of the provided
	// slices contains a nil pointer.
	ErrBadWordList = errors.New("word list contains invalid element(s)")

	// ErrEmptyLists is returned by NewGenerator and NewGeneratorFromWord
	// if any of the user-provided lists is empty or nil.
	ErrEmptyLists = errors.New("empty list provided")

	// ErrEmptyPattern is returned by Generator.Phrase if pattern is empty.
	ErrEmptyPattern = errors.New("provided pattern is empty")

	// ErrEmptyWord is returned from NewWordFromParams if the 'word' parameter
	// is an empty string.
	ErrEmptyWord = errors.New("provided word is an empty string")

	// ErrEscapedStrTerm is returned by Generator.Phrase
	// if pattern ends with '%'.
	ErrEscapedStrTerm = errors.New("escape character at pattern termination")

	// ErrIncompatible is returned by Generator.TransformWord, if given
	// WordClass is incompatible with requested transformations.
	ErrIncompatible = errors.New("WordClass not compatible with the provided Mod(s)")

	// ErrIterLimit is returned by Generator.Adjective, Generator.Adverb
	// or Generator.Noun if iteration limit is reached while trying to
	// generate a valid word to perform the requested gradation
	// or pluralization.
	ErrIterLimit = errors.New("iteration limit reached while trying to draw a comparable or countable word")

	// ErrMalformedIrr is returned from NewWordFromParams, if ft == FT_IRREGULAR
	// and irr has incorrect length or any of its elements is an empty string.
	ErrMalformedIrr = errors.New("irregular forms slice is empty, too long, or contains an empty string")

	// ErrNonComparable is returned by Generator.TransformWord,
	// if non-comparable adjective or adverb is received along
	// with gradation modifier.
	ErrNonComparable = errors.New("gradation requested, but the provided word is non-comparable")

	// ErrNonIrregular is returned if non-nil slice is passed as irr parameter
	// to NewWordsFromPar, but ft != FT_IRREGULAR.
	ErrNonIrregular = errors.New("attempt to assign irregular forms, but Word is not irregular")

	// ErrNotFound is returned by Generator.Find if the specified word
	// is not found in the word database.
	ErrNotFound = errors.New("no matches found")

	// ErrPluralOnly is returned by Generator.TransformWord if a plural-only
	// noun is received along with MOD_INDEF or MOD_INDEF_SILENT.
	ErrPluralOnly = errors.New("indefinite article requested for plural-only noun")

	// ErrSpecStrTerm is returned by Generator.Phrase if a pattern ends
	// with transformation specifier (e.g "%t2").
	ErrSpecStrTerm = errors.New("transformation specifier ends the pattern")

	// ErrUncountable is returned by Generator.TransformWord if an uncountable
	// noun is received along with MOD_INDEF, MOD_INDEF_SILENT or MOD_PLURAL.
	ErrUncountable = errors.New("indefinite article or pluralization requested for uncountable noun")

	// ErrUndefinedFormType is returned from NewWordFromParams if an undefined
	// FormType is passed as ft parameter, e.g. FormType(123).
	ErrUndefinedFormType = errors.New("undefined FormType")

	// ErrUndefinedMod is returned by Generator.TransformWord if an undefined
	// modifier value is received, e.g. Mod(65536).
	ErrUndefinedMod = errors.New("undefined modifier")

	// ErrUndefinedSpecifier is returned by Generator.Phrase if a pattern contains
	// an undefined escaped character.
	ErrUndefinedSpecifier = errors.New("undefined specifier")

	// ErrUndefinedWordClass is returned by Generator.Find if an undefined
	// WordClass value is received, e.g. WordClass(123).
	ErrUndefinedWordClass = errors.New("undefined WordClass")
)
