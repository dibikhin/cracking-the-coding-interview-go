package arrsandstrs

import (
	"strings"
)

// IsAllCharsUniqA implements an algorithm to determine if a string has all unique characters.
func IsAllCharsUniqA(s string) bool {
	m := make(map[rune]int, strLen(s))
	for _, v := range s {
		m[v]++
	}
	for _, v := range m {
		if v > 1 {
			return false
		}
	}
	return true
}

func IsAllCharsUniqA_1(s string) bool {
	m := make(map[rune]int, strLen(s))
	for _, v := range s {
		m[v]++
		if m[v] > 1 {
			return false
		}
	}
	return true
}

// IsAllCharsUniqB gives the same result as IsAllCharsUniqA but without using extra data structs
func IsAllCharsUniqB(s string) bool {
	for _, v := range s {
		ix := strings.Index(s, string(v))
		lix := strings.LastIndex(s, string(v))
		if lix > ix {
			return false
		}
	}
	return true
}
