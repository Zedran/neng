package neng

// Form type or formation type. Indicates the effect that grammatical transformations have on a given word.
type FormType uint8

const (
	// A regular word
	// adj, adv - forms comparative and superlative by inserting 'more' and 'most' before it
	// noun     - can be both singular and plural
	// verb     - a regular verb
	FT_REGULAR FormType = iota

	// An irregular word, has its own special forms for:
	// adj, adv - comparative, superlative
	// noun     - plural
	// verb     - Past Simple, Past Participle
	FT_IRREGULAR

	// A plural-only noun (e.g. scissors), does not get picked in absence of MOD_PLURAL.
	FT_PLURAL_ONLY

	// Adjective or adverb graded by appending '-er' and '-est' suffixes.
	FT_SUFFIXED

	// Uncomparable adjective or adverb, does not get picked
	// if MOD_COMPARATIVE or MOD_SUPERLATIVE is requested.
	// An attempt to grade uncomparable word results
	// in an error.
	FT_UNCOMPARABLE

	// Uncountable noun, does not get picked if MOD_PLURAL was requested.
	// An attempt to pluralize uncountable noun results in an error.
	FT_UNCOUNTABLE
)
