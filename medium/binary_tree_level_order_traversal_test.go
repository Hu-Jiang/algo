package medium

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	want := [][]int{
		{3},
		{9, 20},
		{15, 7},
	}

	got := levelOrder(root)
	if !reflect.DeepEqual(got, want) {
		t.Logf("got: %v; want: %v", got, want)
	}
}
