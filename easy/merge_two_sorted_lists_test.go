package easy

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 4}

	l2 := &ListNode{Val: 1}
	l2.Next = &ListNode{Val: 3}
	l2.Next.Next = &ListNode{Val: 4}

	l1_l2 := &ListNode{Val: 1}
	l1_l2.Next = &ListNode{Val: 1}
	l1_l2.Next.Next = &ListNode{Val: 2}
	l1_l2.Next.Next.Next = &ListNode{Val: 3}
	l1_l2.Next.Next.Next.Next = &ListNode{Val: 4}
	l1_l2.Next.Next.Next.Next.Next = &ListNode{Val: 4}

	tests := []struct {
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			l1:   l1,
			l2:   l2,
			want: l1_l2,
		},
		{
			l1:   l1,
			l2:   nil,
			want: l1,
		},
		{
			l1:   nil,
			l2:   nil,
			want: nil,
		},
	}

	for i, test := range tests {
		got := mergeTwoLists(test.l1, test.l2)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("test %d, want: %+v, got: %+v", i, test.want, got)
		}
	}
}
