package neng

import "testing"

/* Tests comparative function. Fails if incorrect comparative form is returned.*/
func TestComparative(t *testing.T) {
	cases := map[string]string{
		"good":     "better",
		"valuable": "more valuable",
	}

	irregular := [][]string{
		{"good", "better", "best"},
	}

	for input, expected := range cases {
		output := comparative(input, irregular)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, expected, output)
		}
	}
}

/* Tests superlative function. Fails if incorrect superlative form is returned.*/
func TestSuperlative(t *testing.T) {
	cases := map[string]string{
		"good":     "best",
		"valuable": "most valuable",
	}

	irregular := [][]string{
		{"good", "better", "best"},
	}

	for input, expected := range cases {
		output := superlative(input, irregular)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, expected, output)
		}
	}
}
