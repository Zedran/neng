package neng

import (
	"testing"

	"github.com/Zedran/neng/internal/tests"
)

func TestIndefinite(t *testing.T) {
	var cases map[string]string
	if err := tests.ReadData("TestIndefinite.json", &cases); err != nil {
		t.Fatalf("Failed loading test data: %v", err)
	}

	for input, expected := range cases {
		out := indefinite(input)
		if out != expected {
			t.Errorf("Failed for %s: expected %s, got %s.", input, expected, out)
		}
	}
}
