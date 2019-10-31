package medium

// Reverse a linked list from position m to n. Do it in one-pass.
//
// Note: 1 ≤ m ≤ n ≤ length of list.
//
// Example:
// Input: 1->2->3->4->5->NULL, m = 2, n = 4
// Output: 1->4->3->2->5->NULL

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	prevM := &ListNode{Next: head}
	for i := 1; i < m; i++ {
		prevM = prevM.Next
	}

	first := prevM.Next
	second := first.Next
	if second == nil {
		return head
	}
	three := second.Next
	for i := 0; i < n-m; i++ {
		second.Next = first
		first, second = second, three
		if three != nil {
			three = three.Next
		}
	}

	prevM.Next.Next, prevM.Next = second, first

	if m == 1 {
		head = prevM.Next
	}
	return head
}

// The official solution:
//
// Approach 1: Recursion
//
// class Solution {
//
//     // Object level variables since we need the changes
//     // to persist across recursive calls and Java is pass by value.
//     private boolean stop;
//     private ListNode left;
//
//     public void recurseAndReverse(ListNode right, int m, int n) {
//
//         // base case. Don't proceed any further
//         if (n == 1) {
//             return;
//         }
//
//         // Keep moving the right pointer one step forward until (n == 1)
//         right = right.next;
//
//         // Keep moving left pointer to the right until we reach the proper node
//         // from where the reversal is to start.
//         if (m > 1) {
//             this.left = this.left.next;
//         }
//
//         // Recurse with m and n reduced.
//         this.recurseAndReverse(right, m - 1, n - 1);
//
//         // In case both the pointers cross each other or become equal, we
//         // stop i.e. don't swap data any further. We are done reversing at this
//         // point.
//         if (this.left == right || right.next == this.left) {
//             this.stop = true;
//         }
//
//         // Until the boolean stop is false, swap data between the two pointers
//         if (!this.stop) {
//             int t = this.left.val;
//             this.left.val = right.val;
//             right.val = t;
//
//             // Move left one step to the right.
//             // The right pointer moves one step back via backtracking.
//             this.left = this.left.next;
//         }
//     }
//
//     public ListNode reverseBetween(ListNode head, int m, int n) {
//         this.left = head;
//         this.stop = false;
//         this.recurseAndReverse(head, m, n);
//         return head;
//     }
// }
//
// Approach 2: Iterative Link Reversal.
//
// class Solution {
//     public ListNode reverseBetween(ListNode head, int m, int n) {
//
//         // Empty list
//         if (head == null) {
//             return null;
//         }
//
//         // Move the two pointers until they reach the proper starting point
//         // in the list.
//         ListNode cur = head, prev = null;
//         while (m > 1) {
//             prev = cur;
//             cur = cur.next;
//             m--;
//             n--;
//         }
//
//         // The two pointers that will fix the final connections.
//         ListNode con = prev, tail = cur;
//
//         // Iteratively reverse the nodes until n becomes 0.
//         ListNode third = null;
//         while (n > 0) {
//             third = cur.next;
//             cur.next = prev;
//             prev = cur;
//             cur = third;
//             n--;
//         }
//
//         // Adjust the final connections as explained in the algorithm
//         if (con != null) {
//             con.next = prev;
//         } else {
//             head = prev;
//         }
//
//         tail.next = cur;
//         return head;
//     }
// }
