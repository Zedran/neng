package neng

import "testing"

/*
Tests whether Generator.Noun correctly skips uncountable nouns in presence of MOD_PLURAL
and plural-only nouns in absence of plural modifier.
*/
func TestGenerator_Noun(t *testing.T) {
	gen, err := NewGenerator([]string{"big"}, []string{"nicely"}, []string{"binoculars"}, []string{"stash"}, 10)
	if err != nil {
		t.Fatalf("Failed: NewGenerator returned an error: %s", err.Error())
	}

	if n, err := gen.Noun(); err == nil {
		t.Errorf("Failed for singular: plural-only noun was not rejected. Noun returned: %s", n)
	}

	if _, err = gen.Noun(MOD_PLURAL); err != nil {
		t.Errorf("Failed for plural: plural-only noun was rejected: %s", err.Error())
	}

	gen.nouns = []string{"boldness"}

	if n, err := gen.Noun(MOD_PLURAL); err == nil {
		t.Errorf("Failed for plural: uncountable noun was not rejected. Noun returned: %s", n)
	}

	if _, err = gen.Noun(); err != nil {
		t.Errorf("Failed for singular: uncountable noun was rejected: %s", err.Error())
	}
}

/* Tests whether Generator.Phrase correctly parses pattern syntax and generates phrases. */
func TestGenerator_Phrase(t *testing.T) {
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
		"%cn",  // WordClass-Mod incompatibility
	}

	for _, bc := range errCases {
		output, err := gen.Phrase(bc)
		if err == nil {
			t.Errorf("Failed for errCase '%s': no error returned. Output: '%s'", bc, output)
		}
	}
}

/*
Tests whether the Generator.Transform correctly returns errIncompatible, errNonComparable and errUncountable.
errIncompatible should be returned if the requested modification is incompatible with a given WordClass.
errNonComparable should only be returned if gradation was requested for non-comparable adjective or adverb and
errUncountable should only be returned if pluralization was requested for an uncountable noun.
*/
func TestGenerator_Transform(t *testing.T) {
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
		{"WordClass-Mod incompatibility", "aa", []Mod{MOD_COMPARATIVE}, WC_NOUN, false},
		{"Uncountable noun + MOD_PLURAL", "aa", []Mod{MOD_PLURAL}, WC_NOUN, false},
		{"Non-comparable adj + MOD_COMPARATIVE", "own", []Mod{MOD_COMPARATIVE}, WC_ADJECTIVE, false},
		{"Non-comparable adj + MOD_SUPERLATIVE", "own", []Mod{MOD_SUPERLATIVE}, WC_ADJECTIVE, false},
		{"Non-comparable adv + MOD_COMPARATIVE", "cryptographically", []Mod{MOD_COMPARATIVE}, WC_ADVERB, false},
		{"Non-comparable adv + MOD_SUPERLATIVE", "cryptographically", []Mod{MOD_SUPERLATIVE}, WC_ADVERB, false},
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

/* Tests whether Generator.generateModifier correctly skips non-comparable adjectives if gradation is requested. */
func TestGenerator_generateModifier(t *testing.T) {
	gen, err := NewGenerator([][]byte{[]byte("bottomless 4")}, [][]byte{[]byte("cryptographically 4")}, [][]byte{[]byte("snowfall 0")}, [][]byte{[]byte("stash 0")}, 10)
	if err != nil {
		t.Fatalf("Failed: NewGenerator returned an error: %s", err.Error())
	}

	if _, err = gen.Adjective(); err != nil {
		t.Errorf("Failed for positive: non-comparable adjective was rejected: %s", err.Error())
	}

	if a, err := gen.Adjective(MOD_COMPARATIVE); err == nil {
		t.Errorf("Failed for comparative: non-comparable adjective was not rejected. Adjective returned: %s", a)
	}

	if a, err := gen.Adjective(MOD_SUPERLATIVE); err == nil {
		t.Errorf("Failed for superlative: non-comparable adjective was not rejected. Adjective returned: %s", a)
	}

	if _, err = gen.Adverb(); err != nil {
		t.Errorf("Failed for positive: non-comparable adverb was rejected: %s", err.Error())
	}

	if a, err := gen.Adverb(MOD_COMPARATIVE); err == nil {
		t.Errorf("Failed for comparative: non-comparable adverb was not rejected. Adverb returned: %s", a)
	}

	if a, err := gen.Adverb(MOD_SUPERLATIVE); err == nil {
		t.Errorf("Failed for superlative: non-comparable adverb was not rejected. Adverb returned: %s", a)
	}
}

/* Tests NewGenerator function. Fails if providing an empty list, nil or an invalid iterLimit value does not trigger an error. */
func TestNewGenerator(t *testing.T) {
	type testCase struct {
		adj, adv, noun, verb [][]byte
		iterLimit            int
		goodCase             bool
	}

	var (
		good  = [][]byte{[]byte("word 0")}
		empty = [][]byte{}
	)

	cases := []testCase{
		{good, good, good, good, 1, true},               // Words present in every slice
		{good, good, nil, good, 1, false},               // nil pointer
		{empty, good, good, good, 1, false},             // No adjectives
		{good, empty, good, good, 1, false},             // No adverbs
		{good, good, empty, good, 1, false},             // No nouns
		{good, good, good, empty, 1, false},             // No verbs
		{empty, empty, empty, empty, 1, false},          // Empty slices only
		{nil, nil, nil, nil, DEFAULT_ITER_LIMIT, false}, // nil pointers only
		{good, good, good, good, 0, false},              // iterLimit == 0
		{good, good, good, good, -5, false},             // Negative iterLimit
	}

	for _, c := range cases {
		_, err := NewGenerator(c.adj, c.adv, c.noun, c.verb, c.iterLimit)

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
