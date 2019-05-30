package medium

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tl1 := &ListNode{Val: 2}
	tl1.Next = &ListNode{Val: 4}
	tl1.Next.Next = &ListNode{Val: 3}

	tl2 := &ListNode{Val: 5}
	tl2.Next = &ListNode{Val: 6}
	tl2.Next.Next = &ListNode{Val: 4}

	tl1_tl2 := &ListNode{Val: 7}
	tl1_tl2.Next = &ListNode{Val: 0}
	tl1_tl2.Next.Next = &ListNode{Val: 8}

	tests := []struct {
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			l1:   tl1,
			l2:   tl2,
			want: tl1_tl2,
		},
		{
			l1:   tl1,
			l2:   nil,
			want: tl1,
		},
		{
			l1:   nil,
			l2:   nil,
			want: nil,
		},
	}

	for i, tt := range tests {
		got := addTwoNumbers(tt.l1, tt.l2)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("case %d: got: %+v, want: %+v", i, got, tt.want)
		}
	}
}
