package medium

// Given a linked list, swap every two adjacent nodes and return its head.
//
// You may not modify the values in the list's nodes, only nodes itself may be changed.
//
// Example:
// Given 1->2->3->4, you should return the list as 2->1->4->3.

func swapPairs(head *ListNode) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head

	pervNode := dummyHead
	for pervNode.Next != nil && pervNode.Next.Next != nil {
		l1 := pervNode.Next
		l2 := l1.Next
		l3 := l2.Next

		pervNode.Next = l2
		l2.Next = l1
		l1.Next = l3

		pervNode = pervNode.Next.Next
	}

	return dummyHead.Next
}
