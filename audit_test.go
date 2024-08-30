package neng

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

/*
TestAudit applies all possible grammatical transformations to every word in the default database and stores the results
within 'audit' directory. This process creates a standardized snapshot of all transformed words, which can be used
to evaluate transformation quality and analyze the impact of new features and modifications.

This function does not run automatically along other tests. A special environment variable 'AUDIT' must be set
(`AUDIT=1 go test` on Linux) in order to trigger its execution.
*/
func TestAudit(t *testing.T) {
	const (
		AUDIT_DIR string = "audit"
		FORMAT    string = "'%s': %v"
	)

	if os.Getenv("AUDIT") == "" {
		t.Skip("Env 'AUDIT' not set. Skipping TestAudit.")
	}

	createFile := func(fname string) (*os.File, error) {
		f, err := os.Create(filepath.Join(AUDIT_DIR, fname))
		if err != nil {
			return nil, fmt.Errorf("%s: %w", fname, err)
		}
		return f, nil
	}

	gen, err := DefaultGenerator()
	if err != nil {
		t.Fatalf("Failed: NewGenerator returned an error: %v", err)
	}

	if err = os.Mkdir(AUDIT_DIR, 0755); err != nil && !os.IsExist(err) {
		t.Fatalf("Could not create audit directory: %v", err)
	}

	fAdj, err := createFile("adj")
	if err != nil {
		t.Fatal(err)
	}
	defer fAdj.Close()

	fAdv, err := createFile("adv")
	if err != nil {
		t.Fatal(err)
	}
	defer fAdv.Close()

	fNoun, err := createFile("noun")
	if err != nil {
		t.Fatal(err)
	}
	defer fNoun.Close()

	fVerb, err := createFile("verb")
	if err != nil {
		t.Fatal(err)
	}
	defer fVerb.Close()

	for _, a := range gen.adj {
		cmp, err := gen.TransformWord(a, WC_ADJECTIVE, MOD_COMPARATIVE)
		if err != nil {
			if !errors.Is(err, errNonComparable) {
				t.Fatalf(FORMAT, a.word, err)
			}
			cmp = "ncmp"
		}

		sup, err := gen.TransformWord(a, WC_ADJECTIVE, MOD_SUPERLATIVE)
		if err != nil {
			if !errors.Is(err, errNonComparable) {
				t.Fatalf(FORMAT, a.word, err)
			}
			sup = "ncmp"
		}

		fAdj.WriteString(fmt.Sprintf("%s\n%s\n%s\n\n", a.word, cmp, sup))
	}

	for _, m := range gen.adv {
		cmp, err := gen.TransformWord(m, WC_ADVERB, MOD_COMPARATIVE)
		if err != nil {
			if !errors.Is(err, errNonComparable) {
				t.Fatalf(FORMAT, m.word, err)
			}
			cmp = "ncmp"
		}

		sup, err := gen.TransformWord(m, WC_ADVERB, MOD_SUPERLATIVE)
		if err != nil {
			if !errors.Is(err, errNonComparable) {
				t.Fatalf(FORMAT, m.word, err)
			}
			sup = "ncmp"
		}

		fAdv.WriteString(fmt.Sprintf("%s\n%s\n%s\n\n", m.word, cmp, sup))
	}

	for _, n := range gen.noun {
		pl, err := gen.TransformWord(n, WC_NOUN, MOD_PLURAL)
		if err != nil {
			if !errors.Is(err, errUncountable) {
				t.Fatalf(FORMAT, n.word, err)
			}
			pl = "unc"
		}
		fNoun.WriteString(fmt.Sprintf("%s\n%s\n\n", n.word, pl))
	}

	for _, v := range gen.verb {
		pastSimple, err := gen.TransformWord(v, WC_VERB, MOD_PAST_SIMPLE)
		if err != nil {
			t.Fatalf(FORMAT, v.word, err)
		}

		pastParticiple, err := gen.TransformWord(v, WC_VERB, MOD_PAST_PARTICIPLE)
		if err != nil {
			t.Fatalf(FORMAT, v.word, err)
		}

		presSimple, err := gen.TransformWord(v, WC_VERB, MOD_PRESENT_SIMPLE)
		if err != nil {
			t.Fatalf(FORMAT, v.word, err)
		}

		gerund, err := gen.TransformWord(v, WC_VERB, MOD_GERUND)
		if err != nil {
			t.Fatalf(FORMAT, v.word, err)
		}

		fVerb.WriteString(fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n\n", v.word, pastSimple, pastParticiple, presSimple, gerund))
	}
}
