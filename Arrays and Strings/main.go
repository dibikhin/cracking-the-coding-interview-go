package main

import (
	. "aas/arrsandstrs"
)

func main() {
	println(IsAllCharsUniqA("abc"))  // true
	println(IsAllCharsUniqB("日本本語")) // false

	println(IsPermut("cab", "bac")) // true
	println(IsPermut("xcvb", "bac")) // false
}
