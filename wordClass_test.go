package neng

import "testing"

// Tests WordClass.CompatibleWith. Fails if compatibility assessment
// is not consistent with documentation.
func TestWordClass_CompatibleWith(t *testing.T) {
	type testCase struct {
		expected bool
		wc       WordClass
		mods     Mod
	}

	cases := []testCase{
		{true, WC_ADJECTIVE, MOD_COMPARATIVE | MOD_SUPERLATIVE | MOD_CASE_LOWER | MOD_CASE_TITLE | MOD_CASE_UPPER},
		{true, WC_ADVERB, MOD_COMPARATIVE | MOD_SUPERLATIVE | MOD_CASE_LOWER | MOD_CASE_TITLE | MOD_CASE_UPPER},
		{true, WC_NOUN, MOD_PLURAL | MOD_CASE_LOWER | MOD_CASE_TITLE | MOD_CASE_UPPER},
		{true, WC_VERB, MOD_PAST_SIMPLE | MOD_PAST_PARTICIPLE | MOD_PRESENT_SIMPLE | MOD_GERUND | MOD_PLURAL | MOD_CASE_LOWER | MOD_CASE_TITLE | MOD_CASE_UPPER},
		{true, WC_VERB, MOD_PAST_SIMPLE},
		{true, WC_VERB, MOD_PRESENT_SIMPLE},
		{false, WC_ADJECTIVE, MOD_GERUND},
		{false, WC_ADJECTIVE, MOD_PLURAL},
		{false, WC_ADVERB, MOD_PLURAL},
		{false, WC_ADVERB, MOD_PAST_SIMPLE},
		{false, WC_NOUN, MOD_COMPARATIVE},
		{false, WC_VERB, MOD_SUPERLATIVE},
		{false, WC_VERB, MOD_PLURAL},
	}

	for _, c := range cases {
		out := c.wc.CompatibleWith(c.mods)

		if out != c.expected {
			t.Errorf("Failed for '%v': expected %v, got %v", c, c.expected, out)
		}
	}
}
