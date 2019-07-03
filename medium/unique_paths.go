package medium

// A robot is located at the top-left corner of a m x n
// grid (marked 'Start' in the diagram below).
//
// The robot can only move either down or right at any point
// in time. The robot is trying to reach the bottom-right corner
// of the grid (marked 'Finish' in the diagram below).
//
// How many possible unique paths are there?
//
// Image location: [/image/medium/unique_paths.png]
// Above is a 7 x 3 grid. How many possible unique paths are there?
//
// Note: m and n will be at most 100.
//
// Example 1:
// Input: m = 3, n = 2
// Output: 3
// Explanation:
// From the top-left corner, there are a total of 3 ways to reach
// the bottom-right corner:
// 1. Right -> Right -> Down
// 2. Right -> Down -> Right
// 3. Down -> Right -> Right
//
// Example 2:
// Input: m = 7, n = 3
// Output: 28

func uniquePaths(m int, n int) int {
	res := make([][]int, m+1)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, n+1)
	}

	for i := 1; i < len(res); i++ {
		for j := 1; j < len(res[0]); j++ {
			res[i][j] = maxStep(res[i-1][j]+res[i][j-1], 1)
		}
	}
	return res[m][n]
}

func maxStep(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// No extra wall:
//
// func uniquePaths(m int, n int) int {
// 	result := make([][]int, m)
// 	for i := 0; i < m; i++ {
// 		result[i] = make([]int, n)
// 	}
//
// 	for i := 0; i < m; i++ {
// 		result[i][0] = 1
// 	}
//
// 	for j := 1; j < n; j++ {
// 		result[0][j] = 1
// 	}
//
// 	for i := 1; i < m; i++ {
// 		for j := 1; j < n; j++ {
// 			result[i][j] = result[i-1][j] + result[i][j-1]
// 		}
// 	}
//
// 	return result[m-1][n-1]
// }
