package arrsandstrs

import "testing"

func TestIsPermut(t *testing.T) {
	type args struct {
		one string
		two string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"both empty", args{"", ""}, true},
		{"one empty", args{"", "abc"}, false},
		{"another empty", args{"abc", ""}, false},
		{"same", args{"abc", "abc"}, true},
		{"permutation, no dups", args{"cab", "abc"}, true},
		{"permutation, dups", args{"cab", "aaaaabc"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPermut(tt.args.one, tt.args.two); got != tt.want {
				t.Errorf("IsPermut() = %v, want %v", got, tt.want)
			}
		})
	}
}
