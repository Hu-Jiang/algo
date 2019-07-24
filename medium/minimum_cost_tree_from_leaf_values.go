package medium

import "math"

// Given an array arr of positive integers, consider all
// binary trees such that:
//
// Each node has either 0 or 2 children;
// The values of arr correspond to the values of each leaf
// in an in-order traversal of the tree.  (Recall that a node
// is a leaf if and only if it has 0 children.)
// The value of each non-leaf node is equal to the product of
// the largest leaf value in its left and right subtree respectively.
// Among all possible binary trees considered, return the smallest
// possible sum of the values of each non-leaf node.  It is guaranteed
// this sum fits into a 32-bit integer.
//
// Example 1:
// Input: arr = [6,2,4]
// Output: 32
// Explanation:
// There are two possible trees.  The first has non-leaf node sum 36,
// and the second has non-leaf node sum 32.
//
//     24            24
//    /  \          /  \
//   12   4        6    8
//  /  \               / \
// 6    2             2   4
//
// Constraints:
//
// 2 <= arr.length <= 40
// 1 <= arr[i] <= 15
// It is guaranteed that the answer fits into a 32-bit signed
// integer (ie. it is less than 2^31).

func mctFromLeafValues(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}

	ans := 0
	for len(arr) > 1 {
		minVal := arr[0]
		minIdx := 0
		for i := 0; i < len(arr); i++ {
			if arr[i] < minVal {
				minVal = arr[i]
				minIdx = i
			}
		}
		left, right := math.MaxInt32, math.MaxInt32
		if minIdx != 0 {
			left = arr[minIdx-1]
		}
		if minIdx != len(arr)-1 {
			right = arr[minIdx+1]
		}
		ans += minVal * minLeaf(left, right)
		arr = append(arr[:minIdx], arr[minIdx+1:]...)
	}

	return ans
}

func minLeaf(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// DP solution from leetcode discuss:
// 递推公式dp[i][j]代表节点从i到j所能构成的非叶节点的最小和:
// 			dp[i][j] = min(dp[i][j],dp[i][k] + dp[k+1][j] + max[i][k]*max[k+1][j])
// 那么,最终我们要求的结果就是 dp[0][n-1]的值.枚举过程是先枚举区间长度l,再枚举起始点s,
// 根据起始点和长度算出区间终点s+l,最后枚举区间中的点k.
//
// class Solution:
//     def mctFromLeafValues(self, arr: List[int]) -> int:
//         n = len(arr)
//         # 初始值设为最大
//         dp = [[float('inf') for _ in range(n)] for _ in range(n)]
//         # 初始区间查询最大值设为0
//         maxval = [[0 for _ in range(n)] for _ in range(n)]
//         # 求区间[i, j]中最大元素
//         for i in range(n):
//             for j in range(i, n):
//                 maxval[i][j] = max(arr[i:j + 1])
//         # 叶子结点不参与计算
//         for i in range(n):
//             dp[i][i] = 0
//         # 枚举区间长度
//         for l in range(1, n):
//             # 枚举区间起始点
//             for s in range(n - l):
//                 # 枚举划分两棵子树
//                 for k in range(s, s + l):
//                     dp[s][s + l] = min(dp[s][s + l], dp[s][k] + dp[k + 1][s + l] + maxval[s][k] * maxval[k + 1][s + l])
//         return dp[0][n - 1]
