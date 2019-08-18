package medium

import "math"

// Given the root of a binary tree, the level of its root is 1,
// the level of its children is 2, and so on.
//
// Return the smallest level X such that the sum of all the values
// of nodes at level X is maximal.
//
// Example 1:
// Input: [1,7,0,7,-8,null,null]
// Output: 2
// Explanation:
// Level 1 sum = 1.
// Level 2 sum = 7 + 0 = 7.
// Level 3 sum = 7 + -8 = -1.
// So we return the level with the maximum sum which is level 2.
//
// Note:
// The number of nodes in the given tree is between 1 and 10^4.
// -10^5 <= node.val <= 10^5

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		curMaxSum      = math.MinInt32
		curMaxSumLevel = 1
		curLevel       = 1
		q              []*TreeNode
	)
	q = append(q, root)
	for len(q) != 0 {
		size := len(q)
		curLevelSum := 0
		for i := 0; i < size; i++ {
			item := q[0]
			curLevelSum += item.Val
			if item.Left != nil {
				q = append(q, item.Left)
			}
			if item.Right != nil {
				q = append(q, item.Right)
			}
			q = q[1:]
		}
		if curLevelSum > curMaxSum {
			curMaxSum = curLevelSum
			curMaxSumLevel = curLevel
		}
		curLevel++
	}
	return curMaxSumLevel
}
