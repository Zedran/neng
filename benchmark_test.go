// neng -- Non-Extravagant Name Generator
// Copyright (C) 2024  Wojciech Głąb (github.com/Zedran)
//
// This file is part of neng.
//
// neng is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3 only.
//
// neng is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with neng.  If not, see <https://www.gnu.org/licenses/>.

package neng

import "testing"

func BenchmarkDefaultGenerator(b *testing.B) {
	for range b.N {
		DefaultGenerator(nil)
	}
}

func BenchmarkGenerator_Phrase(b *testing.B) {
	b.StopTimer()

	gen, _ := DefaultGenerator(nil)

	b.StartTimer()
	for range b.N {
		gen.Phrase("%tsa %tpn that %m %Npv the %n")
	}
}

func BenchmarkTransformAll_Adj(b *testing.B) {
	b.StopTimer()

	gen, _ := DefaultGenerator(nil)

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

	gen, _ := DefaultGenerator(nil)
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

	gen, _ := DefaultGenerator(nil)

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

	gen, _ := DefaultGenerator(nil)
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

	gen, _ := DefaultGenerator(nil)

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

	gen, _ := DefaultGenerator(nil)
	nouns, _ := gen.Words(WC_NOUN)

	b.StartTimer()
	for range b.N {
		for w := range nouns {
			gen.TransformWord(w, WC_NOUN, MOD_INDEF)
		}
	}
}

func BenchmarkTransformAll_Noun_Plural_Possessive(b *testing.B) {
	b.StopTimer()

	gen, _ := DefaultGenerator(nil)
	nouns, _ := gen.Words(WC_NOUN)

	b.StartTimer()
	for range b.N {
		for w := range nouns {
			gen.TransformWord(w, WC_NOUN, MOD_PLURAL|MOD_POSSESSIVE)
		}
	}
}

func BenchmarkTransformAll_Noun_Possessive(b *testing.B) {
	b.StopTimer()

	gen, _ := DefaultGenerator(nil)
	nouns, _ := gen.Words(WC_NOUN)

	b.StartTimer()
	for range b.N {
		for w := range nouns {
			gen.TransformWord(w, WC_NOUN, MOD_POSSESSIVE)
		}
	}
}

func BenchmarkTransformAll_Verb(b *testing.B) {
	b.StopTimer()

	gen, _ := DefaultGenerator(nil)

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

	gen, _ := DefaultGenerator(nil)
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

	gen, _ := DefaultGenerator(nil)
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

	gen, _ := DefaultGenerator(nil)
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

	gen, _ := DefaultGenerator(nil)
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
