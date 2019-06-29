package easy

import (
	"reflect"
	"testing"
)

func TestNewList(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		l := NewList()
		assertEqual(t, l, (*ListNode)(nil))
	})

	t.Run("normal", func(t *testing.T) {
		l := NewList(1, 2, 3)
		wantRoot := &ListNode{Val: 1}
		wantRoot.Next = &ListNode{Val: 2}
		wantRoot.Next.Next = &ListNode{Val: 3}
		assertEqual(t, l, wantRoot)
	})
}

func TestNewPreorderTree(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		root := NewPreorderTree()
		assertEqual(t, root, (*TreeNode)(nil))
	})

	t.Run("only left", func(t *testing.T) {
		root := NewPreorderTree(1, 2, 3, -1, -1, -1, -1, -1, -1, -1)
		wantRoot := &TreeNode{Val: 1}
		wantRoot.Left = &TreeNode{Val: 2}
		wantRoot.Left.Left = &TreeNode{Val: 3}
		assertEqual(t, root, wantRoot)
	})

	t.Run("only right", func(t *testing.T) {
		root := NewPreorderTree(1, -1, 2, -1, 3)
		wantRoot := &TreeNode{Val: 1}
		wantRoot.Right = &TreeNode{Val: 2}
		wantRoot.Right.Right = &TreeNode{Val: 3}
		assertEqual(t, root, wantRoot)
	})

	t.Run("normal", func(t *testing.T) {
		root := NewPreorderTree(1, 2, -1, -1, 3)
		wantRoot := &TreeNode{Val: 1}
		wantRoot.Left = &TreeNode{Val: 2}
		wantRoot.Right = &TreeNode{Val: 3}
		assertEqual(t, root, wantRoot)
	})
}

func TestPreorderArr(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		root := NewPreorderTree()
		got := PreorderArr(root)
		assertEqual(t, got, ([]int)(nil))
	})

	t.Run("only left", func(t *testing.T) {
		root := NewPreorderTree(1, 2, 3)
		got := PreorderArr(root)
		want := []int{1, 2, 3, -1, -1, -1, -1}
		assertEqual(t, got, want)
	})

	t.Run("only right", func(t *testing.T) {
		root := NewPreorderTree(1, -1, 2, -1, 3)
		got := PreorderArr(root)
		want := []int{1, -1, 2, -1, 3, -1, -1}
		assertEqual(t, got, want)
	})

	t.Run("normal", func(t *testing.T) {
		root := NewPreorderTree(1, 2, -1, -1, 3, -1, -1)
		got := PreorderArr(root)
		want := []int{1, 2, -1, -1, 3, -1, -1}
		assertEqual(t, got, want)
	})

}

func assertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
