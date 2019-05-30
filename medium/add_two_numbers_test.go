package main

import (
	"reflect"
	"testing"
)

func TestAddTwoNumber(t *testing.T) {
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
		l1       *ListNode
		l2       *ListNode
		wantList *ListNode
	}{
		{
			l1:       tl1,
			l2:       tl2,
			wantList: tl1_tl2,
		},
		{
			l1:       tl1,
			l2:       nil,
			wantList: tl1,
		},
		{
			l1:       nil,
			l2:       nil,
			wantList: nil,
		},
	}

	for i, test := range tests {
		if !reflect.DeepEqual(test.wantList, addTwoNumbers(test.l1, test.l2)) {
			t.Fatalf("test %d, l1: %+v, l2: %+v, want: %+v, got: %+v", i, test.l1, test.l2, test.wantList, addTwoNumbers(test.l1, test.l2))
		}
	}
}
