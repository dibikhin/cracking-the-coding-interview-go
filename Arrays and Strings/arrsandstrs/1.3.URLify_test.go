package arrsandstrs

import "testing"

func Test_urlify(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty", args{""}, ""},

		{"no spaces", args{"abc"}, "abc"},

		{"1 whitespace", args{" "}, "%20"},
		{"2 spaces", args{"  "}, "%20%20"},

		{"wrapped w/ spaces", args{" abc "}, "%20abc%20"},

		{"Japanese wrapped w/ spaces", args{" 日 本 語 "}, "%20日%20本%20語%20"},
		{"French wrapped w/ spaces", args{" côtier côtier"}, "%20côtier%20côtier"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Urlify(tt.args.s); got != tt.want {
				t.Errorf("urlify() = %v, want %v", got, tt.want)
			}
		})
	}
}
