package medium

// Given a linked list, remove the n-th node from the end of list and return its head.
//
// Example:
// Given linked list: 1->2->3->4->5, and n = 2.
// After removing the second node from the end, the linked list becomes 1->2->3->5.
//
// Note:
// Given n will always be valid.
//
// Follow up:
// Could you do this in one pass?

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := ListNode{Next: head}

	var length int
	for cursor := &dummyHead; cursor.Next != nil; cursor = cursor.Next {
		length++
	}

	if n <= 0 || n > length {
		return dummyHead.Next
	}

	preIdx := length - n
	preNode := &dummyHead
	for i := 0; i < preIdx; i++ {
		preNode = preNode.Next
	}

	delNode := preNode.Next
	preNode.Next = delNode.Next
	delNode.Next = nil // avoid memory leaks

	return dummyHead.Next
}

// The official solution:
//
// Approach 1: Two pass algorithm
//
// public ListNode removeNthFromEnd(ListNode head, int n) {
//     ListNode dummy = new ListNode(0);
//     dummy.next = head;
//     int length  = 0;
//     ListNode first = head;
//     while (first != null) {
//         length++;
//         first = first.next;
//     }
//     length -= n;
//     first = dummy;
//     while (length > 0) {
//         length--;
//         first = first.next;
//     }
//     first.next = first.next.next;
//     return dummy.next;
// }
//
// Approach 2: One pass algorithm
//
// public ListNode removeNthFromEnd(ListNode head, int n) {
//     ListNode dummy = new ListNode(0);
//     dummy.next = head;
//     ListNode first = dummy;
//     ListNode second = dummy;
//     // Advances first pointer so that the gap between first and second is n nodes apart
//     for (int i = 1; i <= n + 1; i++) {
//         first = first.next;
//     }
//     // Move first to the end, maintaining the gap
//     while (first != null) {
//         first = first.next;
//         second = second.next;
//     }
//     second.next = second.next.next;
//     return dummy.next;
// }
