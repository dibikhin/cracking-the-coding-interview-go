package arrsandstrs

import (
	"strings"
)

// IsPalindPermute checks if a string is a permutation of a palin­drome. The palindrome
// does not need to be limited to just dictionary words.
// The solution below is based on an observation that palindromes have even qtys of chars (even-length ones)
func IsPalindPermute(s string) bool {
	bare := strings.ToLower(strings.ReplaceAll(s, " ", ""))
	size := strLen(bare)
	if size == 0 {
		return false
	}
	m := make(map[rune]int, size)
	for _, r := range bare {
		m[r]++
	}
	// even len
	if size%2 == 0 {
		for _, v := range m {
			if v%2 != 0 {
				return false
			}
		}
		return true
	}
	// odd len
	c := 0
	for _, v := range m {
		if v%2 != 0 {
			c++
		}
	}
	return c == 1
}

func strLen(s string) int {
	return len([]rune(s))
}
