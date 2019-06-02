package easy

import "testing"

func TestTree2Str(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 3}
		root.Left.Left = &TreeNode{Val: 4}

		got, want := tree2str(root), "1(2(4))(3)"
		if got != want {
			t.Fatalf("got: %v; want: %v", got, want)
		}
	})

	t.Run("test2", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 3}
		root.Left.Right = &TreeNode{Val: 4}

		got, want := tree2str(root), "1(2()(4))(3)"
		if got != want {
			t.Fatalf("got: %v; want: %v", got, want)
		}
	})
}
