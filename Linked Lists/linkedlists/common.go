package linkedlists

import (
	"fmt"
	"io"
	"strings"
)

type val = int

type Node struct {
	Val  val
	Next *Node
}

// String implements fmt.Stringer interface
func (n Node) String() string {
	if n.Next != nil {
		return fmt.Sprintf("{%v}", n.Val)
	}
	return fmt.Sprintf("{%v}", n.Val)
}

// TODO:
// func (a *Node) Equal(b *Node) {
// }

// Print pretty prints linked list
func (list *Node) Print() {
	var sb strings.Builder
	sb.Grow(1 << 7) // random

	if list == nil {
		fmt.Fprint(&sb, "nil")
		return
	}
	fprint(&sb, *list)

	p := list.Next
	for p != nil {
		fprint(&sb, *p)
		p = p.Next
	}
	fmt.Fprint(&sb, "•")

	fmt.Println(sb.String())
}

// ToLinkedList creates linked list from slice
func ToLinkedList(ss []val) *Node {
	if ss == nil {
		return nil
	}
	if len(ss) == 0 {
		return &Node{}
	}
	if len(ss) == 1 {
		return &Node{Val: ss[0], Next: nil}
	}
	return &Node{
		Val:  ss[0],
		Next: ToLinkedList(ss[1:]),
	}
}

func fprint(w io.Writer, s fmt.Stringer) {
	fmt.Fprintf(w, "%v -> ", s)
}
