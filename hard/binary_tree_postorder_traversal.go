package hard

// Given a binary tree, return the postorder traversal of its nodes' values.
//
// Example:
// Input: [1,null,2,3]
//    1
//     \
//      2
//     /
//    3
// Output: [3,2,1]
//
// Follow up: Recursive solution is trivial, could you do it iteratively?

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	type trackTreeNode struct {
		root    *TreeNode
		visited bool
	}

	var (
		stack []*trackTreeNode
		res   []int
	)
	stack = append(stack, &trackTreeNode{root: root})
	for len(stack) != 0 {
		node := stack[len(stack)-1] // peek
		if node.visited {
			stack = stack[:len(stack)-1] // pop
			res = append(res, node.root.Val)
			continue
		}
		if node.root.Right != nil {
			stack = append(stack, &trackTreeNode{root: node.root.Right}) // push right child
		}
		if node.root.Left != nil {
			stack = append(stack, &trackTreeNode{root: node.root.Left}) // push left child
		}
		node.visited = true
	}

	return res
}
