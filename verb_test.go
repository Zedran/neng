package neng

import "testing"

/* Tests gerund function. Fails if improper gerund form of a verb is returned. */
func TestGerund(t *testing.T) {
	cases, err := loadTestMapStringString("TestGerund.json")
	if err != nil {
		t.Fatalf("Failed loading test data: %s", err.Error())
	}

	for input, expected := range cases {
		output := gerund(input)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, expected, output)
		}
	}
}

/*
Tests pastParticiple function. Fails if improper Past Participle form of a verb is returned.
Handling of regular verbs is only symbolically checked, as it is the focus of TestPastSimpleRegular.
*/
func TestPastParticiple(t *testing.T) {
	cases, err := loadTestMapStringString("TestPastParticiple.json")
	if err != nil {
		t.Fatalf("Failed loading test data: %s", err.Error())
	}

	irregular, err := loadIrregularWords("res/verb.irr")
	if err != nil {
		t.Fatalf("loadIrregularWords failed: %s", err.Error())
	}

	for input, expected := range cases {
		output := pastParticiple(input, irregular)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, expected, output)
		}
	}
}

/* Tests pastRegular function. Fails if incorrect past tense form of a regular verb is returned. */
func TestPastRegular(t *testing.T) {
	cases, err := loadTestMapStringString("TestPastRegular.json")
	if err != nil {
		t.Fatalf("Failed loading test data: %s", err.Error())
	}

	for input, expected := range cases {
		output := pastRegular(input)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, expected, output)
		}
	}
}

/*
Tests pastSimple function. Fails if improper Past Simple form of a verb is returned.
Handling of regular verbs is only symbolically checked, as it is the focus of TestPastSimpleRegular.
*/
func TestPastSimple(t *testing.T) {
	cases, err := loadSliceTestCasePlural("TestPastSimple.json")
	if err != nil {
		t.Fatalf("Failed loading test data: %s", err.Error())
	}

	irregular, err := loadIrregularWords("res/verb.irr")
	if err != nil {
		t.Fatalf("loadIrregularWords failed: %s", err.Error())
	}

	for _, c := range cases {
		output := pastSimple(c.Input, irregular, c.Plural)

		if output != c.Expected {
			t.Errorf("Failed for '%s' (plural = %v): expected '%s', got '%s'", c.Input, c.Plural, c.Expected, output)
		}
	}
}

/* Tests presentSimple function. Fails if improper Present Simple form of a verb is returned. */
func TestPresentSimple(t *testing.T) {
	cases, err := loadSliceTestCasePlural("TestPresentSimple.json")
	if err != nil {
		t.Fatalf("Failed loading test data: %s", err.Error())
	}

	for _, c := range cases {
		output := presentSimple(c.Input, c.Plural)

		if output != c.Expected {
			number := "sing."

			if c.Plural {
				number = "pl."
			}
			t.Errorf("Failed for '%s': expected '%s' (%s), got '%s'", c.Input, c.Expected, number, output)
		}
	}
}
