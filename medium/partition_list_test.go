package medium

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{}

	tests := []struct {
		head *ListNode
		x    int
		want *ListNode
	}{
		{nil, 10, nil},
		{NewList(100, 100, 1), 10, NewList(1, 100, 100)},
		{NewList(100, 100, 1), 1, NewList(100, 100, 1)},
		{NewList(1, 2, 3), 3, NewList(1, 2, 3)},
		{NewList(1, 4, 3, 2, 5, 2), 3, NewList(1, 2, 2, 4, 3, 5)},
	}

	for i, tt := range tests {
		got := partition(tt.head, tt.x)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
