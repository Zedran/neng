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

	list := make([]string, len(gen.adj))
	for i, w := range gen.adj {
		list[i] = w.word
	}

	wc := WC_ADJECTIVE

	b.StartTimer()
	for range b.N {
		for _, w := range list {
			f, _ := gen.Find(w, wc)

			gen.TransformWord(f, wc, MOD_COMPARATIVE)
			gen.TransformWord(f, wc, MOD_SUPERLATIVE)
		}
	}
}

func BenchmarkTransformAll_Adj_Indef(b *testing.B) {
	b.StopTimer()

	gen, _ := DefaultGenerator()
	nouns, _ := gen.Words(WC_ADJECTIVE)

	b.StartTimer()
	for range b.N {
		for w := range nouns {
			gen.TransformWord(w, WC_ADJECTIVE, MOD_INDEF)
		}
	}
}

func BenchmarkTransformAll_Adv(b *testing.B) {
	b.StopTimer()

	gen, _ := DefaultGenerator()

	list := make([]string, len(gen.adv))
	for i, w := range gen.adv {
		list[i] = w.word
	}

	wc := WC_ADVERB

	b.StartTimer()
	for range b.N {
		for _, w := range list {
			f, _ := gen.Find(w, wc)

			gen.TransformWord(f, wc, MOD_COMPARATIVE)
			gen.TransformWord(f, wc, MOD_SUPERLATIVE)
		}
	}
}

func BenchmarkTransformAll_Adv_Indef(b *testing.B) {
	b.StopTimer()

	gen, _ := DefaultGenerator()
	nouns, _ := gen.Words(WC_ADVERB)

	b.StartTimer()
	for range b.N {
		for w := range nouns {
			gen.TransformWord(w, WC_ADVERB, MOD_INDEF)
		}
	}
}

func BenchmarkTransformAll_Noun(b *testing.B) {
	b.StopTimer()

	gen, _ := DefaultGenerator()

	list := make([]string, len(gen.noun))
	for i, w := range gen.noun {
		list[i] = w.word
	}

	wc := WC_NOUN

	b.StartTimer()
	for range b.N {
		for _, w := range list {
			f, _ := gen.Find(w, wc)

			gen.TransformWord(f, wc, MOD_PLURAL)
		}
	}
}

func BenchmarkTransformAll_Noun_Indef(b *testing.B) {
	b.StopTimer()

	gen, _ := DefaultGenerator()
	nouns, _ := gen.Words(WC_NOUN)

	b.StartTimer()
	for range b.N {
		for w := range nouns {
			gen.TransformWord(w, WC_NOUN, MOD_INDEF)
		}
	}
}

func BenchmarkTransformAll_Verb(b *testing.B) {
	b.StopTimer()

	gen, _ := DefaultGenerator()

	list := make([]string, len(gen.verb))
	for i, w := range gen.verb {
		list[i] = w.word
	}

	wc := WC_VERB

	b.StartTimer()
	for range b.N {
		for _, w := range list {
			f, _ := gen.Find(w, wc)

			gen.TransformWord(f, wc, MOD_PAST_SIMPLE)
			gen.TransformWord(f, wc, MOD_PAST_PARTICIPLE)
			gen.TransformWord(f, wc, MOD_PRESENT_SIMPLE)
			gen.TransformWord(f, wc, MOD_GERUND)
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
		f, _ := gen.Find(w, wc)

		gen.TransformWord(f, wc, MOD_COMPARATIVE)
		gen.TransformWord(f, wc, MOD_SUPERLATIVE)
	}
}

func BenchmarkTransformOne_Adv(b *testing.B) {
	b.StopTimer()

	gen, _ := DefaultGenerator()
	w := "far"
	wc := WC_ADVERB

	b.StartTimer()
	for range b.N {
		f, _ := gen.Find(w, wc)

		gen.TransformWord(f, wc, MOD_COMPARATIVE)
		gen.TransformWord(f, wc, MOD_SUPERLATIVE)
	}
}

func BenchmarkTransformOne_Noun(b *testing.B) {
	b.StopTimer()

	gen, _ := DefaultGenerator()
	w := "modulus"
	wc := WC_NOUN

	b.StartTimer()
	for range b.N {
		f, _ := gen.Find(w, wc)

		gen.TransformWord(f, wc, MOD_PLURAL)
	}
}

func BenchmarkTransformOne_Verb(b *testing.B) {
	b.StopTimer()

	gen, _ := DefaultGenerator()
	w := "overcome"
	wc := WC_VERB

	b.StartTimer()
	for range b.N {
		f, _ := gen.Find(w, wc)

		gen.TransformWord(f, wc, MOD_PAST_SIMPLE)
		gen.TransformWord(f, wc, MOD_PAST_PARTICIPLE)
		gen.TransformWord(f, wc, MOD_PRESENT_SIMPLE)
		gen.TransformWord(f, wc, MOD_GERUND)
	}
}
