package neng

// Indicates a specific transformation constraint of a word
type WordType uint8

const (
	// A regular word
	// adj, adv - forms comparative and superlative by adding 'more' to it
	// noun     - can be both singular and plural
	// verb     - a regular verb
	WT_REGULAR WordType = iota

	// An irregular word, has its own special forms for:
	// adj, adv - comparative, superlative
	// noun     - plural
	// verb     - Past Simple, Past Participle
	WT_IRREGULAR

	// A plural-only noun (e.g. scissors)
	WT_PLURAL_ONLY

	// Adjective or adverb graded by appending '-er' and '-est' suffixes
	WT_SUFFIXED

	// Uncomparable adjective or adverb
	WT_UNCOMPARABLE

	// Uncountable noun
	WT_UNCOUNTABLE
)
