package main

import (
	. "aas/arrsandstrs"
	"fmt"
)

func main() {
	println(IsAllCharsUniqA("abc"))  // true
	println(IsAllCharsUniqB("日本本語")) // false

	println(IsPermut("cab", "bac"))  // true
	println(IsPermut("xcvb", "bac")) // false

	println(Urlify(" "))              // "%20"
	println(Urlify(" côtier côtier")) // "%20côtier%20côtier"

	println(IsPalindPermute("Tact Coa")) // true
	println(IsPalindPermute("zxcvbn"))   // false

	println(IsOneEditAway("pale", "bale")) // true
	println(IsOneEditAway("pale", "bake")) // false

	println(CompressString("zxcv"))        // zxcv
	println(CompressString("aabcccccaaa")) // a2b1c5a3

	fmt.Println(RotateMatrix([][]int{{123}}))              // [[123]]
	fmt.Println(RotateMatrix([][]int{{12, 34}, {56, 78}})) // [[56 12] [78 34]]

	// fmt.Println(RotateMatrixInPlace([][]int{{12, 34}, {56, 78}})) // [[56 12] [78 34]]
	// fmt.Println(RotateMatrixInPlace([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})) // [[7 4 1] [8 5 2] [9 6 3]]

	fmt.Println(ZeroMatrix([][]int{{123}}))             // [[123]]
	fmt.Println(ZeroMatrix([][]int{{1, 0}, {3, 4}})) // [[56 12] [78 34]]
}
