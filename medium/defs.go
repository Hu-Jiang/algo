package medium

import (
	"strconv"
)

// ListNode represents a node for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(nums ...int) *ListNode {
	var dummy ListNode
	cursor := &dummy
	for _, v := range nums {
		cursor.Next = &ListNode{Val: v}
		cursor = cursor.Next
	}
	return dummy.Next
}

func (l *ListNode) String() string {
	s, sep := "[", ""
	for ; l != nil; l = l.Next {
		s += sep + strconv.Itoa(l.Val)
		sep = " "
	}
	s += "]"

	return s
}

// TreeNode represents a node for a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewPreorderTree returns preorder tree relate to nums, nums must has not Ambiguity.
// -1 in nums represents child is null.
//
// Example 1:
// 	Input: [1, 2, 3]
// 	Generate:
//   	    1
//    	   /
//   	  2
// 	 	 /
//   	3
//
// Example 2:
// 	Input: [1, 2, -1, -1, 3] or [1, 2, -1, -1, 3, -1, -1]
// 	Generate:
//         1
//        / \
//       2   3
func NewPreorderTree(nums ...int) *TreeNode {
	var root *TreeNode
	pos := 0
	constructTree(&root, nums, &pos, len(nums))
	return root
}

func constructTree(t **TreeNode, nums []int, pos *int, len int) {
	if *pos >= len {
		return
	}

	if nums[*pos] == -1 {
		*t = nil
		*pos++
		return
	}

	*t = &TreeNode{Val: nums[*pos]}
	*pos++
	constructTree(&((*t).Left), nums, pos, len)
	constructTree(&((*t).Right), nums, pos, len)
}

func PreorderArr(t *TreeNode) []int {
	if t == nil {
		return nil
	}

	var res []int
	preorderTraversal(t, &res)
	return res
}

func preorderTraversal(t *TreeNode, res *[]int) {
	if t == nil {
		*res = append(*res, -1)
		return
	}
	*res = append(*res, t.Val)

	preorderTraversal(t.Left, res)
	preorderTraversal(t.Right, res)
}
