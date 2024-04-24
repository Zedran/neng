package neng

import "testing"

/* Tests casing transformations handled by caser struct. */
func TestCaser(t *testing.T) {
	type testCase struct {
		input    string
		expected string
		function func(string) string
	}

	caser := newCaser()

	cases := []testCase{
		{"LOWER", "lower", caser.toLower},
		{"LoWeR", "lower", caser.toLower},
		{"title", "Title", caser.toTitle},
		{"tItLe", "Title", caser.toTitle},
		{"upper", "UPPER", caser.toUpper},
		{"uPpEr", "UPPER", caser.toUpper},
	}

	for _, c := range cases {
		output := c.function(c.input)

		if output != c.expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", c.input, c.expected, output)
		}
	}
}
