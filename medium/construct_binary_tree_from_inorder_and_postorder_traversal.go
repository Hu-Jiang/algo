package medium

// Given inorder and postorder traversal of a tree, construct the binary tree.
//
// Note:
// You may assume that duplicates do not exist in the tree.
//
// For example, given
// inorder = [9,3,15,20,7]
// postorder = [9,15,7,20,3]
// Return the following binary tree:
//     3
//    / \
//   9  20
//     /  \
//    15   7

func buildTrees(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	pivotVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: pivotVal}
	pivotIdx := 0
	for pivotIdx < len(inorder) {
		if inorder[pivotIdx] == pivotVal {
			break
		}
		pivotIdx++
	}
	root.Left = buildTrees(inorder[:pivotIdx], postorder[:pivotIdx])
	root.Right = buildTrees(inorder[pivotIdx+1:], postorder[pivotIdx:len(postorder)-1])

	return root
}
