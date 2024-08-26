package neng

import (
	"slices"
	"testing"
)

/*
Ensures that call to DefaultGenerator does not return an error
and tests whether word lists of the resulting instance are sorted.
*/
func TestDefaultGenerator(t *testing.T) {
	gen, err := DefaultGenerator()
	if err != nil {
		t.Fatalf("Failed: DefaultGenerator returned an error: %s", err.Error())
	}

	for i, wl := range [][]*Word{gen.adj, gen.adv, gen.noun, gen.verb} {
		if !slices.IsSortedFunc(wl, cmpWord) {
			t.Fatalf("Failed for list %d - not sorted", i)
		}
	}
}

/* Tests whether Generator.Find correctly returns found words or errors upon failure. */
func TestGenerator_Find(t *testing.T) {
	type testCase struct {
		good  bool
		wc    WordClass
		query string
	}

	testCases := []testCase{
		{true, WC_ADJECTIVE, "big"},         // Existing adjective
		{true, WC_ADVERB, "nicely"},         // Existing adverb
		{true, WC_NOUN, "snowfall"},         // Existing noun
		{true, WC_VERB, "stash"},            // Existing verb
		{false, WC_NOUN, "box"},             // Missing noun
		{false, WC_NOUN, "big"},             // Missing "noun", present in other lists
		{false, WordClass(255), "snowfall"}, // Undefined WordClass
	}

	gen, err := NewGenerator([]string{"3big"}, []string{"0nicely"}, []string{"0snowfall"}, []string{"0stash"}, DEFAULT_ITER_LIMIT, false)
	if err != nil {
		t.Fatalf("Failed: NewGenerator returned an error: %s", err.Error())
	}

	for _, c := range testCases {
		out, err := gen.Find(c.query, c.wc)

		if c.good {
			switch true {
			case err != nil:
				t.Errorf("Failed for case %v: error returned: '%s'", c, err.Error())
			case out.word != c.query:
				t.Errorf("Failed for case %v: expected word '%s', found '%s'", c, c.query, out.word)
			}
		} else {
			if err == nil {
				t.Errorf("Failed for case %v: no error returned, got %v", c, out)
			}
		}
	}
}

/*
Tests whether Generator.Noun correctly skips uncountable nouns in presence of MOD_PLURAL
and plural-only nouns in absence of plural modifier.
*/
func TestGenerator_Noun(t *testing.T) {
	gen, err := NewGenerator([]string{"3big"}, []string{"0nicely"}, []string{"2binoculars"}, []string{"0stash"}, 10, false)
	if err != nil {
		t.Fatalf("Failed: NewGenerator returned an error: %s", err.Error())
	}

	if n, err := gen.Noun(); err == nil {
		t.Errorf("Failed for singular: plural-only noun was not rejected. Noun returned: %s", n)
	}

	if _, err = gen.Noun(MOD_PLURAL); err != nil {
		t.Errorf("Failed for plural: plural-only noun was rejected: %s", err.Error())
	}

	gen.noun = []*Word{{ft: FT_UNCOUNTABLE, irr: nil, word: "boldness"}}

	if n, err := gen.Noun(MOD_PLURAL); err == nil {
		t.Errorf("Failed for plural: uncountable noun was not rejected. Noun returned: %s", n)
	}

	if _, err = gen.Noun(); err != nil {
		t.Errorf("Failed for singular: uncountable noun was rejected: %s", err.Error())
	}
}

