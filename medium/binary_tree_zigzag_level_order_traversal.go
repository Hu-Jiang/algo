package medium

// Given a binary tree, return the zigzag level order traversal of
// its nodes' values. (ie, from left to right, then right to left
// for the next level and alternate between).
//
// For example:
// Given binary tree [3,9,20,null,null,15,7],
//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its zigzag level order traversal as:
// [
//   [3],
//   [20,9],
//   [15,7]
// ]

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	arr := levelOrderArr(root)
	for i := 0; i < len(arr); i++ {
		if i%2 != 0 {
			revarsalArr(arr[i])
		}
	}
	return arr
}

func levelOrderArr(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ans [][]int
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
		ans = append(ans, tmp)
	}

	return ans
}

func revarsalArr(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}
