package easy

// Given a sorted linked list, delete all duplicates such that each
// element appear only once.
//
// Example 1:
// Input: 1->1->2
// Output: 1->2
//
// Example 2:
// Input: 1->1->2->3->3
// Output: 1->2->3

func deleteDuplicates(head *ListNode) *ListNode {
	cursor := head
	for cursor != nil && cursor.Next != nil {
		if cursor.Next.Val == cursor.Val {
			cursor.Next = cursor.Next.Next
		} else {
			cursor = cursor.Next
		}
	}

	return head
}

// The official solution:
//
// Approach 1: Straight-Forward Approach
//
// public ListNode deleteDuplicates(ListNode head) {
//     ListNode current = head;
//     while (current != null && current.next != null) {
//         if (current.next.val == current.val) {
//             current.next = current.next.next;
//         } else {
//             current = current.next;
//         }
//     }
//     return head;
// }
