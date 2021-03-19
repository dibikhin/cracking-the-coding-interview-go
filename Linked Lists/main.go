package main

import (
	. "ll/linkedlists"
)

func main() {
	var nilNode *Node
	nilNode.Print() // nil

	empty := &Node{}
	empty.Print() // {0} -> •

	single := &Node{Val: 986}
	single.Print() // {986} -> •

	z := &Node{
		Val: 123,
		Next: &Node{
			Val: 234,
			Next: &Node{
				Val:  345,
				Next: nil,
			},
		},
	}
	z.Print()

	b := []int{987, 876, 765, 654, 543, 432, 321}
	x := ToLinkedList(b)
	x.Print()

	dupes := ToLinkedList([]int{123, 345, 345, 234})
	dupes.Print() // {123} -> {345} -> {345} -> {234} -> •
	uniq := RemoveDups(dupes)
	uniq.Print() // {123} -> {345} -> {234} -> •
}
