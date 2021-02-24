package main

import "strings"

func main() {
	println(isAllCharsUniqA("abc"))  // true
	println(isAllCharsUniqA("日本本語")) // false

	println(isAllCharsUniqB("abc"))  // true
	println(isAllCharsUniqB("日本本語")) // false
}

// IsAllCharsUniqA implements an algorithm to determine if a string has all unique characters.
func isAllCharsUniqA(s string) bool {
	m := make(map[rune]int)
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

// IsAllCharsUniqB gives the same result as IsAllCharsUniqA but without extra data structs
func isAllCharsUniqB(s string) bool {
	for _, v := range s {
		ix := strings.Index(s, string(v))
		lix := strings.LastIndex(s, string(v))
		if lix > ix {
			return false
		}
	}
	return true
}