/* Tests whether Generator.Phrase correctly parses pattern syntax and generates phrases. */
func TestGenerator_Phrase(t *testing.T) {
	gen, err := NewGenerator([]string{"3big"}, []string{"0nicely"}, []string{"0snowfall"}, []string{"0stash"}, DEFAULT_ITER_LIMIT, false)
	if err != nil {
		t.Fatalf("Failed: NewGenerator returned an error: %s", err.Error())
	}

	var cases map[string]string
	if err := loadTestData("TestPhrase.json", &cases); err != nil {
		t.Fatalf("Failed loading test data: %v", err)
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
		"",     // Pattern is empty
		"abc%", // Escape character at pattern termination
		"%q",   // Unknown command
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
Tests basic dispatching done by Generator.Transform. More detailed tests are performed for Generator.TransformWord,
which receives input from Generator.Transform.
*/
func TestGenerator_Transform(t *testing.T) {
	type testCase struct {
		good     bool
		word     string
		expected string
		wc       WordClass
	}

	testCases := []testCase{
		{true, "word", "word", WC_NOUN},    // Existing
		{false, "theta", "theta", WC_NOUN}, // Non-existent
	}

	gen, err := DefaultGenerator()
	if err != nil {
		t.Fatalf("Failed: DefaultGenerator returned an error: %s", err.Error())
	}

	for _, c := range testCases {
		out, err := gen.Transform(c.word, c.wc)

		if c.good {
			if err != nil {
				t.Errorf("Failed for '%v': error returned: '%s'", c, err.Error())
			}

			if out != c.expected {
				t.Errorf("Failed for '%v': got '%s' - expected '%s'", c, out, c.expected)
			}
		} else {
			if err == nil {
				t.Errorf("Failed for '%v': error not returned", c)
			}
		}
	}
}

/*
Tests whether the Generator.TransformWord correctly returns errIncompatible, errNonComparable and errUncountable.
errIncompatible should be returned if the requested modification is incompatible with a given WordClass.
errNonComparable should only be returned if gradation was requested for non-comparable adjective or adverb and
errUncountable should only be returned if pluralization was requested for an uncountable noun.
*/
func TestGenerator_TransformWord(t *testing.T) {
	type testCase struct {
		description string
		word        string
		mods        []Mod
		wc          WordClass
		goodCase    bool
	}

	gen, err := DefaultGenerator()
	if err != nil {
		t.Fatalf("Failed: DefaultGenerator returned an error: %s", err.Error())
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
	}

	for _, c := range cases {
		word, err := gen.Find(c.word, c.wc)
		if err != nil {
			t.Fatalf("'%s' (WordClass %d) does not exist in the word database.", c.word, c.wc)
		}

		out, err := gen.TransformWord(word, c.wc, c.mods...)

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
	gen, err := NewGenerator([]string{"4bottomless"}, []string{"4cryptographically"}, []string{"0snowfall"}, []string{"0stash"}, 10, false)
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

/*
Tests NewGenerator function. Fails if it does not return an error upon receiving an empty list, nil or invalid iterLimit value.
Malformed or empty slice elements trigger an error as well. Takes safe value into account during testing. Slice order is not
checked - NewGenerator calls NewGeneratorFromWord which handles sorting.
*/
func TestNewGenerator(t *testing.T) {
	type testCase struct {
		adj, adv, noun, verb []string
		iterLimit            int
		safe                 bool
		goodCase             bool
	}

	var (
		empty  = []string{}
		good   = []string{"0word"}
		has0   = []string{"0word", ""}
		hasBad = []string{"1word"}
	)

	cases := []testCase{
		{good, good, good, good, 1, true, true},               // Lines present in every slice
		{empty, good, good, good, 1, false, true},             // Empty list, safe false
		{good, good, nil, good, 1, false, true},               // nil pointer, safe false
		{good, good, nil, good, 1, true, false},               // Error: nil pointer, safe true
		{good, has0, good, good, 1, true, false},              // Error: contains zero-length element, safe true
		{good, has0, good, good, 1, false, false},             // Error: contains zero-length element, safe false
		{empty, good, good, good, 1, true, false},             // Error: no adjectives
		{good, empty, good, good, 1, true, false},             // Error: no adverbs
		{good, good, empty, good, 1, true, false},             // Error: no nouns
		{good, good, good, empty, 1, true, false},             // Error: no verbs
		{empty, empty, empty, empty, 1, true, false},          // Error: empty slices only
		{nil, nil, nil, nil, DEFAULT_ITER_LIMIT, true, false}, // Error: nil pointers only
		{good, good, good, good, 0, true, false},              // Error: iterLimit == 0
		{good, good, good, good, -5, true, false},             // Error: negative iterLimit
		{hasBad, good, good, good, 1, true, false},            // Error: malformed adjective
		{good, hasBad, good, good, 1, true, false},            // Error: malformed adverb
		{good, good, hasBad, good, 1, true, false},            // Error: malformed noun
		{good, good, good, hasBad, 1, true, false},            // Error: malformed verb
	}

	for _, c := range cases {
		_, err := NewGenerator(c.adj, c.adv, c.noun, c.verb, c.iterLimit, c.safe)

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

/* Tests NewGeneratorFromWord ensuring that the intended checks correctly trigger errors. */
func TestNewGeneratorFromWord(t *testing.T) {
	type testCase struct {
		adj, adv, noun, verb []*Word
		iterLimit            int
		safe                 bool
		goodCase             bool
	}

	w, err := NewWord("0word")
	if err != nil {
		t.Fatalf("Failed: NewWord returned an error: %s", err.Error())
	}

	w2, err := NewWord("4a")
	if err != nil {
		t.Fatalf("Failed: NewWord returned an error: %s", err.Error())
	}

	badW := &Word{ft: 255, irr: &[]string{"too", "many", "forms"}, word: "z"}

	var (
		good   = []*Word{w, w}
		empty  = []*Word{}
		hasNil = []*Word{w, nil}
		hasBad = []*Word{w, badW}
		unsort = []*Word{w, w2}
	)

	cases := []testCase{
		{good, good, good, good, 1, true, true},               // Lines present in every slice
		{good, good, good, hasNil, 1, false, true},            // One of the slices contains nil, but safe is false
		{good, empty, good, good, 1, false, true},             // One of the slices is empty, but safe is false
		{good, hasBad, good, good, 1, true, true},             // One of the slices has an invalid element, safe is true
		{good, hasBad, good, good, 1, false, true},            // One of the slices has an invalid element, safe is false
		{unsort, good, good, good, 1, true, true},             // Contains unsorted slice, safe is true
		{unsort, good, good, good, 1, false, true},            // Contains unsorted slice, but safe is false
		{good, good, nil, good, 1, true, false},               // Error: nil pointer
		{empty, good, good, good, 1, true, false},             // Error: No adjectives
		{good, empty, good, good, 1, true, false},             // Error: No adverbs
		{good, good, empty, good, 1, true, false},             // Error: No nouns
		{good, good, good, empty, 1, true, false},             // Error: No verbs
		{empty, empty, empty, empty, 1, true, false},          // Error: Empty slices only
		{nil, nil, nil, nil, DEFAULT_ITER_LIMIT, true, false}, // Error: nil pointers only
		{good, good, good, good, 0, true, false},              // Error: iterLimit == 0
		{good, good, good, good, -5, true, false},             // Error: Negative iterLimit
		{good, good, good, good, -5, false, false},            // Error: Negative iterLimit, safe is false
		{good, good, good, hasNil, 1, true, false},            // Error: One of the slices contains nil
	}

	for i, c := range cases {
		_, err := NewGeneratorFromWord(c.adj, c.adv, c.noun, c.verb, c.iterLimit, c.safe)

		switch c.goodCase {
		case true:
			if err != nil {
				t.Errorf("Failed for case %d: NewGenerator returned an error: %v", i, err)
			}

			if c.safe {
				for _, wl := range [][]*Word{c.adj, c.adv, c.noun, c.verb} {
					if len(wl) == 0 {
						t.Errorf("Failed for case %d: NewGenerator allowed empty list", i)
					}

					if !slices.IsSortedFunc(wl, cmpWord) {
						t.Errorf("Failed for case %d: NewGenerator did not sort the list (safe == %v)", i, c.safe)
					}
				}
			}

		default:
			if err == nil {
				t.Errorf("Failed for case '%d': NewGenerator did not return an error.", i)
			}
		}
	}
}
