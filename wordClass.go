package neng

import "slices"

// WordClass type helps Generator.Transform differentiate
// parts of speech and process them accordingly.
type WordClass uint8

const (
	WC_ADJECTIVE WordClass = iota
	WC_ADVERB
	WC_NOUN
	WC_VERB
)

/* Returns true if wc is compatible with all of the received mods. */
func (wc WordClass) CompatibleWith(mods ...Mod) bool {
	for _, m := range mods {
		switch m {

		case MOD_PLURAL:
			if wc == WC_ADJECTIVE || wc == WC_ADVERB {
				return false
			}

			if wc == WC_VERB && !slices.Contains(mods, MOD_PAST_SIMPLE) && !slices.Contains(mods, MOD_PRESENT_SIMPLE) {
				return false
			}

		case MOD_PAST_SIMPLE, MOD_PAST_PARTICIPLE, MOD_PRESENT_SIMPLE, MOD_GERUND:
			if wc != WC_VERB {
				return false
			}

		case MOD_COMPARATIVE, MOD_SUPERLATIVE:
			if wc != WC_ADJECTIVE && wc != WC_ADVERB {
				return false
			}

		case MOD_CASE_LOWER, MOD_CASE_TITLE, MOD_CASE_UPPER:
			continue

		default: // Undefined Mod
			return false
		}
	}

	return true
}
