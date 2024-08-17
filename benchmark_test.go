package neng

import "testing"

func BenchmarkDefaultGenerator(b *testing.B) {
	for range b.N {
		DefaultGenerator()
	}
}

func BenchmarkGenerator_Phrase(b *testing.B) {
	b.StopTimer()

	gen, _ := DefaultGenerator()

	b.StartTimer()
	for range b.N {
		gen.Phrase("%tsa %tpn that %m %Npv the %n")
	}
}

func BenchmarkTransformAll_Adj(b *testing.B) {
	b.StopTimer()

	gen, _ := DefaultGenerator()
	list := gen.adjectives
	wc := WC_ADJECTIVE

	b.StartTimer()
	for range b.N {
		for _, w := range list {
			gen.Transform(w, wc, MOD_COMPARATIVE)
			gen.Transform(w, wc, MOD_SUPERLATIVE)
		}
	}
}

func BenchmarkTransformAll_Adv(b *testing.B) {
	b.StopTimer()

	gen, _ := DefaultGenerator()
	list := gen.adverbs
	wc := WC_ADVERB

	b.StartTimer()
	for range b.N {
		for _, w := range list {
			gen.Transform(w, wc, MOD_COMPARATIVE)
			gen.Transform(w, wc, MOD_SUPERLATIVE)
		}
	}
}

func BenchmarkTransformAll_Noun(b *testing.B) {
	b.StopTimer()

	gen, _ := DefaultGenerator()
	list := gen.nouns
	wc := WC_NOUN

	b.StartTimer()
	for range b.N {
		for _, w := range list {
			gen.Transform(w, wc, MOD_PLURAL)
		}
	}
}

func BenchmarkTransformAll_Verb(b *testing.B) {
	b.StopTimer()

	gen, _ := DefaultGenerator()
	list := gen.verbs
	wc := WC_VERB

	b.StartTimer()
	for range b.N {
		for _, w := range list {
			gen.Transform(w, wc, MOD_PAST_SIMPLE)
			gen.Transform(w, wc, MOD_PAST_PARTICIPLE)
			gen.Transform(w, wc, MOD_PRESENT_SIMPLE)
			gen.Transform(w, wc, MOD_GERUND)
		}
	}
}

func BenchmarkTransformOne_Adj(b *testing.B) {
	b.StopTimer()

	gen, _ := DefaultGenerator()
	w := "little"
	wc := WC_ADJECTIVE

	b.StartTimer()
	for range b.N {
		gen.Transform(w, wc, MOD_COMPARATIVE)
		gen.Transform(w, wc, MOD_SUPERLATIVE)
	}
}

func BenchmarkTransformOne_Adv(b *testing.B) {
	b.StopTimer()

	gen, _ := DefaultGenerator()
	w := "far"
	wc := WC_ADVERB

	b.StartTimer()
	for range b.N {
		gen.Transform(w, wc, MOD_COMPARATIVE)
		gen.Transform(w, wc, MOD_SUPERLATIVE)
	}
}

func BenchmarkTransformOne_Noun(b *testing.B) {
	b.StopTimer()

	gen, _ := DefaultGenerator()
	w := "modulus"
	wc := WC_NOUN

	b.StartTimer()
	for range b.N {
		gen.Transform(w, wc, MOD_PLURAL)
	}
}

func BenchmarkTransformOne_Verb(b *testing.B) {
	b.StopTimer()

	gen, _ := DefaultGenerator()
	w := "overcome"
	wc := WC_VERB

	b.StartTimer()
	for range b.N {
		gen.Transform(w, wc, MOD_PAST_SIMPLE)
		gen.Transform(w, wc, MOD_PAST_PARTICIPLE)
		gen.Transform(w, wc, MOD_PRESENT_SIMPLE)
		gen.Transform(w, wc, MOD_GERUND)
	}
}
