package medium

import (
	"math"
)

// Given a m x n grid filled with non-negative numbers, find a path
// from top left to bottom right which minimizes the sum of all
// numbers along its path.
//
// Note: You can only move either down or right at any point in time.
//
// Example:
// Input:
// [
//   [1,3,1],
//   [1,5,1],
//   [4,2,1]
// ]
// Output: 7
// Explanation: Because the path 1→3→1→1→1 minimizes the sum.

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	dp := make([][]int, len(grid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid[0]))
	}

	var left, up int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
				continue
			}
			if j-1 < 0 {
				left = math.MaxInt32
			} else {
				left = dp[i][j-1]
			}
			if i-1 < 0 {
				up = math.MaxInt32
			} else {
				up = dp[i-1][j]
			}
			dp[i][j] = minInt(left, up) + grid[i][j]
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}

func minInt(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// Trik: using "wall" to surround grid to simplify code
//
// My first version, which got "Time Limit Exceeded" in leetcode:
//
// func minPathSum(grid [][]int) int {
// 	if len(grid) == 0 {
// 		return 0
// 	}
//
// 	ans := math.MaxInt32
// 	backtracking(grid, 0, 0, 0, &ans)
// 	return ans
// }
//
// func backtracking(grid [][]int, row, col int, prevSum int, minPathSum *int) {
// 	totalRow := len(grid)
// 	totalCol := len(grid[0])
//
// 	if row == totalRow-1 && col == totalCol-1 {
// 		if prevSum+grid[row][col] < *minPathSum {
// 			*minPathSum = prevSum + grid[row][col]
// 		}
// 		return
// 	}
//
// 	if row != totalRow-1 {
// 		backtracking(grid, row+1, col, prevSum+grid[row][col], minPathSum)
// 	}
//
// 	if col != totalCol-1 {
// 		backtracking(grid, row, col+1, prevSum+grid[row][col], minPathSum)
// 	}
// }
