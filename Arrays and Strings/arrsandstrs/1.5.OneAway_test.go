package arrsandstrs

import "testing"

func TestIsOneEditAway(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want bool
	}{
		{"both empty", "", "", true},

		{"first empty", "", "zxcv", false},
		{"second empty", "zxcv", "", false},

		{"zero edits", "zero", "zero", true},

		{"inserted one char", "ple", "pale", true},
		{"removed one char", "pales", "pale", true},
		{"replaced one char", "pale", "bale", true},

		{"inserted one char, duped orig", "plle", "palle", true},
		{"removed one char, duped orig", "paales", "paale", true},
		{"replaced one char, duped orig", "palee", "balee", true},

		{"replaced two chars", "pale", "bake", false},

		{"replaced two chars, duped orig", "ppale", "pbake", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOneEditAway(tt.a, tt.b); got != tt.want {
				t.Errorf("IsOneEditAway() = %v, want: %v, a: %q, b: %q", got, tt.want, tt.a, tt.b)
			}
		})
	}
}
