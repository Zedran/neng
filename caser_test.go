package neng

import "testing"

// Tests case transformations handled by caser.
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
		{"a man", "A man", caser.toSentence},
		{"A MAN", "A man", caser.toSentence},
		{"title", "Title", caser.toSentence},
		{"title", "Title", caser.toTitle},
		{"tItLe", "Title", caser.toTitle},
		{"a man", "A Man", caser.toTitle},
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
