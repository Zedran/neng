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

// Tests WordClass.CompatibleWith. Fails if compatibility assessment
// is not consistent with documentation.
func TestWordClass_CompatibleWith(t *testing.T) {
	type testCase struct {
		expected bool
		wc       WordClass
		mods     Mod
	}

	cases := []testCase{
		{true, WC_ADJECTIVE, MOD_COMPARATIVE | MOD_SUPERLATIVE | MOD_CASE_LOWER | MOD_CASE_TITLE | MOD_CASE_UPPER},
		{true, WC_ADVERB, MOD_COMPARATIVE | MOD_SUPERLATIVE | MOD_CASE_LOWER | MOD_CASE_TITLE | MOD_CASE_UPPER},
		{true, WC_NOUN, MOD_PLURAL | MOD_CASE_LOWER | MOD_CASE_TITLE | MOD_CASE_UPPER},
		{true, WC_NOUN, MOD_POSSESSIVE},
		{true, WC_NOUN, MOD_INDEF},
		{true, WC_NOUN, MOD_INDEF_SILENT},
		{true, WC_VERB, MOD_PAST_SIMPLE | MOD_PAST_PARTICIPLE | MOD_PRESENT_SIMPLE | MOD_GERUND | MOD_PLURAL | MOD_CASE_LOWER | MOD_CASE_TITLE | MOD_CASE_UPPER},
		{true, WC_VERB, MOD_PAST_SIMPLE},
		{true, WC_VERB, MOD_PRESENT_SIMPLE},
		{false, WC_ADJECTIVE, MOD_GERUND},
		{false, WC_ADJECTIVE, MOD_PLURAL},
		{false, WC_ADJECTIVE, MOD_INDEF | MOD_SUPERLATIVE},
		{false, WC_ADJECTIVE, MOD_INDEF_SILENT},
		{false, WC_ADJECTIVE, MOD_POSSESSIVE},
		{false, WC_ADVERB, MOD_PLURAL},
		{false, WC_ADVERB, MOD_PAST_SIMPLE},
		{false, WC_ADVERB, MOD_POSSESSIVE},
		{false, WC_ADVERB, MOD_INDEF | MOD_SUPERLATIVE},
		{false, WC_ADVERB, MOD_INDEF_SILENT},
		{false, WC_NOUN, MOD_INDEF | MOD_PLURAL},
		{false, WC_NOUN, MOD_INDEF | MOD_INDEF_SILENT},
		{false, WC_NOUN, MOD_PLURAL | MOD_INDEF_SILENT},
		{false, WC_NOUN, MOD_COMPARATIVE},
		{false, WC_VERB, MOD_SUPERLATIVE},
		{false, WC_VERB, MOD_PLURAL},
		{false, WC_VERB, MOD_POSSESSIVE},
		{false, WC_VERB, MOD_INDEF},
		{false, WC_VERB, MOD_INDEF_SILENT},
	}

	for _, c := range cases {
		out := c.wc.CompatibleWith(c.mods)

		if out != c.expected {
			t.Errorf("Failed for '%v': expected %v, got %v", c, c.expected, out)
		}
	}
}
