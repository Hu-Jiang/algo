package medium

// Given a binary tree, flatten it to a linked list in-place.
//
// For example, given the following tree:
//     1
//    / \
//   2   5
//  / \   \
// 3   4   6
//
// The flattened tree should look like:
// 1
//  \
//   2
//    \
//     3
//      \
//       4
//        \
//         5
//          \
//           6

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	cursor = new(TreeNode)
	dummy := cursor
	preorderFlatten(root)
	*root = *dummy.Right
}

var cursor *TreeNode

func preorderFlatten(root *TreeNode) {
	if root == nil {
		return
	}

	cursor.Right = &TreeNode{Val: root.Val}
	cursor = cursor.Right
	preorderFlatten(root.Left)
	preorderFlatten(root.Right)
}
