package neng

import "testing"

/* Tests plural function. Fails if incorrect plural form of a noun is returned. */
func TestPlural(t *testing.T) {
	cases, err := loadTestMapStringString("TestPlural.json")
	if err != nil {
		t.Fatalf("Failed loading test data: %s", err.Error())
	}

	gen, err := DefaultGenerator()
	if err != nil {
		t.Fatalf("Failed: DefaultGenerator returned an error: %s", err.Error())
	}

	for input, expected := range cases {
		word, err := gen.Find(input, WC_NOUN)
		if err != nil {
			t.Logf("Test case '%s' does not exist in the word database. Assume it is regular and proceed.", input)

			word, err = NewWordFromParams(input, 0, nil)
			if err != nil {
				t.Errorf("Failed for '%s' - error from NewWordFromParams: %v", input, err)
			}
		}

		output := plural(word)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, expected, output)
		}
	}
}
