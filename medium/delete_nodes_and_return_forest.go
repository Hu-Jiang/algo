package medium

// Given the root of a binary tree, each node in the tree has a distinct value.
//
// After deleting all nodes with a value in to_delete, we are left with a forest
// (a disjoint union of trees).
//
// Return the roots of the trees in the remaining forest.  You may return the result
// in any order.
//
// Example 1:
// Image location: [/image/medium/delete_nodes_and_return_forest.png]
// Input: root = [1,2,3,4,5,6,7], to_delete = [3,5]
// Output: [[1,2,null,4],[6],[7]]
//
// Constraints:
// The number of nodes in the given tree is at most 1000.
// Each node has a distinct value between 1 and 1000.
// to_delete.length <= 1000
// to_delete contains distinct values between 1 and 1000.

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	if root == nil {
		return nil
	}
	res := []*TreeNode{root}
	for j := 0; j < len(to_delete); j++ {
		for i := 0; i < len(res); i++ {
			if res[i].Val == -1 {
				continue
			}
			if parrent, current := traversalNode(res[i], to_delete[j]); current != nil {
				if current.Left != nil {
					res = append(res, current.Left)
				}
				if current.Right != nil {
					res = append(res, current.Right)
				}
				if parrent != nil {
					if parrent.Left != nil && parrent.Left.Val == current.Val {
						parrent.Left = nil
					} else {
						parrent.Right = nil
					}
				}
				current.Val = -1 // -1 flag delete
				break
			}
		}
	}

	for i := 0; i < len(res); i++ {
		if res[i].Val == -1 {
			res = append(res[:i], res[i+1:]...)
			i--
		}
	}

	return res
}

func traversalNode(root *TreeNode, v int) (parent, current *TreeNode) {
	if root == nil {
		return nil, nil
	}
	if root.Val == v {
		return nil, root
	}
	if p, c := traversalNode(root.Left, v); c != nil {
		if p == nil {
			p = root
		}
		return p, c
	}
	if p, c := traversalNode(root.Right, v); c != nil {
		if p == nil {
			p = root
		}
		return p, c
	}
	return nil, nil
}

// A recursion solution from discuss in leetcode of this problem:
//
// Intuition
// As I keep saying in my "courses",
// solve tree problem with recursion first.
//
// Explanation
// If a node is root (has no parent) and isn't deleted,
// when will we add it to the result.
//
// Complexity
// Time O(N)
// Space O(N)
//
// Java:
//
//     Set<Integer> to_delete_set;
//     List<TreeNode> res = new ArrayList<>();
//     public List<TreeNode> delNodes(TreeNode root, int[] to_delete) {
//         res = new ArrayList<>();
//         to_delete_set = new HashSet<>();
//         for (int i : to_delete)
//             to_delete_set.add(i);
//         helper(root, true);
//         return res;
//     }
//
//     private TreeNode helper(TreeNode node, boolean is_root) {
//         if (node == null) return null;
//         boolean deleted = to_delete_set.contains(node.val);
//         if (is_root && !deleted) res.add(node);
//         node.left = helper(node.left, deleted);
//         node.right =  helper(node.right, deleted);
//         return deleted ? null : node;
// 	}
