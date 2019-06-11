package medium

import (
	"reflect"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	t.Run("even node", func(t *testing.T) {
		root := &ListNode{Val: 1}
		root.Next = &ListNode{Val: 2}
		root.Next.Next = &ListNode{Val: 3}
		root.Next.Next.Next = &ListNode{Val: 4}

		wantRoot := &ListNode{Val: 2}
		wantRoot.Next = &ListNode{Val: 1}
		wantRoot.Next.Next = &ListNode{Val: 4}
		wantRoot.Next.Next.Next = &ListNode{Val: 3}

		gotRoot := swapPairs(root)

		assertEqual(t, gotRoot, wantRoot)
	})

	t.Run("odd node", func(t *testing.T) {
		root := &ListNode{Val: 1}
		root.Next = &ListNode{Val: 2}
		root.Next.Next = &ListNode{Val: 3}

		wantRoot := &ListNode{Val: 2}
		wantRoot.Next = &ListNode{Val: 1}
		wantRoot.Next.Next = &ListNode{Val: 3}

		gotRoot := swapPairs(root)

		assertEqual(t, gotRoot, wantRoot)

	})

	t.Run("none node", func(t *testing.T) {
		gotRoot := swapPairs(nil)

		assertEqual(t, gotRoot, nil)
	})
}

func assertNodeEqual(t *testing.T, got, want *ListNode) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got: %v; want: %v", got, want)
	}
}
