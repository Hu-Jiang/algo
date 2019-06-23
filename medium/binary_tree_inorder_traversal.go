package medium

// Given a binary tree, return the inorder traversal of its nodes' values.
//
// Example:
// Given binary tree [1,null,2,3]
//    1
//     \
//      2
//     /
//    3
// Output: [1,3,2]
//
// Follow up: Recursive solution is trivial, could you do it iteratively?

func inorderTraversal(root *TreeNode) []int {
	var a []int
	traversal(root, &a)
	return a
}

func traversal(root *TreeNode, a *[]int) {
	if root == nil {
		return
	}
	traversal(root.Left, a)
	*a = append(*a, root.Val)
	traversal(root.Right, a)
}

// The official soloution:
//
// Approach 1: Recursive Approach
//
// class Solution {
//     public List < Integer > inorderTraversal(TreeNode root) {
//         List < Integer > res = new ArrayList < > ();
//         helper(root, res);
//         return res;
//     }
//
//     public void helper(TreeNode root, List < Integer > res) {
//         if (root != null) {
//             if (root.left != null) {
//                 helper(root.left, res);
//             }
//             res.add(root.val);
//             if (root.right != null) {
//                 helper(root.right, res);
//             }
//         }
//     }
// }
//
// Approach 2: Iterating method using Stack
//
// public class Solution {
//     public List < Integer > inorderTraversal(TreeNode root) {
//         List < Integer > res = new ArrayList < > ();
//         Stack < TreeNode > stack = new Stack < > ();
//         TreeNode curr = root;
//         while (curr != null || !stack.isEmpty()) {
//             while (curr != null) {
//                 stack.push(curr);
//                 curr = curr.left;
//             }
//             curr = stack.pop();
//             res.add(curr.val);
//             curr = curr.right;
//         }
//         return res;
//     }
// }
//
// Approach 3: Morris Traversal
//
// class Solution {
//     public List < Integer > inorderTraversal(TreeNode root) {
//         List < Integer > res = new ArrayList < > ();
//         TreeNode curr = root;
//         TreeNode pre;
//         while (curr != null) {
//             if (curr.left == null) {
//                 res.add(curr.val);
//                 curr = curr.right; // move to next right node
//             } else { // has a left subtree
//                 pre = curr.left;
//                 while (pre.right != null) { // find rightmost
//                     pre = pre.right;
//                 }
//                 pre.right = curr; // put cur after the pre node
//                 TreeNode temp = curr; // store cur node
//                 curr = curr.left; // move cur to the top of the new tree
//                 temp.left = null; // original cur left be null, avoid infinite loops
//             }
//         }
//         return res;
//     }
// }
