package hard

// The n-queens puzzle is the problem of placing n queens on an n×n chessboard
// such that no two queens attack each other.
//
// Image location: [/image/hard/n-queens-ii.png]
//
// Given an integer n, return the number of distinct solutions to the n-queens puzzle.
//
// Example:
// Input: 4
// Output: 2
// Explanation: There are two distinct solutions to the 4-queens puzzle as shown below.
// [
//  [".Q..",  // Solution 1
//   "...Q",
//   "Q...",
//   "..Q."],
//  ["..Q.",  // Solution 2
//   "Q...",
//   "...Q",
//   ".Q.."]
// ]

func totalNQueens(n int) int {
	if n == 0 {
		return 0
	}

	board := newChessBoard(n + 1)
	ans := 0
	backtracking(board, 1, n, &ans)

	return ans
}

func backtracking(board [][]int, row, n int, res *int) {
	if row == n+1 {
		*res++
		return
	}

	for i := 1; i <= n; i++ {
		if canPlace(board, row, i, n) {
			board[row][i] = 1
			backtracking(board, row+1, n, res)
			board[row][i] = 0
		}
	}
}

func canPlace(board [][]int, row, col, n int) bool {
	for j := 1; j < row; j++ {
		for k := 1; k <= n; k++ {
			if board[j][k] == 1 {
				if row == j {
					return false
				}
				if col == k {
					return false
				}
				if absInt(row-j) == absInt(col-k) {
					return false
				}
			}
		}
	}
	return true
}

func absInt(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func newChessBoard(n int) [][]int {
	var board [][]int
	for i := 0; i < n; i++ {
		board = append(board, make([]int, n))
	}
	return board
}
