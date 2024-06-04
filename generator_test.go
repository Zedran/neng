package neng

import (
	"testing"
)

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

/* Tests whether the Generator.Transform returns error upon receiving gradation modifier together with non-comparable word. */
func TestTransform(t *testing.T) {
	gen, err := DefaultGenerator()
	if err != nil {
		t.Fatalf("Failed: NewGenerator returned an error: %s", err.Error())
	}

	if output, err := gen.Transform("aa", MOD_PLURAL); err == nil {
		t.Errorf("Failed for MOD_PLURAL: no error returned. Output: '%s'", output)
	}

	if output, err := gen.Transform("own", MOD_COMPARATIVE); err == nil {
		t.Errorf("Failed for MOD_COMPARATIVE: no error returned. Output: '%s'", output)
	}

	if output, err := gen.Transform("own", MOD_SUPERLATIVE); err == nil {
		t.Errorf("Failed for MOD_SUPERLATIVE: no error returned. Output: '%s'", output)
	}
}
