package linkedlists

import "fmt"

type val = int

type Node struct {
	Val  val
	Next *Node
}

func (n *Node) Print() {
	printNode(n)

	p := n.Next
	for p != nil {
		printNode(p)
		p = n.Next.Next
	}
	if p == nil {
		fmt.Print(p)
	}
	println()
}

func printNode(n *Node) {
	fmt.Printf("%v -> ", *n)
}

func ToList(s []val) Node {
	p := &Node{Val: s[0], Next: nil}
	for i := 1; i < len(s); i++ {
		p.Next = &Node{Val: s[i], Next: nil}

		p = p.Next
	}
}
