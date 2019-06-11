package hard

import (
	"reflect"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	t.Run("k=0", func(t *testing.T) {
		root := newList()

		gotRoot := reverseKGroup(root, 0)

		assertEqual(t, gotRoot, root)
	})

	t.Run("k=2", func(t *testing.T) {
		root := newList()

		wantRoot := &ListNode{Val: 2}
		wantRoot.Next = &ListNode{Val: 1}
		wantRoot.Next.Next = &ListNode{Val: 4}
		wantRoot.Next.Next.Next = &ListNode{Val: 3}
		wantRoot.Next.Next.Next.Next = &ListNode{Val: 5}

		gotRoot := reverseKGroup(root, 2)

		assertEqual(t, gotRoot, wantRoot)
	})

	t.Run("k=3", func(t *testing.T) {
		root := newList()

		wantRoot := &ListNode{Val: 3}
		wantRoot.Next = &ListNode{Val: 2}
		wantRoot.Next.Next = &ListNode{Val: 1}
		wantRoot.Next.Next.Next = &ListNode{Val: 4}
		wantRoot.Next.Next.Next.Next = &ListNode{Val: 5}

		gotRoot := reverseKGroup(root, 3)

		assertEqual(t, gotRoot, wantRoot)

	})

	t.Run("k=5", func(t *testing.T) {
		root := newList()

		wantRoot := &ListNode{Val: 5}
		wantRoot.Next = &ListNode{Val: 4}
		wantRoot.Next.Next = &ListNode{Val: 3}
		wantRoot.Next.Next.Next = &ListNode{Val: 2}
		wantRoot.Next.Next.Next.Next = &ListNode{Val: 1}

		gotRoot := reverseKGroup(root, 5)

		assertEqual(t, gotRoot, wantRoot)
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
