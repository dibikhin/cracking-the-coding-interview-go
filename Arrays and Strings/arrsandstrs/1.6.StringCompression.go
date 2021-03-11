package arrsandstrs

import (
	"strconv"
)

// CompressString compresses a string and returns the result if it's less than the original string.
// Supposed that strings consist of "a-zA-Z" only.
// Examples:
//     println(CompressString("zxcv"))        // zxcv
//     println(CompressString("aabcccccaaa")) // a2b1c5a3
func CompressString(s string) string {
	n := strLen(s)
	if n == 0 {
		return s
	}
	res := make([]rune, 0, n)

	c := 1
	rr := []rune(s)
	var prev, cur rune
	for i := 0; i <= n; i++ {
		prev, cur = move(i, n, prev, cur, rr)
		if prev == cur {
			c++
		} else if i > 0 {
			res = append(res, prev)
			res = append(res, []rune(strconv.Itoa(c))...)
			c = 1
		}
	}
	if len(res) >= n {
		return s
	}
	return string(res)
}

func move(i, n int, prev, cur rune, rr []rune) (rune, rune) {
	switch {
	case i == 0:
		prev = 0
		cur = rr[i]
	case i > 0 && i < n:
		prev = cur
		cur = rr[i]
	case i == n:
		prev = cur
		cur = 0
	}
	return prev, cur
}
