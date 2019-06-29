package medium

// Given a sorted linked list, delete all nodes that have duplicate
// numbers, leaving only distinct numbers from the original list.
//
// Example 1:
// Input: 1->2->3->3->4->4->5
// Output: 1->2->5
//
// Example 2:
// Input: 1->1->1->2->3
// Output: 2->3

func deleteDuplicates(head *ListNode) *ListNode {
	dummyHead := ListNode{Next: head}
	prev := &dummyHead
	cursor := dummyHead.Next
	for cursor != nil && cursor.Next != nil {
		p := cursor
		for p != nil && p.Next != nil && p.Val == p.Next.Val {
			p = p.Next
		}
		if p != cursor {
			prev.Next = p.Next
			cursor = p.Next
		} else {
			prev = cursor
			cursor = cursor.Next
		}
	}
	return dummyHead.Next
}

// Recursion 1:
//
// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
//
// 	if head.Val == head.Next.Val {
// 		for head.Next != nil && head.Val == head.Next.Val {
// 			head = head.Next
// 		}
// 		return deleteDuplicates(head.Next)
// 	}
//
// 	head.Next = deleteDuplicates(head.Next)
// 	return head
// }
//
// Recursion 2:
//
// func deleteDuplicates(head *ListNode) *ListNode {
//     if head == nil      { return nil }
//     if head.Next == nil { return head }
//
//     p := head.Next
//     for p != nil && p.Val == head.Val {
//         p = p.Next
//     }
//     if p == head.Next {               // no dupliates at head
//         head.Next = deleteDuplicates(p)
//         return head
//     } else {                          // duplicates! dump head
//         return deleteDuplicates(p)
//     }
// }
