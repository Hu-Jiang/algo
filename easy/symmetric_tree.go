package easy

// Given a binary tree, check whether it is a mirror of
// itself (ie, symmetric around its center).

// For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

//     1
//    / \
//   2   2
//  / \ / \
// 3  4 4  3

// But the following [1,2,2,null,3,null,3] is not:

//     1
//    / \
//   2   2
//    \   \
//    3    3

// Note:
// Bonus points if you could solve it both recursively and iteratively.

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return doIsSymmetric(root.Left, root.Right)
}

func doIsSymmetric(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	if t1.Val != t2.Val {
		return false
	}

	return doIsSymmetric(t1.Left, t2.Right) && doIsSymmetric(t1.Right, t2.Left)
}
