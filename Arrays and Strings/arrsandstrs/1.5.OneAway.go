package arrsandstrs

import "math"

// IsOneEditAway checks if the strings are one edit away (or zero edits). Edits are insert a character,
// remove a character, or replace a character.
func IsOneEditAway(one, two string) bool {
	if one == two {
		return true
	}
	maxLen := maxInt(strLen(one), strLen(two))

	m := make(map[rune]int, maxLen)
	countRunes(one, m)
	countRunes(two, m)

	c := 0
	for _, v := range m {
		if v%2 != 0 {
			c++
		}
		if c > 2 {
			return false
		}
	}
	return true
}

func countRunes(s string, m map[rune]int) {
	for _, r := range s {
		m[r]++
	}
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
