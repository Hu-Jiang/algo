package medium

// Given a 2D board and a word, find if the word exists in the grid.
//
// The word can be constructed from letters of sequentially adjacent
// cell, where "adjacent" cells are those horizontally or vertically
// neighboring. The same letter cell may not be used more than once.
//
// Example:
// board =
// [
//   ['A','B','C','E'],
//   ['S','F','C','S'],
//   ['A','D','E','E']
// ]
// Given word = "ABCCED", return true.
// Given word = "SEE", return true.
// Given word = "ABCB", return false.

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(word) == 0 {
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfsMatch(board, word, i, j) {
				return true
			}
		}
	}
	return false
}

func dfsMatch(board [][]byte, word string, i, j int) bool {
	if len(word) == 0 {
		return true
	}
	if board[i][j] != word[0] {
		return false
	}
	if len(word) == 1 {
		return true
	}
	tmp := board[i][j]
	board[i][j] = 0
	isMatched := (j+1 < len(board[0]) && dfsMatch(board, word[1:], i, j+1)) ||
		(j-1 >= 0 && dfsMatch(board, word[1:], i, j-1)) ||
		(i+1 < len(board) && dfsMatch(board, word[1:], i+1, j)) ||
		(i-1 >= 0 && dfsMatch(board, word[1:], i-1, j))
	board[i][j] = tmp

	return isMatched
}
