package hard

import (
	"reflect"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		lists []*ListNode
		want  *ListNode
	}{
		{nil, nil},
		{make([]*ListNode, 10), nil},
		{[]*ListNode{NewList(1, 4, 5), NewList(1, 3, 4), nil, nil, NewList(2, 6)}, NewList(1, 1, 2, 3, 4, 4, 5, 6)},
	}

	for i, tt := range tests {
		got := mergeKLists(tt.lists)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%d. got %v, want %v", i, got, tt.want)
		}
	}
}
