package linkedlists

import (
	"reflect"
	"testing"
)

func TestRemoveDups(t *testing.T) {
	tests := []struct {
		name string
		arg  *Node
		want *Node
	}{
		{"nil", (*Node)(nil), (*Node)(nil)},
		{"empty", &Node{}, &Node{}},
		{"one", &Node{Val: 1234, Next: nil}, &Node{Val: 1234, Next: nil}},

		{"2 dupes, nothing more", ToLinkedList([]int{234, 234}), &Node{Val: 234}},
		{"2 dupes, other ints", ToLinkedList([]int{456, 567, 456}), ToLinkedList([]int{456, 567})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDups(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDups() = %v, want %v", got, tt.want)
			}
		})
	}
}
