package neng

import "testing"

// Tests whether Mod methods return the correct results.
func TestMod(t *testing.T) {
	m := mod_undefined - 1 // Set every defined Mod value

	for i := MOD_PLURAL; i < mod_undefined; i <<= 1 {
		if !m.Enabled(i) {
			t.Errorf("Mod %d not detected", i)
		}
	}

	m++

	if !m.Undefined() {
		t.Error("Undefined Mod was not detected")
	}

	m = MOD_PLURAL | MOD_CASE_UPPER

	if !m.Enabled(MOD_PLURAL) || !m.Enabled(MOD_CASE_UPPER) {
		t.Error("Single mod not found")
	}

	if !m.Enabled(MOD_PLURAL | MOD_CASE_LOWER | MOD_GERUND) {
		t.Error("None of the group found")
	}

	if m.Enabled(MOD_GERUND) {
		t.Error("False positive reported - single")
	}

	if m.Enabled(MOD_GERUND | MOD_CASE_LOWER | MOD_PRESENT_SIMPLE) {
		t.Error("False positive reported - group")
	}
}
