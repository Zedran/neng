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

// WordClass helps the Generator to differentiate parts of speech.
type WordClass uint8

const (
	WC_ADJECTIVE WordClass = iota
	WC_ADVERB
	WC_NOUN
	WC_VERB
)

// CompatibleWith returns true if WordClass is compatible with all of the
// received mods. This method tests defined Mod values only - undefined Mod
// values have undefined compatibility. Use Mod.Undefined to ensure that
// all bits in Mod have defined values.
func (wc WordClass) CompatibleWith(mods Mod) bool {
	switch wc {
	case WC_ADJECTIVE, WC_ADVERB:
		if mods.Enabled(MOD_PLURAL | MOD_PAST_SIMPLE | MOD_PAST_PARTICIPLE | MOD_PRESENT_SIMPLE | MOD_GERUND | MOD_INDEF_SILENT) {
			return false
		}
		if mods.Enabled(MOD_INDEF) && mods.Enabled(MOD_SUPERLATIVE) {
			return false
		}
	case WC_NOUN:
		if mods.Enabled(MOD_PAST_SIMPLE | MOD_PAST_PARTICIPLE | MOD_PRESENT_SIMPLE | MOD_GERUND | MOD_COMPARATIVE | MOD_SUPERLATIVE) {
			return false
		}
		if mods.Enabled(MOD_INDEF) && mods.Enabled(MOD_PLURAL) {
			return false
		}
		if mods.Enabled(MOD_INDEF_SILENT) && mods.Enabled(MOD_PLURAL|MOD_INDEF) {
			return false
		}
	case WC_VERB:
		if mods.Enabled(MOD_INDEF | MOD_COMPARATIVE | MOD_SUPERLATIVE | MOD_INDEF_SILENT) {
			return false
		}
		if mods.Enabled(MOD_PLURAL) && !mods.Enabled(MOD_PAST_SIMPLE|MOD_PRESENT_SIMPLE) {
			return false
		}
	}
	return true
}
