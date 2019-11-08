package medium

// Given a binary tree containing digits from 0-9 only,
// each root-to-leaf path could represent a number.
//
// An example is the root-to-leaf path 1->2->3 which
// represents the number 123.
//
// Find the total sum of all root-to-leaf numbers.
//
// Note: A leaf is a node with no children.
//
// Example:
// Input: [1,2,3]
//     1
//    / \
//   2   3
// Output: 25
// Explanation:
// The root-to-leaf path 1->2 represents the number 12.
// The root-to-leaf path 1->3 represents the number 13.
// Therefore, sum = 12 + 13 = 25.
//
// Example 2:
// Input: [4,9,0,5,1]
//     4
//    / \
//   9   0
//  / \
// 5   1
// Output: 1026
// Explanation:
// The root-to-leaf path 4->9->5 represents the number 495.
// The root-to-leaf path 4->9->1 represents the number 491.
// The root-to-leaf path 4->0 represents the number 40.
// Therefore, sum = 495 + 491 + 40 = 1026.

func sumNumbers(root *TreeNode) int {
	var totalSum int
	sumHelper(root, 0, &totalSum)
	return totalSum
}

func sumHelper(root *TreeNode, prevVal int, totalSum *int) {
	if root == nil {
		return
	}

	curVal := prevVal*10 + root.Val
	if root.Left == nil && root.Right == nil { // leaf node
		*totalSum += curVal
		return
	}
	sumHelper(root.Left, curVal, totalSum)
	sumHelper(root.Right, curVal, totalSum)
}
