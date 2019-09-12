package medium

import (
	"reflect"
	"testing"
)

func TestSortedListToBST(t *testing.T) {
	tests := []struct {
		head *ListNode
		want *TreeNode
	}{
		{
			head: nil,
			want: nil,
		},
		{
			head: NewList(-10, -3, 0, 5, 9),
			want: NewPreorderTree(0, -3, -10, -1, -1, -1, 9, 5),
		},
		{
			head: NewList(0, 1, 2, 3, 4, 5),
			want: NewPreorderTree(3, 1, 0, -1, -1, 2, -1, -1, 5, 4),
		},
	}

	for i, tt := range tests {
		got := sortedListToBST(tt.head)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
