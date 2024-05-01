package neng

import "testing"

/* Tests gerund function. Fails if improper gerund form of a verb is returned. */
func TestGerund(t *testing.T) {
	cases := map[string]string{
		"abet":      "abetting",
		"abhor":     "abhorring",
		"agree":     "agreeing",
		"alibi":     "alibiing",
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
		"shanghai":  "shanghaiing",
		"shed":      "shedding",
		"sic":       "siccing",
		"sit":       "sitting",
		"ski":       "skiing",
		"sponsor":   "sponsoring",
		"stop":      "stopping",
		"suds":      "sudsing",
		"sublet":    "subletting",
		"take":      "taking",
		"talc":      "talcing",
		"taxi":      "taxiing",
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
		"alibi":    "alibis",
		"be":       "is",
		"buzz":     "buzzes",
		"dismiss":  "dismisses",
		"dodge":    "dodges",
		"go":       "goes",
		"have":     "has",
		"honey":    "honeys",
		"learn":    "learns",
		"shanghai": "shanghais",
		"sic":      "sics",
		"ski":      "skis",
		"study":    "studies",
		"taxi":     "taxis",
	}

	for input, expected := range cases {
		output := presentSimple(input)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, expected, output)
		}
	}
}

/* Tests pastRegular function. Fails if incorrect past tense form of a regular verb is returned. */
func TestPastRegular(t *testing.T) {
	cases := map[string]string{
		"abet":       "abetted",
		"abhor":      "abhorred",
		"agree":      "agreed",
		"alibi":      "alibied",
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
		"shanghai":   "shanghaied",
		"sic":        "sicced",
		"ski":        "skied",
		"suds":       "sudsed",
		"sponsor":    "sponsored",
		"stop":       "stopped",
		"taxi":       "taxied",
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
