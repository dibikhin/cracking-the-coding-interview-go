package linkedlists

type void = struct{}

// RemoveDups removes duplicates from an unsorted linked list.
// The target list is modified. Additional data structs are used.
// Example:
// 	dupes := ToLinkedList([]int{123, 345, 345, 234})
// 	dupes.Print() // {123 *linkedlists.Node} -> {345 *linkedlists.Node} -> {345 *linkedlists.Node} -> {234 nil} -> •
// 	uniq := RemoveDups(dupes)
// 	uniq.Print() // {345 *linkedlists.Node} -> {234 *linkedlists.Node} -> {123 nil} -> •
func RemoveDups(list *Node) *Node {
	if list == nil {
		return nil
	}
	if *list == (Node{}) {
		return list
	}
	if list.Next == nil {
		return list
	}
	set := toSet(list)
	keys := toList(set)
	return ToLinkedList(keys)
}

func toSet(n *Node) map[int]void {
	set := make(map[val]void)
	for p := n; p != nil; p = p.Next {
		set[p.Val] = void{}
	}
	return set
}

func toList(set map[int]struct{}) []int {
	keys := make([]val, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}
	return keys
}
