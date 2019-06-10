package easy

// Merge two sorted linked lists and return it as a new list.
// The new list should be made by splicing together the nodes of the first two lists.
//
// Example:
// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := new(ListNode)
	prevNode := dummyHead

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prevNode.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			prevNode.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		}
		prevNode = prevNode.Next
	}

	if l1 != nil {
		prevNode.Next = l1
	} else if l2 != nil {
		prevNode.Next = l2
	}

	return dummyHead.Next
}
