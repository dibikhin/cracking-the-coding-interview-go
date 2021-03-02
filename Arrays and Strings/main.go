package main

import (
	. "aas/arrsandstrs"
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
}
