package hard

import "strconv"

// ListNode represents a node for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(nums ...int) *ListNode {
	var dummy ListNode
	cursor := &dummy
	for _, v := range nums {
		cursor.Next = &ListNode{Val: v}
		cursor = cursor.Next
	}
	return dummy.Next
}

func (l *ListNode) String() string {
	s, sep := "[", ""
	for ; l != nil; l = l.Next {
		s += sep + strconv.Itoa(l.Val)
		sep = " "
	}
	s += "]"

	return s
}
