package hard

import "math"

// Given a non-empty binary tree, find the maximum path sum.
//
// For this problem, a path is defined as any sequence of nodes
// from some starting node to any node in the tree along the
// parent-child connections. The path must contain at least one
// node and does not need to go through the root.
//
// Example 1:
// Input: [1,2,3]
//        1
//       / \
//      2   3
// Output: 6
//
// Example 2:
// Input: [-10,9,20,null,null,15,7]
//    -10
//    / \
//   9  20
//     /  \
//    15   7
// Output: 42

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	helper(root, &res)
	return res
}

func helper(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	left := helper(root.Left, res)
	right := helper(root.Right, res)

	*res = max(left+right+root.Val, *res)

	return max(0, max(left, right)+root.Val)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
