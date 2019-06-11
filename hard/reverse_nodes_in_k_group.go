package hard

import (
	"errors"
)

// Given a linked list, reverse the nodes of a linked list k at a time
// and return its modified list.
//
// k is a positive integer and is less than or equal to the length of
// the linked list. If the number of nodes is not a multiple of k then
// left-out nodes in the end should remain as it is.
//
// Example:
// Given this linked list: 1->2->3->4->5
// For k = 2, you should return: 2->1->4->3->5
// For k = 3, you should return: 3->2->1->4->5
//
// Note:
// Only constant extra memory is allowed.
// You may not alter the values in the list's nodes, only nodes itself
// may be changed.

var errLen error = errors.New("length of list not enough k")

func reverseKGroup(head *ListNode, k int) *ListNode {
	res := head
	firstSwap := true
	var pervNode *ListNode

	for {
		l, err := swapK(head, k)
		if err != nil {
			break
		}

		if firstSwap {
			res = l
			firstSwap = false
		} else {
			pervNode.Next = l
		}

		pervNode = head
		head = pervNode.Next
	}

	return res
}

func swapK(head *ListNode, k int) (*ListNode, error) {
	if !hasK(head, k) {
		return nil, errLen
	}

	if k == 1 {
		return head, nil
	}

	l1, tmp := head, head
	l2 := head.Next
	l3 := head.Next.Next

	for ; k-1 != 0; k-- {
		l2.Next = l1

		l1 = l2
		l2 = l3
		if l3 != nil {
			l3 = l3.Next
		}
	}

	tmp.Next = l2
	return l1, nil
}

// hasK returns true if length of head GE k.
// Otherwise, it returns false.
func hasK(head *ListNode, k int) bool {
	if head == nil || k == 0 {
		return false
	}

	for head != nil && k != 0 {
		head = head.Next
		k--
	}

	return k == 0
}
