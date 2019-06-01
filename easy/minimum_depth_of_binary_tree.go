package easy

// Given a binary tree, find its minimum depth.
//
// The minimum depth is the number of nodes along the
// shortest path from the root node down to the nearest leaf node.
//
// Note: A leaf is a node with no children.
//
// Example:
//
// Given binary tree [3,9,20,null,null,15,7],
//
//     3
//    / \
//   9  20
//     /  \
//    15   7
//
// return its minimum depth = 2.

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return minSmart(minDepth(root.Left), minDepth(root.Right)) + 1
}

// minSmart returns a if b equals zero, or b if a equals zero.
// otherwise it returns the minimum of a or b.
func minSmart(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a < b {
		return a
	}
	return b
}
