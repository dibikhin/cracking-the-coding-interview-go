package linkedlists

import (
	"reflect"
	"testing"
)

func TestToLinkedList(t *testing.T) {
	tests := []struct {
		name string
		arg  []val
		want *Node
	}{
		{"nil", []val(nil), (*Node)(nil)},
		{"empty", []val{}, &Node{}},

		{"one", []val{123}, &Node{Val: 123, Next: nil}},

		{"two", []val{123, 234}, &Node{
			Val: 123,
			Next: &Node{
				Val:  234,
				Next: nil,
			}}},
		{"three", []val{123, 234, 345}, &Node{
			Val: 123,
			Next: &Node{
				Val: 234,
				Next: &Node{
					Val:  345,
					Next: nil,
				},
			}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToLinkedList(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToLinkedList() = %v, want: %v", got, tt.want)
			}
		})
	}
}
