package neng

import "testing"

/* Tests whether findIrregular returns a correct index. */
func TestFindIrregular(t *testing.T) {
	type testCase struct {
		input             string
		expectedBool      bool
		expectedStemIndex int
	}

	cases := map[string]int{ // irregular stem verbs:
		"be"     :  0,       // be
		"do"     :  1,       // do
		"freeze" :  2,       // freeze
		"forgive":  3,       // give
		"panic"  : -1,       // regular verb
	}

	irregular := [][]string{
		{"be"     , "was"    , "been"    },
		{"do"     , "did"    , "done"    },
		{"freeze" , "froze"  , "frozen"  },
		{"give"   , "gave"   , "given"   },
	}

	for input, expected := range cases {
		output := findIrregular(input, irregular)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%d', got '%d'", input, expected, output)
		}
	}
}
