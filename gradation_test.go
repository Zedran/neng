package neng

import "testing"

/* Tests comparative function. Fails if incorrect comparative form is returned.*/
func TestComparative(t *testing.T) {
	cases, err := loadTestMapStringString("TestComparative.json")
	if err != nil {
		t.Fatalf("Failed loading test data: %s", err.Error())
	}

	irregular, err := loadIrregularWords("res/adj.irr")
	if err != nil {
		t.Fatalf("loadIrregularWords failed: %s", err.Error())
	}

	suffixed, err := loadWords("res/adj.suf")
	if err != nil {
		t.Fatalf("loadWords failed: %s", err.Error())
	}

	for input, expected := range cases {
		output := comparative(input, irregular, suffixed)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, expected, output)
		}
	}
}

/* Tests sufGrad function. Fails if incorrect graded form is returned. */
func TestSufGrad(t *testing.T) {
	cases, err := loadTest2DSliceString("TestSufGrad.json")
	if err != nil {
		t.Fatalf("Failed loading test data: %s", err.Error())
	}

	for _, c := range cases {
		cmp := sufGrad(c[0], "er")
		sup := sufGrad(c[0], "est")

		if cmp != c[1] || sup != c[2] {
			t.Errorf("Failed for '%s': expected '%s' - '%s', got '%s' - '%s'", c[0], c[1], c[2], cmp, sup)
		}
	}
}

/* Tests superlative function. Fails if incorrect superlative form is returned.*/
func TestSuperlative(t *testing.T) {
	cases, err := loadTestMapStringString("TestSuperlative.json")
	if err != nil {
		t.Fatalf("Failed loading test data: %s", err.Error())
	}

	irregular, err := loadIrregularWords("res/adj.irr")
	if err != nil {
		t.Fatalf("loadIrregularWords failed: %s", err.Error())
	}

	suffixed, err := loadWords("res/adj.suf")
	if err != nil {
		t.Fatalf("loadWords failed: %s", err.Error())
	}

	for input, expected := range cases {
		output := superlative(input, irregular, suffixed)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, expected, output)
		}
	}
}
