package hard

import "math"

// Two elements of a binary search tree (BST) are swapped by mistake.
//
// Recover the tree without changing its structure.
//
// Example 1:
// Input: [1,3,null,null,2]
//    1
//   /
//  3
//   \
//    2
// Output: [3,1,null,null,2]
//    3
//   /
//  1
//   \
//    2
//
// Example 2:
// Input: [3,1,4,null,null,2]
//   3
//  / \
// 1   4
//    /
//   2
// Output: [2,1,4,null,null,3]
//   2
//  / \
// 1   4
//    /
//   3
//
// Follow up:
// A solution using O(n) space is pretty straight forward.
// Could you devise a constant space solution?

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}
	wn := new(wrongNode)
	wn.inorder(root)
	wn.findWrongNode()
	wn.changeWrongNode()
}

type wrongNode struct {
	first  *TreeNode
	second *TreeNode

	inorderArr []*TreeNode
}

func (w *wrongNode) inorder(root *TreeNode) {
	if root == nil {
		return
	}

	w.inorder(root.Left)
	w.inorderArr = append(w.inorderArr, root)
	w.inorder(root.Right)
}

func (w *wrongNode) findWrongNode() {
	arr := append([]*TreeNode{}, &TreeNode{Val: math.MinInt32})
	arr = append(arr, w.inorderArr...)
	arr = append(arr, &TreeNode{Val: math.MaxInt32})
	for i := 1; i < len(arr)-1; i++ {
		if arr[i-1].Val <= arr[i].Val && arr[i].Val > arr[i+1].Val {
			w.first = arr[i]
			break
		}
	}

	for i := len(arr) - 2; i > 0; i-- {
		if arr[i-1].Val > arr[i].Val && arr[i].Val <= arr[i+1].Val {
			w.second = arr[i]
			break
		}
	}
}

func (w *wrongNode) changeWrongNode() {
	w.first.Val, w.second.Val = w.second.Val, w.first.Val
}
