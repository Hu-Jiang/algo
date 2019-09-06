package medium

import "math"

// Given a binary tree, determine if it is a valid binary search tree (BST).
//
// Assume a BST is defined as follows:
//
// The left subtree of a node contains only nodes with keys less than the node's key.
// The right subtree of a node contains only nodes with keys greater than the node's key.
// Both the left and right subtrees must also be binary search trees.
//
// Example 1:
//     2
//    / \
//   1   3
// Input: [2,1,3]
// Output: true
//
// Example 2:
//     5
//    / \
//   1   4
//      / \
//     3   6
// Input: [5,1,4,null,null,3,6]
// Output: false
// Explanation: The root node's value is 5 but its right child's value is 4.

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, low, high int64) bool {
	if root == nil {
		return true
	}

	if int64(root.Val) <= low || int64(root.Val) >= high {
		return false
	}

	return helper(root.Left, low, int64(root.Val)) && helper(root.Right, int64(root.Val), high)
}

// The official soloution:
//
// Approach 1: Recursion
//
// class Solution {
// 	public boolean helper(TreeNode node, Integer lower, Integer upper) {
// 	  if (node == null) return true;
//
// 	  int val = node.val;
// 	  if (lower != null && val <= lower) return false;
// 	  if (upper != null && val >= upper) return false;
//
// 	  if (! helper(node.right, val, upper)) return false;
// 	  if (! helper(node.left, lower, val)) return false;
// 	  return true;
// 	}
//
// 	public boolean isValidBST(TreeNode root) {
// 	  return helper(root, null, null);
// 	}
// }
//
// Approach 2: Iteration
//
// class Solution {
// 	LinkedList<TreeNode> stack = new LinkedList();
// 	LinkedList<Integer> uppers = new LinkedList(),
// 			lowers = new LinkedList();
//
// 	public void update(TreeNode root, Integer lower, Integer upper) {
// 	  stack.add(root);
// 	  lowers.add(lower);
// 	  uppers.add(upper);
// 	}
//
// 	public boolean isValidBST(TreeNode root) {
// 	  Integer lower = null, upper = null, val;
// 	  update(root, lower, upper);
//
// 	  while (!stack.isEmpty()) {
// 		root = stack.poll();
// 		lower = lowers.poll();
// 		upper = uppers.poll();
//
// 		if (root == null) continue;
// 		val = root.val;
// 		if (lower != null && val <= lower) return false;
// 		if (upper != null && val >= upper) return false;
// 		update(root.right, val, upper);
// 		update(root.left, lower, val);
// 	  }
// 	  return true;
// 	}
// }
//
// Approach 3: Inorder traversal
//
// class Solution {
// 	public boolean isValidBST(TreeNode root) {
// 	  Stack<TreeNode> stack = new Stack();
// 	  double inorder = - Double.MAX_VALUE;
//
// 	  while (!stack.isEmpty() || root != null) {
// 		while (root != null) {
// 		  stack.push(root);
// 		  root = root.left;
// 		}
// 		root = stack.pop();
// 		// If next element in inorder traversal
// 		// is smaller than the previous one
// 		// that's not BST.
// 		if (root.val <= inorder) return false;
// 		inorder = root.val;
// 		root = root.right;
// 	  }
// 	  return true;
// 	}
// }
