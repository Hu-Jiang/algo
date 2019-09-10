package easy

// Given a binary tree, return the bottom-up level order
// traversal of its nodes' values. (ie, from left to right,
// level by level from leaf to root).
//
// For example:
// Given binary tree [3,9,20,null,null,15,7],
//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its bottom-up level order traversal as:
// [
//   [15,7],
//   [9,20],
//   [3]
// ]

func levelOrderBottom(root *TreeNode) [][]int {
	res := levelOrderUp(root)
	reverseSlice(res)
	return res
}

func levelOrderUp(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		l := len(q)
		var tmp []int
		for i := 0; i < l; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		res = append(res, tmp)
	}
	return res
}

func reverseSlice(s [][]int) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

// Trick:
// If using list, we can insert to list head, need not reverse item like using array.
