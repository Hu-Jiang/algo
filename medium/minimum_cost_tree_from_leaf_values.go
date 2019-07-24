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
