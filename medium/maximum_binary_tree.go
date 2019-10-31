package medium

// Given an integer array with no duplicates. A maximum tree
// building on this array is defined as follow:
//
// 1. The root is the maximum number in the array.
// 2. The left subtree is the maximum tree constructed from
// left part subarray divided by the maximum number.
// 3. The right subtree is the maximum tree constructed from
// right part subarray divided by the maximum number.
//
// Construct the maximum tree by the given array and output
// the root node of this tree.
//
// Example 1:
// Input: [3,2,1,6,0,5]
// Output: return the tree root node representing the following tree:
//       6
//     /   \
//    3     5
//     \    /
//      2  0
//        \
//         1
// Note:
// The size of the given array will be in the range [1,1000].

func constructMaximumBinaryTree(nums []int) *TreeNode {
	len := len(nums)
	l, r := 0, len-1
	return doConstructor(nums, l, r)
}

func doConstructor(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	i := GetMaxValueIndex(nums, l, r)
	node := &TreeNode{Val: nums[i]}
	node.Left = doConstructor(nums, l, i-1)
	node.Right = doConstructor(nums, i+1, r)
	return node
}

func GetMaxValueIndex(nums []int, l, r int) int { // [l..r]
	mvi := l
	for i := l; i <= r; i++ {
		if nums[i] > nums[mvi] {
			mvi = i
		}
	}
	return mvi
}

// my idea:
// 1. for l..r find max value
// 2. use max value be the root
// 3. l..max_index-1 be left root, max_index+1..r be right root

// The official solution:
//
// Approach 1: Recursive Solution
//
// public class Solution {
//     public TreeNode constructMaximumBinaryTree(int[] nums) {
//         return construct(nums, 0, nums.length);
//     }
//     public TreeNode construct(int[] nums, int l, int r) {
//         if (l == r)
//             return null;
//         int max_i = max(nums, l, r);
//         TreeNode root = new TreeNode(nums[max_i]);
//         root.left = construct(nums, l, max_i);
//         root.right = construct(nums, max_i + 1, r);
//         return root;
//     }
//     public int max(int[] nums, int l, int r) {
//         int max_i = l;
//         for (int i = l; i < r; i++) {
//             if (nums[max_i] < nums[i])
//                 max_i = i;
//         }
//         return max_i;
//     }
// }
