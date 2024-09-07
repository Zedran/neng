package neng

// WordClass type helps the Generator to differentiate parts of speech.
type WordClass uint8

const (
	WC_ADJECTIVE WordClass = iota
	WC_ADVERB
	WC_NOUN
	WC_VERB
)

/* Returns true if wc is compatible with all of the received mods. */
func (wc WordClass) CompatibleWith(mods Mod) bool {
	switch wc {
	case WC_ADJECTIVE, WC_ADVERB:
		if mods.Enabled(MOD_PLURAL | MOD_PAST_SIMPLE | MOD_PAST_PARTICIPLE | MOD_PRESENT_SIMPLE | MOD_GERUND) {
			return false
		}
	case WC_NOUN:
		if mods.Enabled(MOD_PAST_SIMPLE | MOD_PAST_PARTICIPLE | MOD_PRESENT_SIMPLE | MOD_GERUND | MOD_COMPARATIVE | MOD_SUPERLATIVE) {
			return false
		}
	case WC_VERB:
		if mods.Enabled(MOD_COMPARATIVE | MOD_SUPERLATIVE) {
			return false
		}

		if mods.Enabled(MOD_PLURAL) && !mods.Enabled(MOD_PAST_SIMPLE|MOD_PRESENT_SIMPLE) {
			return false
		}
	}

	return true
}
