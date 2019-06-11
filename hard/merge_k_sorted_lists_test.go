package hard

import (
	"reflect"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 5}

	l2 := &ListNode{Val: 1}
	l2.Next = &ListNode{Val: 3}
	l2.Next.Next = &ListNode{Val: 4}

	l3 := &ListNode{Val: 2}
	l3.Next = &ListNode{Val: 6}

	var lists []*ListNode
	lists = append(lists, nil, l1, l2, l3, nil)

	got := mergeKLists(lists)

	want := &ListNode{Val: 1}
	want.Next = &ListNode{Val: 1}
	want.Next.Next = &ListNode{Val: 2}
	want.Next.Next.Next = &ListNode{Val: 3}
	want.Next.Next.Next.Next = &ListNode{Val: 4}
	want.Next.Next.Next.Next.Next = &ListNode{Val: 4}
	want.Next.Next.Next.Next.Next.Next = &ListNode{Val: 5}
	want.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	assertEqual(t, got, want)

	got2 := mergeKLists(nil)
	assertEqual(t, got2, nil)

	got3 := mergeKLists(make([]*ListNode, 10))
	assertEqual(t, got3, nil)
}

func assertEqual(t *testing.T, got, want *ListNode) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got: %v; want: %v", got, want)
	}
}
