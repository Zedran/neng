package neng

// Mod holds modification parameters for a generated word.
type Mod uint

// Do not transform the word in any way.
const MOD_NONE Mod = 0

const (
	// Transform a noun or a verb (Past Simple, Present Simple)
	// into its plural form.
	MOD_PLURAL Mod = 1 << iota

	// Transform a verb into its Past Simple form.
	MOD_PAST_SIMPLE

	// Transform a verb into its Past Participle form.
	MOD_PAST_PARTICIPLE

	// Add Present Simple suffix to a verb (-s, -es).
	MOD_PRESENT_SIMPLE

	// Create gerund form of a verb (-ing).
	MOD_GERUND

	// Transform an adjective or an adverb into comparative (good -> better).
	MOD_COMPARATIVE

	// Transform an adjective or an adverb into superlative (good -> best).
	MOD_SUPERLATIVE

	// Transform a word to lower case.
	MOD_CASE_LOWER

	// Transform a word to Title Case.
	MOD_CASE_TITLE

	// Transform a word to UPPER CASE.
	MOD_CASE_UPPER

	// Internal value, declared to mark the end of usable Mod values.
	mod_undefined
)

// Enabled returns true if any of the specified mods are enabled in m.
// Do not use this method to test for MOD_NONE. Use a simple comparison instead.
func (m Mod) Enabled(mods Mod) bool {
	return m&mods != 0
}

// Undefined returns true if m holds an undefined Mod value.
func (m Mod) Undefined() bool {
	return m >= mod_undefined
}

// flagToMod translates flag character into Mod value.
func flagToMod(flag rune) Mod {
	switch flag {
	case 'p':
		return MOD_PLURAL
	case '2':
		return MOD_PAST_SIMPLE
	case '3':
		return MOD_PAST_PARTICIPLE
	case 'N':
		return MOD_PRESENT_SIMPLE
	case 'c':
		return MOD_COMPARATIVE
	case 'g':
		return MOD_GERUND
	case 'l':
		return MOD_CASE_LOWER
	case 's':
		return MOD_SUPERLATIVE
	case 't':
		return MOD_CASE_TITLE
	case 'u':
		return MOD_CASE_UPPER
	default:
		return Mod(mod_undefined)
	}
}
