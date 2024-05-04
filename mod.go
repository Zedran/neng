package neng

// Modification parameter for a generated word
type Mod uint8

const (
	// Transform a noun or a verb (Present Simple) into its plural form
	MOD_PLURAL Mod = iota

	// Add Past Simple suffix to a verb or substitute its irregular form
	MOD_PAST_SIMPLE

	// Add Past Simple suffix to a regular verb or substitute Past Participle form to an irregular one
	MOD_PAST_PARTICIPLE

	// Add Present Simple suffix to a verb (-s, -es)
	MOD_PRESENT_SIMPLE

	// Create gerund form of a verb (-ing)
	MOD_GERUND

	// Transform a word to lower case
	MOD_CASE_LOWER

	// Transform a word to Title Case
	MOD_CASE_TITLE

	// Transform a word to UPPER CASE
	MOD_CASE_UPPER
)
