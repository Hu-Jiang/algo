package easy

import (
	"testing"
)

func TestPathSum(t *testing.T) {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 11}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}
	root.Right.Right.Right = &TreeNode{Val: 1}

	t.Run("should has path sum", func(t *testing.T) {
		if !hasPathSum(root, 22) {
			t.Fatal("expected has path sum for 22")
		}
	})

	t.Run("should not has path sum", func(t *testing.T) {
		if hasPathSum(root, 9) {
			t.Fatalf("expected not has path sum for 9")
		}
	})
}
