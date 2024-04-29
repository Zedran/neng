package neng

import "testing"

/* Tests gerund function. Fails if improper gerund form of a verb is returned. */
func TestGerund(t *testing.T) {
	cases := map[string]string{
		"abet":      "abetting",
		"agree":     "agreeing",
		"be":        "being",
		"become":    "becoming",
		"begin":     "beginning",
		"blossom":   "blossoming",
		"buzz":      "buzzing",
		"callus":    "callusing",
		"callous":   "callousing",
		"care":      "caring",
		"carry":     "carrying",
		"clip":      "clipping",
		"commit":    "committing",
		"degas":     "degassing",
		"dismay":    "dismaying",
		"do":        "doing",
		"dip":       "dipping",
		"dye":       "dyeing",
		"enlighten": "enlightening",
		"fidget":    "fidgeting",
		"forget":    "forgetting",
		"freeze":    "freezing",
		"gas":       "gassing",
		"go":        "going",
		"glom":      "glomming",
		"hold":      "holding",
		"interpret": "interpreting",
		"iron":      "ironing",
		"jump":      "jumping",
		"knit":      "knitting",
		"occur":     "occurring",
		"offset":    "offsetting",
		"overrun":   "overrunning",
		"quit":      "quitting",
		"panic":     "panicking",
		"pocket":    "pocketing",
		"precis":    "precising",
		"recommit":  "recommitting",
		"reposit":   "repositing",
		"shed":      "shedding",
		"sic":       "siccing",
		"sit":       "sitting",
		"sponsor":   "sponsoring",
		"stop":      "stopping",
		"suds":      "sudsing",
		"sublet":    "subletting",
		"take":      "taking",
		"talc":      "talcing",
		"tie":       "tying",
		"tighten":   "tightening",
		"underrun":  "underrunning",
		"value":     "valuing",
		"verdigris": "verdigrising",
		"vex":       "vexing",
		"zoom":      "zooming",
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
	cases := map[string]string{
		"be":      "been",
		"do":      "done",
		"forgive": "forgiven",
		"freeze":  "frozen",
		"panic":   "panicked",
	}

	irregular := [][]string{
		{"be", "was", "been"},
		{"do", "did", "done"},
		{"forgive", "forgave", "forgiven"},
		{"freeze", "froze", "frozen"},
		{"give", "gave", "given"},
	}

	for input, expected := range cases {
		output := pastParticiple(input, irregular)

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
	cases := map[string]string{
		"be":      "was",
		"do":      "did",
		"forgive": "forgave",
		"freeze":  "froze",
		"panic":   "panicked",
	}

	irregular := [][]string{
		{"be", "was", "been"},
		{"do", "did", "done"},
		{"forgive", "forgave", "forgiven"},
		{"freeze", "froze", "frozen"},
		{"give", "gave", "given"},
	}

	for input, expected := range cases {
		output := pastSimple(input, irregular)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, expected, output)
		}
	}
}

/* Tests presentSimple function. Fails if improper Present Simple form of a verb is returned. */
func TestPresentSimple(t *testing.T) {
	cases := map[string]string{
		"be":      "is",
		"buzz":    "buzzes",
		"dismiss": "dismisses",
		"dodge":   "dodges",
		"go":      "goes",
		"have":    "has",
		"honey":   "honeys",
		"learn":   "learns",
		"sic":     "sics",
		"study":   "studies",
	}

	for input, expected := range cases {
		output := presentSimple(input)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, expected, output)
		}
	}
}

/* Tests pastSimpleRegular function. Fails if improper Past Simple form of a regular verb is returned. */
func TestPastSimpleRegular(t *testing.T) {
	cases := map[string]string{
		"abet":       "abetted",
		"agree":      "agreed",
		"blossom":    "blossomed",
		"buzz":       "buzzed",
		"callus":     "callused",
		"callous":    "calloused",
		"care":       "cared",
		"carry":      "carried",
		"commission": "commissioned",
		"commit":     "committed",
		"covenant":   "covenanted",
		"degas":      "degassed",
		"dismay":     "dismayed",
		"enlighten":  "enlightened",
		"fidget":     "fidgeted",
		"interpret":  "interpreted",
		"iron":       "ironed",
		"ford":       "forded",
		"gas":        "gassed",
		"glom":       "glommed",
		"panic":      "panicked",
		"precis":     "precised",
		"recommit":   "recommitted",
		"reposit":    "reposited",
		"sic":        "sicced",
		"suds":       "sudsed",
		"sponsor":    "sponsored",
		"stop":       "stopped",
		"talc":       "talced",
		"tighten":    "tightened",
		"torpedo":    "torpedoed",
		"verdigris":  "verdigrised",
		"vex":        "vexed",
		"zoom":       "zoomed",
	}

	for input, expected := range cases {
		output := pastRegular(input)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, expected, output)
		}
	}
}
