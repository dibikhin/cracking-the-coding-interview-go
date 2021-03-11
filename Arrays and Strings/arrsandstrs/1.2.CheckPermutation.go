package arrsandstrs

import (
	"sort"
)

// IsPermut decides if one is a permutation of the other.
func IsPermut(one, two string) bool {
	s1, s2 := []rune(one), []rune(two)
	sortRuneSliceStable(s1)
	sortRuneSliceStable(s2)
	return string(s1) == string(s2)
}

func sortRuneSliceStable(s []rune) {
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}
