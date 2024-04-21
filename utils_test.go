package neng

import "testing"

/* Tests whether findIrregular returns a correct verb line. */
func TestFindIrregular(t *testing.T) {
	type testCase struct {
		input             string
		expectedBool      bool
		expectedStemIndex int
	}

	cases := []string{
		"be"     ,
		"do"     ,
		"freeze" ,
		"forgive",
	}

	irregular := [][]string{
		{"be"     , "was"    , "been"    },
		{"do"     , "did"    , "done"    },
		{"forgive", "forgave", "forgiven"},
		{"freeze" , "froze"  , "frozen"  },
		{"give"   , "gave"   , "given"   },
	}

	for _, input := range cases {
		output := findIrregular(input, irregular)

		if output == nil {
			t.Errorf("Failed for '%s': nil returned", input)
		} else if output[0] != input {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, input, output)
		}
	}

	regularVerb := "panic"

	if output := findIrregular(regularVerb, irregular); output != nil {
		t.Errorf("Failed for '%s': nil not returned", regularVerb)
	}
}
