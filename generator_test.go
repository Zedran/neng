package neng

import "testing"

/* Tests whether Generator.Phrase correctly parses syntax and generates phrases. */
func TestGetPhrase(t *testing.T) {
	gen := Generator{
		adjectives: []string{"revocable"},
		nouns     : []string{"snowfall"},
	}

	cases := map[string]string{
		"a pretty %a %n": "a pretty revocable snowfall"   ,
		"%a %n of %n"   : "revocable snowfall of snowfall",
		"a n"           : "a n"                           ,
		"%%a"           : "%a"                            ,
	}

	for input, expected := range cases {
		output, err := gen.Phrase(input)
		if err != nil {
			t.Errorf("Failed for case '%s': error returned: '%s'", input, err.Error())
		}else if output != expected {
			t.Errorf("Failed for case '%s': got '%s', expected '%s'", input, output, expected)
		}
	}

	errCases := []string{
		""    , // pattern is empty
		"abc%", // escape character at pattern termination
		"%q"  , // unknown command
	}

	for _, bc := range errCases {
		output, err := gen.Phrase(bc)
		if err == nil {
			t.Errorf("Failed for errCase '%s': no error returned. Output: '%s'", bc, output)
		}
	}
}
