package medium

// Given a binary tree, return the preorder traversal of its nodes' values.
//
// Example:
// Input: [1,null,2,3]
//    1
//     \
//      2
//     /
//    3
//
// Output: [1,2,3]
//
// Follow up: Recursive solution is trivial, could you do it iteratively?

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var (
		stack []*TreeNode
		res   []int
	)
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]  // peek
		stack = stack[:len(stack)-1] // pop
		res = append(res, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right) // push right child
		}
		if node.Left != nil {
			stack = append(stack, node.Left) // push left child
		}
	}

	return res
}
