package medium

// Sort a linked list in O(n log n) time using constant space complexity.
//
// Example 1:
// Input: 4->2->1->3
// Output: 1->2->3->4
//
// Example 2:
// Input: -1->5->3->4->0
// Output: -1->0->3->4->5

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast, slow := head.Next.Next, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	mid := slow.Next
	slow.Next = nil

	return mergeList(sortList(head), sortList(mid))
}

func mergeList(l1, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: 0, Next: nil}
	cursor := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cursor.Next = l1
			l1 = l1.Next
		} else {
			cursor.Next = l2
			l2 = l2.Next
		}
		cursor = cursor.Next
	}

	if l1 != nil {
		cursor.Next = l1
	}

	if l2 != nil {
		cursor.Next = l2
	}

	return dummyHead.Next
}
