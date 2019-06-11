package hard

// Merge k sorted linked lists and return it as one sorted list.
// Analyze and describe its complexity.
//
// Example:
// Input:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// Output: 1->1->2->3->4->4->5->6

func mergeKLists(lists []*ListNode) *ListNode {
	dummyHead := new(ListNode)
	pervNode := dummyHead

	removeNil(&lists)

	for {
		i := findMinIdx(lists)
		if i == -1 {
			break
		}

		pervNode.Next = &ListNode{Val: lists[i].Val}
		pervNode = pervNode.Next

		lists[i] = lists[i].Next
		if lists[i] == nil {
			lists = append(lists[:i], lists[i+1:]...)
		}
	}

	return dummyHead.Next
}

func removeNil(lists *[]*ListNode) {
	for i := 0; i < len(*lists); i++ {
		if (*lists)[i] == nil {
			*lists = append((*lists)[:i], (*lists)[i+1:]...)
			i--
		}
	}
}

// findMinIdx find minimum of head element value for every list in lists.
// If length of lists is not zero, every list in lists must be ensure not nil by invoker.
func findMinIdx(lists []*ListNode) int {
	if len(lists) == 0 {
		return -1
	}

	minVal := lists[0].Val
	minIdx := 0
	for i := 1; i < len(lists); i++ {
		if lists[i].Val < minVal {
			minVal = lists[i].Val
			minIdx = i
		}
	}

	return minIdx
}

// The official soloution:
//
// Approach 1: Brute Force
//
// class Solution(object):
//     def mergeKLists(self, lists):
//         """
//         :type lists: List[ListNode]
//         :rtype: ListNode
//         """
//         self.nodes = []
//         head = point = ListNode(0)
//         for l in lists:
//             while l:
//                 self.nodes.append(l.val)
//                 l = l.next
//         for x in sorted(self.nodes):
//             point.next = ListNode(x)
//             point = point.next
//         return head.next
//
// Approach 2: Compare one by one [no official code, but my soloution just implements it]
//
// Approach 3: Optimize Approach 2 by Priority Queue
//
// from Queue import PriorityQueue
//
// class Solution(object):
//     def mergeKLists(self, lists):
//         """
//         :type lists: List[ListNode]
//         :rtype: ListNode
//         """
//         head = point = ListNode(0)
//         q = PriorityQueue()
//         for l in lists:
//             if l:
//                 q.put((l.val, l))
//         while not q.empty():
//             val, node = q.get()
//             point.next = ListNode(val)
//             point = point.next
//             node = node.next
//             if node:
//                 q.put((node.val, node))
//         return head.next
//
// Approach 4: Merge lists one by one
//
// Approach 5: Merge with Divide And Conquer
//
// class Solution(object):
//     def mergeKLists(self, lists):
//         """
//         :type lists: List[ListNode]
//         :rtype: ListNode
//         """
//         amount = len(lists)
//         interval = 1
//         while interval < amount:
//             for i in range(0, amount - interval, interval * 2):
//                 lists[i] = self.merge2Lists(lists[i], lists[i + interval])
//             interval *= 2
//         return lists[0] if amount > 0 else lists
//
//     def merge2Lists(self, l1, l2):
//         head = point = ListNode(0)
//         while l1 and l2:
//             if l1.val <= l2.val:
//                 point.next = l1
//                 l1 = l1.next
//             else:
//                 point.next = l2
//                 l2 = l1
//                 l1 = point.next.next
//             point = point.next
//         if not l1:
//             point.next=l2
//         else:
//             point.next=l1
//         return head.next
