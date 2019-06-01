package medium

import "strconv"

// ListNode represents a node for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
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
