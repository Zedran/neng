package neng

// Indicates a specific transformation constraint of a word
type FormType uint8

const (
	// A regular word
	// adj, adv - forms comparative and superlative by adding 'more' to it
	// noun     - can be both singular and plural
	// verb     - a regular verb
	FT_REGULAR FormType = iota

	// An irregular word, has its own special forms for:
	// adj, adv - comparative, superlative
	// noun     - plural
	// verb     - Past Simple, Past Participle
	FT_IRREGULAR

	// A plural-only noun (e.g. scissors)
	FT_PLURAL_ONLY

	// Adjective or adverb graded by appending '-er' and '-est' suffixes
	FT_SUFFIXED

	// Uncomparable adjective or adverb
	FT_UNCOMPARABLE

	// Uncountable noun
	FT_UNCOUNTABLE
)
