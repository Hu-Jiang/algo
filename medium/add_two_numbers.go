package medium

// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order and each of their nodes contain a single digit.
// Add the two numbers and return it as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Example:

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		sentinel  = &ListNode{Val: 0}
		prevNode  = sentinel
		prevCarry = 0
	)

	doAdd := func(values ...int) {
		total := 0
		for _, v := range values {
			total += v
		}
		sum := total + prevCarry
		val := 0
		val, prevCarry = sum%10, sum/10
		prevNode.Next = &ListNode{Val: val}
		prevNode = prevNode.Next
	}

	for l1 != nil && l2 != nil {
		doAdd(l1.Val, l2.Val)
		l1 = l1.Next
		l2 = l2.Next
	}

	// maybe l1 or l2 have extra element
	for ; l1 != nil; l1 = l1.Next {
		doAdd(l1.Val)
	}

	for ; l2 != nil; l2 = l2.Next {
		doAdd(l2.Val)
	}

	if prevCarry != 0 {
		prevNode.Next = &ListNode{Val: prevCarry}
	}

	return sentinel.Next
}
