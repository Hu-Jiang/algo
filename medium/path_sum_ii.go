package medium

// Given a binary tree and a sum, find all root-to-leaf
// paths where each path's sum equals the given sum.
//
// Note: A leaf is a node with no children.
//
// Example:
// Given the below binary tree and sum = 22,
//       5
//      / \
//     4   8
//    /   / \
//   11  13  4
//  /  \    / \
// 7    2  5   1
// Return:
// [
//    [5,4,11,2],
//    [5,8,4,5]
// ]

func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	var path []int
	doPathSum(root, 0, sum, path, &res)
	return res
}

func doPathSum(root *TreeNode, curSum, wantSum int, path []int, res *[][]int) {
	if root == nil {
		return
	}

	curSum = curSum + root.Val
	path = append(path, root.Val)

	if isLeft(root) {
		if curSum == wantSum {
			*res = append(*res, append([]int{}, path...))
		}
		return
	}

	doPathSum(root.Left, curSum, wantSum, path, res)
	doPathSum(root.Right, curSum, wantSum, path, res)
}

func isLeft(node *TreeNode) bool {
	if node == nil {
		return false
	}
	return node.Left == nil && node.Right == nil
}
