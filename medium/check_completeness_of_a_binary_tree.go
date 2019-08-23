package medium

// Given a binary tree, determine if it is a complete binary tree.
//
// Definition of a complete binary tree from Wikipedia:
// In a complete binary tree every level, except possibly the last,
// is completely filled, and all nodes in the last level are as far
// left as possible. It can have between 1 and 2^h nodes inclusive
// at the last level h.
//
// Example 1:
// 			1
// 		  /	  \
// 		 2     3
// 	    / \	  /
//	   4   5  6
// Input: [1,2,3,4,5,6]
// Output: true
// Explanation: Every level before the last is full (ie. levels with
// node-values {1} and {2, 3}), and all nodes in the last level
// ({4, 5, 6}) are as far left as possible.
//
// Example 2:
// 			1
// 		  /	  \
// 		 2     3
// 	    / \	   	\
//	   4   5     7
// Input: [1,2,3,4,5,null,7]
// Output: false
// Explanation: The node with value 7 isn't as far left as possible.
//
// Note:
// The tree will have between 1 and 100 nodes.

func isCompleteTree(root *TreeNode) bool {
	q := []*TreeNode{root}
	for len(q) != 0 && q[0] != nil {
		node := q[0]
		q = append(q, node.Left, node.Right)
		q = q[1:]
	}

	for i := 0; i < len(q); i++ {
		if q[i] != nil {
			return false
		}
	}

	return true
}
