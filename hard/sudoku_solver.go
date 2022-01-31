package hard

// Write a program to solve a Sudoku puzzle by filling the empty cells.
//
// A sudoku solution must satisfy all of the following rules:
// Each of the digits 1-9 must occur exactly once in each row.
// Each of the digits 1-9 must occur exactly once in each column.
// Each of the the digits 1-9 must occur exactly once in each of the 9 3x3
// sub-boxes of the grid.
//
// Empty cells are indicated by the character '.'.
//
// Image location: [/image/hard/sudoku_solver.svg.png]
// A sudoku puzzle...
//
// Image location: [/image/hard/sudoku_solver_solution.svg.png]
// ...and its solution numbers marked in red.
//
// Note:
// The given board contain only digits 1-9 and the character '.'.
// You may assume that the given Sudoku puzzle will have a single unique solution.
// The given board size is always 9x9.

func solveSudoku(board [][]byte) {
	doSolveSudoku(board)
}

func doSolveSudoku(board [][]byte) bool {
	candidates := "123456789"
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for k := 0; k < 9; k++ {
					board[i][j] = candidates[k]
					if isValidSudoku(board) && doSolveSudoku(board) {
						return true
					}
				}
				board[i][j] = '.'
				return false
			}
		}
	}

	return isValidSudoku(board)
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		hasNum := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if hasNum[board[i][j]] {
				return false
			} else {
				hasNum[board[i][j]] = true
			}
		}
	}

	for j := 0; j < 9; j++ {
		hasNum := make(map[byte]bool)
		for i := 0; i < 9; i++ {
			if board[i][j] == '.' {
				continue
			}
			if hasNum[board[i][j]] {
				return false
			} else {
				hasNum[board[i][j]] = true
			}
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if !isValidThree(board, i, j) {
				return false
			}
		}
	}

	return true
}

func isValidThree(board [][]byte, x, y int) bool {
	hasNum := make(map[byte]bool)
	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			if board[i][j] == '.' {
				continue
			}
			if hasNum[board[i][j]] {
				return false
			} else {
				hasNum[board[i][j]] = true
			}
		}
	}

	return true
}
