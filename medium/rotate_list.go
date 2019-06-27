package medium

// Given a linked list, rotate the list to the right by k places, where k is non-negative.
//
// Example 1:
// Input: 1->2->3->4->5->NULL, k = 2
// Output: 4->5->1->2->3->NULL
// Explanation:
// rotate 1 steps to the right: 5->1->2->3->4->NULL
// rotate 2 steps to the right: 4->5->1->2->3->NULL
//
// Example 2:
// Input: 0->1->2->NULL, k = 4
// Output: 2->0->1->NULL
// Explanation:
// rotate 1 steps to the right: 2->0->1->NULL
// rotate 2 steps to the right: 1->2->0->NULL
// rotate 3 steps to the right: 0->1->2->NULL
// rotate 4 steps to the right: 2->0->1->NULL

func rotateRight(head *ListNode, k int) *ListNode {
	length := listLen(head)
	if length <= 1 || k == 0 {
		return head
	}

	rotate := k % length
	if rotate == 0 {
		return head
	}

	skip := length - rotate
	prev := head
	for i := 0; i < skip-1; i++ {
		prev = prev.Next
	}
	newHead := prev.Next
	prev.Next = nil
	last := newHead
	for last.Next != nil {
		last = last.Next
	}
	last.Next = head
	return newHead
}

func listLen(head *ListNode) int {
	len := 0
	for head != nil {
		len++
		head = head.Next
	}

	return len
}
