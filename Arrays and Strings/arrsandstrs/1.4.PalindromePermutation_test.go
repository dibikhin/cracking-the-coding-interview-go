package arrsandstrs

import "testing"

func TestIsPalindPermute(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{"empty", "", false},
		{"whitespace", "     ", false},

		{"not a palindrome", "sdfge", false},

		{"digits, one word", "9753579", true},

		{"digits, two words, even length", "864 468", true},

		{"not a palindrome, digits, two words, even length", "863 468", false},

		{"a palindrome, one word", "Kayak", true},
		{"a palindrome, 2 words", "Tact Coa", true},

		{"a palindrome, one word, French", "étêté", true},

		{"a palindrome, one word, Chinese", "中山王族王山中", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindPermute(tt.arg); got != tt.want {
				t.Errorf("IsPalindPermute() = %v, want: %v, arg: %q", got, tt.want, tt.arg)
			}
		})
	}
}
