package neng

import "testing"

/* Tests NewGenerator function. Fails if providing an empty list or nil does not trigger an error. */
func TestNewGenerator(t *testing.T) {
	type testCase struct {
		adj, adv, noun, verb []string
		goodCase             bool
	}

	cases := []testCase{
		{[]string{"adj"}, []string{"adv"}, []string{"noun"}, []string{"verb"}, true},
		{[]string{"adj", "adj2"}, []string{"adv"}, []string{"noun"}, []string{"verb"}, true},
		{[]string{"adj"}, []string{"adv"}, []string{"noun"}, []string{}, false},
		{[]string{"adj"}, []string{"adv"}, nil, []string{"verb"}, false},
	}

	for _, c := range cases {
		_, err := NewGenerator(c.adj, c.adv, c.noun, c.verb, DEFAULT_ITER_LIMIT)

		switch c.goodCase {
		case true:
			if err != nil {
				t.Errorf("Failed for '%v': NewGenerator returned an error: %s", c, err.Error())
			}
		default:
			if err == nil {
				t.Errorf("Failed for '%v': NewGenerator did not return an error.", c)
			}
		}
	}
}

/* Tests whether Generator.Phrase correctly parses pattern syntax and generates phrases. */
func TestPhrase(t *testing.T) {
	gen, err := NewGenerator([]string{"big"}, []string{"nicely"}, []string{"snowfall"}, []string{"stash"}, DEFAULT_ITER_LIMIT)
	if err != nil {
		t.Fatalf("Failed: NewGenerator returned an error: %s", err.Error())
	}

	cases, err := loadTestMapStringString("TestPhrase.json")
	if err != nil {
		t.Fatalf("Failed loading test data: %s", err.Error())
	}

	for input, expected := range cases {
		output, err := gen.Phrase(input)
		if err != nil {
			t.Errorf("Failed for case '%s': error returned: '%s'", input, err.Error())
		} else if output != expected {
			t.Errorf("Failed for case '%s': got '%s', expected '%s'", input, output, expected)
		}
	}

	errCases := []string{
		"",     // pattern is empty
		"abc%", // escape character at pattern termination
		"%q",   // unknown command
	}

	for _, bc := range errCases {
		output, err := gen.Phrase(bc)
		if err == nil {
			t.Errorf("Failed for errCase '%s': no error returned. Output: '%s'", bc, output)
		}
	}
}

/*
Tests whether the Generator.Transform returns errNonComparable and errUncountable for appropriate WordClass.
errNonComparable should only be returned if gradation was requested for non-comparable adjective or adverb and
errUncountable should only be returned if pluralization was requested for an uncountable noun.
*/
func TestTransform(t *testing.T) {
	type testCase struct {
		description string
		word        string
		mods        []Mod
		wc          WordClass
		goodCase    bool
	}

	gen, err := DefaultGenerator()
	if err != nil {
		t.Fatalf("Failed: NewGenerator returned an error: %s", err.Error())
	}

	cases := []testCase{
		{"Uncountable noun + MOD_PLURAL", "aa", []Mod{MOD_PLURAL}, WC_NOUN, false},
		{"Non-comparable adj + MOD_COMPARATIVE", "own", []Mod{MOD_COMPARATIVE}, WC_ADJECTIVE, false},
		{"Non-comparable adj + MOD_SUPERLATIVE", "own", []Mod{MOD_SUPERLATIVE}, WC_ADJECTIVE, false},
		{"Noun in adj.ncmp + MOD_SUPERLATIVE", "arctic", []Mod{MOD_PLURAL}, WC_NOUN, true},
		{"Verb in adj.ncmp + MOD_PLURAL", "present", []Mod{MOD_PRESENT_SIMPLE, MOD_PLURAL}, WC_VERB, true},
		{"Adj in noun.unc + MOD_SUPERLATIVE", "cool", []Mod{MOD_SUPERLATIVE}, WC_ADJECTIVE, true},
		{"Adv in noun.unc + MOD_SUPERLATIVE", "cool", []Mod{MOD_SUPERLATIVE}, WC_ADVERB, true},
	}

	for _, c := range cases {
		out, err := gen.Transform(c.word, c.wc, c.mods...)

		switch c.goodCase {
		case true:
			if err != nil {
				t.Errorf("Failed for '%s - %s': error returned: '%s'", c.description, c.word, err.Error())
			}
		default:
			if err == nil {
				t.Errorf("Failed for '%s - %s': error not returned, output: %s.", c.description, c.word, out)
			}
		}
	}
}
