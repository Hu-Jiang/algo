package easy

// Given a binary tree and a sum, determine if the tree has a
// root-to-leaf path such that adding up all the values along
// the path equals the given sum.
//
// Note: A leaf is a node with no children.
//
// Example:
//
// Given the below binary tree and sum = 22,
//
//       5
//      / \
//     4   8
//    /   / \
//   11  13  4
//  /  \      \
// 7    2      1
//
// return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.

func hasPathSum(root *TreeNode, sum int) bool {
	return doHasPathSum(root, 0, sum)

}

func doHasPathSum(node *TreeNode, curSum, wantSum int) bool {
	if node == nil {
		return false
	}

	if isLeaf(node) && curSum+node.Val == wantSum {
		return true
	}
	return doHasPathSum(node.Left, curSum+node.Val, wantSum) ||
		doHasPathSum(node.Right, curSum+node.Val, wantSum)
}

func isLeaf(node *TreeNode) bool {
	if node == nil {
		return false
	}
	return node.Left == nil && node.Right == nil
}
