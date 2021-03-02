package arrsandstrs

import (
	"strings"
)

// Urlify replaces whitespaces w/ "%20" in place using []rune
func Urlify(s string) string {
	const p = "%20"

	c := strings.Count(s, " ")
	r := []rune(s)
	rs := make([]rune, len(r)+c*2) // adding margin for shifting
	copy(rs, r)

	ix := index(rs, ' ')
	for ix > -1 {
		copy(rs[ix+2:], rs[ix:]) // shift

		rs[ix] = '%'
		rs[ix+1] = '2'
		rs[ix+2] = '0'

		ix = index(rs, ' ')
	}
	return string(rs)
}

func index(rs []rune, r rune) int {
	for i, v := range rs {
		if v == r {
			return i
		}
	}
	return -1
}
