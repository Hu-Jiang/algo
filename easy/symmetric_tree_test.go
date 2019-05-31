package easy

import "testing"

func TestSymmetricTree(t *testing.T) {
	t.Run("should be symmetric", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 2}
		root.Left.Left = &TreeNode{Val: 3}
		root.Left.Right = &TreeNode{Val: 4}
		root.Right.Left = &TreeNode{Val: 4}
		root.Right.Right = &TreeNode{Val: 3}

		if !isSymmetric(root) {
			t.Fatalf("got: %v; want: %v", isSymmetric(root), true)
		}
	})

	t.Run("should not be symmetric", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 2}
		root.Left.Right = &TreeNode{Val: 3}
		root.Right.Right = &TreeNode{Val: 3}

		if isSymmetric(root) {
			t.Fatalf("got: %v; want: %v", isSymmetric(root), false)
		}
	})
}
