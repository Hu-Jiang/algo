package easy

// Given a binary tree, check whether it is a mirror of
// itself (ie, symmetric around its center).
//
// For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
//
//     1
//    / \
//   2   2
//  / \ / \
// 3  4 4  3
//
// But the following [1,2,2,null,3,null,3] is not:
//
//     1
//    / \
//   2   2
//    \   \
//    3    3
//
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

// Trick:
// func doIsSymmetric(t1, t2 *TreeNode) bool {
// 	if t1 == nil || t2 == nil {
// 		return t1 == t2
// 	}
//
// 	return t1.Val == t2.Val &&
// 		doIsSymmetric(t1.Left, t2.Right) &&
// 		doIsSymmetric(t1.Right, t2.Left)
// }
//
// The official solution:
//
// Approach 1: Recursive
//
// public boolean isSymmetric(TreeNode root) {
//     return isMirror(root, root);
// }
//
// public boolean isMirror(TreeNode t1, TreeNode t2) {
//     if (t1 == null && t2 == null) return true;
//     if (t1 == null || t2 == null) return false;
//     return (t1.val == t2.val)
//         && isMirror(t1.right, t2.left)
//         && isMirror(t1.left, t2.right);
// }
//
// Approach 2: Iterative
//
// public boolean isSymmetric(TreeNode root) {
//     Queue<TreeNode> q = new LinkedList<>();
//     q.add(root);
//     q.add(root);
//     while (!q.isEmpty()) {
//         TreeNode t1 = q.poll();
//         TreeNode t2 = q.poll();
//         if (t1 == null && t2 == null) continue;
//         if (t1 == null || t2 == null) return false;
//         if (t1.val != t2.val) return false;
//         q.add(t1.left);
//         q.add(t2.right);
//         q.add(t1.right);
//         q.add(t2.left);
//     }
//     return true;
// }
