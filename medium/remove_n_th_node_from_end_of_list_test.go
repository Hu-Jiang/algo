package medium

import (
	"reflect"
	"testing"
)

func TestRemoveNthFormEnd(t *testing.T) {
	t.Run("remove mid", func(t *testing.T) {
		root := newList()

		gotRoot := removeNthFromEnd(root, 2)

		wantRoot := &ListNode{Val: 1}
		wantRoot.Next = &ListNode{Val: 2}
		wantRoot.Next.Next = &ListNode{Val: 3}
		wantRoot.Next.Next.Next = &ListNode{Val: 5}

		assertListEqual(t, gotRoot, wantRoot)
	})

	t.Run("remove head", func(t *testing.T) {
		root := newList()

		gotRoot := removeNthFromEnd(root, 5)

		wantRoot := &ListNode{Val: 2}
		wantRoot.Next = &ListNode{Val: 3}
		wantRoot.Next.Next = &ListNode{Val: 4}
		wantRoot.Next.Next.Next = &ListNode{Val: 5}

		assertListEqual(t, gotRoot, wantRoot)
	})

	t.Run("remove tail", func(t *testing.T) {
		root := newList()

		gotRoot := removeNthFromEnd(root, 1)

		wantRoot := &ListNode{Val: 1}
		wantRoot.Next = &ListNode{Val: 2}
		wantRoot.Next.Next = &ListNode{Val: 3}
		wantRoot.Next.Next.Next = &ListNode{Val: 4}

		assertListEqual(t, gotRoot, wantRoot)
	})
}

func newList() *ListNode {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 4}
	root.Next.Next.Next.Next = &ListNode{Val: 5}

	return root
}

func assertListEqual(t *testing.T, got, want *ListNode) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got: %v; want: %v", got, want)
	}
}
