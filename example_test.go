package neng_test

import (
	"fmt"

	"github.com/Zedran/neng"
)

func ExampleDefaultGenerator() {
	gen, _ := neng.DefaultGenerator()

	phrase, _ := gen.Phrase("The %tsa %tpn that %m %Npv the %n")
	fmt.Println(phrase)
}

func ExampleGenerator_Adjective() {
	gen, _ := neng.DefaultGenerator()

	adj, _ := gen.Adjective(neng.MOD_COMPARATIVE)
	fmt.Println(adj)
}

func ExampleGenerator_Adverb() {
	gen, _ := neng.DefaultGenerator()

	adv, _ := gen.Adverb()
	fmt.Println(adv)
}

func ExampleGenerator_Find() {
	gen, _ := neng.DefaultGenerator()

	// Combine Find with TransformWord to efficiently
	// perform multiple transformations on a single word

	verb, _ := gen.Find("go", neng.WC_VERB)

	inf := verb.Word()
	ger, _ := gen.TransformWord(verb, neng.WC_VERB, neng.MOD_GERUND)
	pas, _ := gen.TransformWord(verb, neng.WC_VERB, neng.MOD_PAST_SIMPLE)
	pap, _ := gen.TransformWord(verb, neng.WC_VERB, neng.MOD_PAST_PARTICIPLE)
	prs, _ := gen.TransformWord(verb, neng.WC_VERB, neng.MOD_PRESENT_SIMPLE)

	fmt.Printf("%s:\ng: %s\n2: %s\n3: %s\nN: %s\n", inf, ger, pas, pap, prs)
	// Output:
	// go:
	// g: going
	// 2: went
	// 3: gone
	// N: goes
}

func ExampleGenerator_Noun() {
	gen, _ := neng.DefaultGenerator()

	noun, _ := gen.Noun(neng.MOD_PLURAL)
	fmt.Println(noun)
}

func ExampleGenerator_Phrase() {
	gen, _ := neng.DefaultGenerator()

	phrase, _ := gen.Phrase("%tpn %Npv %n")
	fmt.Println(phrase)
}

func ExampleGenerator_Transform() {
	gen, _ := neng.DefaultGenerator()

	// Suitable for one-time modification. Inefficient when transforming the same word multiple times,
	// because it searches the database for the specified string every time. Refer to Generator.Find
	// for an example of bulk transformation.

	v, _ := gen.Transform("muffin", neng.WC_NOUN, neng.MOD_PLURAL, neng.MOD_CASE_TITLE)

	fmt.Println(v)
	// Output:
	// Muffins
}

func ExampleGenerator_TransformWord() {
	gen, _ := neng.DefaultGenerator()

	w, _ := gen.Find("write", neng.WC_VERB)

	t, _ := gen.TransformWord(w, neng.WC_VERB, neng.MOD_PAST_PARTICIPLE)

	fmt.Printf("%s : %s\n", w.Word(), t)
	// Output:
	// write : written
}

func ExampleGenerator_Verb() {
	gen, _ := neng.DefaultGenerator()

	verb, _ := gen.Verb(neng.MOD_PAST_SIMPLE)
	fmt.Println(verb)
}

func ExampleNewGenerator() {
	gen, _ := neng.NewGenerator(
		[]string{"3strong"},    // Adjectives
		[]string{"4optically"}, // Adverbs
		[]string{"0moon"},      // Nouns
		[]string{"0exist"},     // Verbs
		2,                      // iterLimit
	)

	adj, _ := gen.Adjective()
	adv, _ := gen.Adverb(neng.MOD_CASE_TITLE)
	noun, _ := gen.Noun(neng.MOD_PLURAL)
	verb, _ := gen.Verb(neng.MOD_PAST_SIMPLE)

	// Without iterLimit, this call to Adverb would cause an infinite loop,
	// because the only adverb present in the database is uncomparable.
	_, err := gen.Adverb(neng.MOD_SUPERLATIVE)

	fmt.Printf("%s %s %s once %s.\n", adv, adj, noun, verb)
	fmt.Printf("Uncomparable: %v", err)
	// Output:
	// Optically strong moons once existed.
	// Uncomparable: iteration limit reached while trying to draw a valid comparative adjective or adverb
}

func ExampleNewWord() {
	w, _ := neng.NewWord("1write,wrote,written")

	fmt.Println(w.Word())
	// Output:
	// write
}

func ExampleNewWordFromParams() {
	// Regular verb
	rv, _ := neng.NewWordFromParams("create", neng.FT_REGULAR, nil)

	// Irregular verb
	iv, _ := neng.NewWordFromParams("think", neng.FT_IRREGULAR, []string{"thought", "thought"})

	// Plural-only noun
	pn, _ := neng.NewWordFromParams("odds", neng.FT_PLURAL_ONLY, nil)

	// Adjective that forms comparative and superlative by adding '-er' and '-est'
	sa, _ := neng.NewWordFromParams("strong", neng.FT_SUFFIXED, nil)

	// Uncomparable adjective
	ua, _ := neng.NewWordFromParams("tenth", neng.FT_UNCOMPARABLE, nil)

	// Uncountable noun
	un, _ := neng.NewWordFromParams("magnesium", neng.FT_UNCOUNTABLE, nil)

	for _, w := range []*neng.Word{rv, iv, pn, sa, ua, un} {
		fmt.Println(w.Word())
	}
	// Output:
	// create
	// think
	// odds
	// strong
	// tenth
	// magnesium
}

func ExampleWordClass_CompatibleWith() {
	wc := neng.WC_VERB

	fmt.Println(wc.CompatibleWith(neng.MOD_PLURAL, neng.MOD_PRESENT_SIMPLE))
	fmt.Println(wc.CompatibleWith(neng.MOD_PLURAL))
	// Output:
	// true
	// false
}
