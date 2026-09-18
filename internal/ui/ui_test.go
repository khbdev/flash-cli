package ui

import "testing"

func TestMask(t *testing.T) {
	tests := map[string]string{
		"":                 "",
		"abc":              "***",
		"ghp_abcdefgh1234": "********1234",
	}
	for input, want := range tests {
		if got := Mask(input); got != want {
			t.Errorf("Mask(%q) = %q, %q kutilgan", input, got, want)
		}
	}
}
