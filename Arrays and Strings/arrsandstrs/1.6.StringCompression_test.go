package arrsandstrs

import "testing"

func TestCompressString(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"empty", "", ""},

		{"one char", "b", "b"},

		{"no dups", "zxcv", "zxcv"},

		{"dups", "aabcccccaaa", "a2b1c5a3"},
		{"dups, diff case", "AAAaCCCCCaaa", "A3a1C5a3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompressString(tt.arg); got != tt.want {
				t.Errorf("CompressString() = %v, want %v, arg: %q", got, tt.want, tt.arg)
			}
		})
	}
}
