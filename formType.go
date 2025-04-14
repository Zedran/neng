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

// FormType (formation type). Indicates the effect that
// grammatical transformations have on a given word.
type FormType uint8

const (
	// A regular word:
	//   - adj, adv - forms comparative and superlative by adding
	//                'more' and 'most' before itself
	//   - noun     - can be both singular and plural
	//   - verb     - has regular past tense forms
	FT_REGULAR FormType = iota

	// An irregular word. It has its own special forms for:
	//   - adj, adv - comparative, superlative
	//   - noun     - plural
	//   - verb     - Past Simple, Past Participle
	FT_IRREGULAR

	// A plural-only noun (e.g. scissors). It does not get picked
	// in absence of MOD_PLURAL.
	FT_PLURAL_ONLY

	// Adjective or adverb graded by appending '-er' and '-est' suffixes.
	FT_SUFFIXED

	// Non-comparable adjective or adverb. It does not get picked
	// if MOD_COMPARATIVE or MOD_SUPERLATIVE is requested.
	// An attempt to grade a non-comparable word results in an error.
	FT_NON_COMPARABLE

	// Uncountable noun. It does not get picked if MOD_PLURAL is requested.
	// An attempt to pluralize an uncountable noun results in an error.
	FT_UNCOUNTABLE
)
