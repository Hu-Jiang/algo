package medium

// Given an array A of integers, a ramp is a tuple (i, j) for which
// i < j and A[i] <= A[j].  The width of such a ramp is j - i.
//
// Find the maximum width of a ramp in A.  If one doesn't exist, return 0.
//
// Example 1:
// Input: [6,0,8,2,1,5]
// Output: 4
// Explanation:
// The maximum width ramp is achieved at (i, j) = (1, 5): A[1] = 0 and A[5] = 5.
//
// Example 2:
// Input: [9,8,1,0,1,9,4,0,4,1]
// Output: 7
// Explanation:
// The maximum width ramp is achieved at (i, j) = (2, 9): A[2] = 1 and A[9] = 1.
//
// Note:
// 2 <= A.length <= 50000
// 0 <= A[i] <= 50000

func maxWidthRamp(A []int) int {
	maxLen := 0
	for i := 0; i < len(A)-1; i++ {
		if len(A)-1-i < maxLen {
			break
		}
		for j := len(A) - 1; j > 0; j-- {
			if A[i] <= A[j] && j-i > maxLen {
				maxLen = j - i
				break
			}
		}
	}
	return maxLen
}
