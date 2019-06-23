package medium

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}

	tests := []struct {
		root *TreeNode
		want []int
	}{
		{nil, nil},
		{root, []int{1, 3, 2}},
	}

	for i, tt := range tests {
		got := inorderTraversal(tt.root)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
