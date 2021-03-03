package arrsandstrs

import (
	"strconv"
)

// The string has only "a-zA-Z"
func CompressString(s string) string {
	n := strLen(s)
	if n == 0 {
		return ""
	}
	res := make([]rune, 0, n)

	c := 1
	rr := []rune(s)
	var prev, cur rune
	for i := 0; i <= n; i++ {
		if i == 0 {
			prev = 0
			cur = rr[i]
		}
		if i > 0 && i < n {
			prev = cur
			cur = rr[i]
		}
		if i == n {
			prev = cur
			cur = 0
		}
		if prev == cur {
			c++
		}
		if prev != cur {
			if i > 0 {
				res = append(res, prev)
				res = append(res, []rune(strconv.Itoa(c))...)
				c = 1
			}
		}
	}
	if len(res) >= n {
		return s
	}
	return string(res)
}
