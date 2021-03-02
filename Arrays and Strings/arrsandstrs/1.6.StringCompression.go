package arrsandstrs

import (
	"fmt"
	"strconv"
)

// The string has only "a-zA-Z"
func CompressString(s string) string {
	n := strLen(s)
	rs := make([]rune, 0, n)

	var c int
	var prev rune
	for i, cur := range s {
		if i == 0 {
			prev = cur
			c++
			continue
		}
		if prev == cur {
			c++
		}
		if prev != cur {
			rs = append(rs, prev)
			rs = append(rs, []rune(strconv.Itoa(c))...)
			c = 1
		}
		prev = cur
	}

	fmt.Println(string(rs))

	if len(rs) >= n {
		return s
	}
	return string(rs)
}
