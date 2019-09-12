package easy

// Given a binary tree, determine if it is height-balanced.
//
// For this problem, a height-balanced binary tree is defined as:
// a binary tree in which the depth of the two subtrees of every
// node never differ by more than 1.
//
// Example 1:
// Given the following tree [3,9,20,null,null,15,7]:
//     3
//    / \
//   9  20
//     /  \
//    15   7
// Return true.
//
// Example 2:
// Given the following tree [1,2,2,3,3,null,null,4,4]:
//        1
//       / \
//      2   2
//     / \
//    3   3
//   / \
//  4   4
// Return false.

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	lh := height(root.Left)
	rh := height(root.Right)

	if absDiff(lh, rh) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func absDiff(a, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + maxHeight(height(root.Left), height(root.Right))
}

func maxHeight(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
