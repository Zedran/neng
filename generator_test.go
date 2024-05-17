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
		_, err := NewGenerator(c.adj, c.adv, c.noun, c.verb)

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
	gen, err := NewGenerator([]string{"revocable"}, []string{"nicely"}, []string{"snowfall"}, []string{"stash"})
	if err != nil {
		t.Errorf("Failed: NewGenerator returned an error: %s", err.Error())
		t.FailNow()
	}

	cases := map[string]string{
		"a pretty %a %n": "a pretty revocable snowfall",
		"%a %n of %n":    "revocable snowfall of snowfall",
		"a n":            "a n",
		"%%a":            "%a",
		"a %2v %n":       "a stashed snowfall",
		"%gv":            "stashing",
		"%Nv":            "stashes",
		"%3v":            "stashed",
		"%nn":            "snowfalln",
		"%%":             "%",
		"%tNv":           "Stashes",
		"%ta %un of %ln": "Revocable SNOWFALL of snowfall",
		"%ttua":          "REVOCABLE",
		"%tpn":           "Snowfalls",
		"%upNv":          "STASH",
		"%pNv %p2v":      "stash stashed",
		"%Nv %n %m":      "stashes snowfall nicely",
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
