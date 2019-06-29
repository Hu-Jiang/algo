package easy

import "reflect"

// Given two binary trees, write a function to check if they
// are the same or not.
//
// Two binary trees are considered the same if they are
// structurally identical and the nodes have the same value.
//
// Example 1:
// Input:     1         1
//           / \       / \
//          2   3     2   3
//
//         [1,2,3],   [1,2,3]
//
// Output: true
//
// Example 2:
// Input:     1         1
//           /           \
//          2             2
//
//         [1,2],     [1,null,2]
//
// Output: false
//
// Example 3:
//
// Input:     1         1
//           / \       / \
//          2   1     1   2
//
//         [1,2,1],   [1,1,2]
//
// Output: false

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// another version of mine:
func isSameTree_v1(p *TreeNode, q *TreeNode) bool {
	return reflect.DeepEqual(PreorderArr_v1(p), PreorderArr_v1(q))
}

func PreorderArr_v1(t *TreeNode) []int {
	if t == nil {
		return nil
	}

	var res []int
	preorderTraversal_v1(t, &res)
	return res
}

func preorderTraversal_v1(t *TreeNode, res *[]int) {
	if t == nil {
		*res = append(*res, -1)
		return
	}
	*res = append(*res, t.Val)

	preorderTraversal_v1(t.Left, res)
	preorderTraversal_v1(t.Right, res)
}

// The official soloution:
//
// Approach 1: Recursion
//
// class Solution {
// 	public boolean isSameTree(TreeNode p, TreeNode q) {
// 		// p and q are both null
// 		if (p == null && q == null) return true;
// 		// one of p and q is null
// 		if (q == null || p == null) return false;
// 		if (p.val != q.val) return false;
// 		return isSameTree(p.right, q.right) &&
// 				isSameTree(p.left, q.left);
// 	}
// }
//
// Approach 2: Iteration
//
// class Solution {
// 	public boolean check(TreeNode p, TreeNode q) {
// 	  // p and q are null
// 	  if (p == null && q == null) return true;
// 	  // one of p and q is null
// 	  if (q == null || p == null) return false;
// 	  if (p.val != q.val) return false;
// 	  return true;
// 	}
//
// 	public boolean isSameTree(TreeNode p, TreeNode q) {
// 	  if (p == null && q == null) return true;
// 	  if (!check(p, q)) return false;
//
// 	  // init deques
// 	  ArrayDeque<TreeNode> deqP = new ArrayDeque<TreeNode>();
// 	  ArrayDeque<TreeNode> deqQ = new ArrayDeque<TreeNode>();
// 	  deqP.addLast(p);
// 	  deqQ.addLast(q);
//
// 	  while (!deqP.isEmpty()) {
// 		p = deqP.removeFirst();
// 		q = deqQ.removeFirst();
//
// 		if (!check(p, q)) return false;
// 		if (p != null) {
// 		  // in Java nulls are not allowed in Deque
// 		  if (!check(p.left, q.left)) return false;
// 		  if (p.left != null) {
// 			deqP.addLast(p.left);
// 			deqQ.addLast(q.left);
// 		  }
// 		  if (!check(p.right, q.right)) return false;
// 		  if (p.right != null) {
// 			deqP.addLast(p.right);
// 			deqQ.addLast(q.right);
// 		  }
// 		}
// 	  }
// 	  return true;
// 	}
// }
