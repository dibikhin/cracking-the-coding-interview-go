package main

import (
	"testing"
)

type args struct {
	s string
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"empty", args{""}, true},

	{"unique English", args{"abcdef"}, true},
	{"dups in row English", args{"abcdefc"}, false},
	{"dups English", args{"abccccdef"}, false},

	{"unique French", args{"côtier"}, true},
	{"dups in row French", args{"côtttier"}, false},
	{"dups French", args{"côtierô"}, false},

	{"unique Japanese", args{"日本語"}, true},
	{"dups in row Japanese", args{"日本本本本語"}, false},
	{"dups Japanese", args{"日本日語"}, false},
}

func Test_isAllCharsUniqA(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAllCharsUniqA(tt.args.s); got != tt.want {
				t.Errorf("isAllCharsUniq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAllCharsUniqB(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAllCharsUniqB(tt.args.s); got != tt.want {
				t.Errorf("isAllCharsUniqB() = %v, want %v", got, tt.want)
			}
		})
	}
}
