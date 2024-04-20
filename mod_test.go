package neng

import "testing"

/* Tests gerund function. Fails if improper gerund form of a verb is returned. */
func TestGerund(t *testing.T) {
	cases := map[string]string{
		"agree" : "agreeing" ,
		"be"    : "being"    ,
		"care"  : "caring"   ,
		"carry" : "carrying" ,
		"do"    : "doing"    ,
		"freeze": "freezing" ,
		"go"    : "going"    ,
		"hold"  : "holding"  ,
		"panic" : "panicking",
		"sit"   : "sitting"  ,
		"stop"  : "stopping" ,
		"take"  : "taking"   ,
		"vex"   : "vexing"   ,
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
		"be"    : "been"     ,
		"do"    : "done"     ,
		"freeze": "frozen"   ,
		"panic" : "panicked" ,
	}

	irregular := map[string][]string{
		"be"    : {"was"  , "been"},
		"do"    : {"did"  , "done"},
		"freeze": {"froze", "frozen"},
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
		"be"    : "was"     ,
		"do"    : "did"     ,
		"freeze": "froze"   ,
		"panic" : "panicked",
	}

	irregular := map[string][]string{
		"be"    : {"was"  , "been"},
		"do"    : {"did"  , "done"},
		"freeze": {"froze", "frozen"},
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
		"be"     : "is"       ,
		"dismiss": "dismisses",
		"dodge"  : "dodges"   ,
		"learn"  : "learns"   ,
		"study"  : "studies"  ,
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
		"agree" : "agreed"  ,
		"care"  : "cared"   ,
		"carry" : "carried" ,
		"panic" : "panicked",
		"stop"  : "stopped" ,
		"vex"   : "vexed"   ,
	}

	for input, expected := range cases {
		output := pastSimpleRegular(input)

		if output != expected {
			t.Errorf("Failed for '%s': expected '%s', got '%s'", input, expected, output)
		}
	}
}
