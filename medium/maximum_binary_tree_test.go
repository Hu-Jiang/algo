package medium

import (
	"reflect"
	"testing"
)

func TestConstructMaximumBinaryTree(t *testing.T) {
	tests := []struct {
		nums []int
		want *TreeNode
	}{
		{nil, nil},
		{[]int{1}, &TreeNode{Val: 1}},
		{[]int{3, 2, 1, 6, 0, 5}, NewPreorderTree(6, 3, -1, 2, -1, 1, -1, -1, 5, 0)},
	}

	for i, tt := range tests {
		got := constructMaximumBinaryTree(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
