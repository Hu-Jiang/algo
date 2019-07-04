package medium

// Given preorder and inorder traversal of a tree, construct the binary tree.
//
// Note:
// You may assume that duplicates do not exist in the tree.
//
// For example, given
// preorder = [3,9,20,15,7]
// inorder = [9,3,15,20,7]
// Return the following binary tree:
//     3
//    / \
//   9  20
//     /  \
//    15   7

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	pivotVal := preorder[0]
	root := &TreeNode{Val: pivotVal}
	pivotIdx := 0
	for pivotIdx < len(inorder) {
		if inorder[pivotIdx] == pivotVal {
			break
		}
		pivotIdx++
	}
	root.Left = buildTree(preorder[1:pivotIdx+1], inorder[:pivotIdx])
	root.Right = buildTree(preorder[pivotIdx+1:], inorder[pivotIdx+1:])

	return root
}
