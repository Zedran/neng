package neng

import "testing"

/* Tests plural function. Fails if incorrect plural form of a noun is returned. */
func TestPlural(t *testing.T) {
	cases := map[string]string{
		"abaca":         "abacas",
		"agromania":     "agromania",
		"alligatorfish": "alligatorfish",
		"antic":         "antics",
		"belongings":    "belongings",
		"box":           "boxes",
		"brother":       "brothers",
		"bush":          "bushes",
		"cactus":        "cacti",
		"car":           "cars",
		"ceramics":      "ceramics",
		"codex":         "codices",
		"consciousness": "consciousness",
		"couch":         "couches",
		"cliff":         "cliffs",
		"craft":         "crafts",
		"criterion":     "criteria",
		"criterium":     "criteria",
		"daisy":         "daisies",
		"elipsis":       "elipses",
		"friendship":    "friendships",
		"dwarf":         "dwarves",
		"fez":           "fezzes",
		"fish":          "fish",
		"fix":           "fixes",
		"focus":         "foci",
		"grotto":        "grottos",
		"guess":         "guesses",
		"handcraft":     "handcraft",
		"hoof":          "hooves",
		"house":         "houses",
		"index":         "indices",
		"megahertz":     "megahertz",
		"paparazzo":     "paparazzi",
		"photo":         "photos",
		"potato":        "potatoes",
		"quartz":        "quartzes",
		"radio":         "radios",
		"ray":           "rays",
		"roof":          "roofs",
		"sesbania":      "sesbanias",
		"sheep":         "sheep",
		"ship":          "ships",
		"tax":           "taxes",
		"vertex":        "vertices",
		"virus":         "viruses",
		"volcano":       "volcanoes",
		"wife":          "wives",
		"wolf":          "wolves",
	}

	irregular, err := loadIrregularWords("res/noun.irr")
	if err != nil {
		t.Fatalf("loadIrregularWords failed: %s", err.Error())
	}

	for input, expected := range cases {
		output := plural(input, irregular)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, expected, output)
		}
	}
}
