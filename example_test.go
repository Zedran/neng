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

func ExampleNewGenerator() {
	gen, _ := neng.NewGenerator(
		[]string{"big"},      // adjectives
		[]string{"nicely"},   // adverbs
		[]string{"snowfall"}, // nouns
		[]string{"stash"},    // verbs
		neng.DEFAULT_ITER_LIMIT,
	)

	phrase, _ := gen.Phrase("a %a %n")
	fmt.Println(phrase)
	// Output: a big snowfall
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
	var (
		gen, _   = neng.DefaultGenerator()
		word, _  = gen.Verb()
		modified string
	)

	fmt.Println(word)

	for _, mod := range []neng.Mod{neng.MOD_PRESENT_SIMPLE, neng.MOD_PAST_SIMPLE, neng.MOD_PAST_PARTICIPLE, neng.MOD_GERUND} {
		modified, _ = gen.Transform(word, neng.WC_VERB, mod)
		fmt.Println(modified)
	}
}

func ExampleGenerator_Verb() {
	gen, _ := neng.DefaultGenerator()

	verb, _ := gen.Verb(neng.MOD_PAST_SIMPLE)
	fmt.Println(verb)
}

func ExampleWordClass_CompatibleWith() {
	wc := neng.WC_VERB

	fmt.Println(wc.CompatibleWith(neng.MOD_PLURAL, neng.MOD_PRESENT_SIMPLE))
	fmt.Println(wc.CompatibleWith(neng.MOD_PLURAL))
	// Output:
	// true
	// false
}
