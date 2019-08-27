package hard

// Given a 2D binary matrix filled with 0's and 1's, find the largest
// rectangle containing only 1's and return its area.
//
// Example:
// Input:
// [
//   ["1","0","1","0","0"],
//   ["1","0","1","1","1"],
//   ["1","1","1","1","1"],
//   ["1","0","0","1","0"]
// ]
// Output: 6

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	matrixHigh := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		matrixHigh[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix[0]); i++ {
		matrixHigh[0][i] = int(matrix[0][i] - '0')
	}

	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				matrixHigh[i][j] = 0
			} else {
				matrixHigh[i][j] = matrixHigh[i-1][j] + 1
			}
		}
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	maxArea := 0
	for i := 0; i < len(matrixHigh); i++ {
		for j := 0; j < len(matrixHigh[0]); j++ {
			if matrixHigh[i][j] == 0 {
				continue
			}
			width := 1
			for k := j + 1; k < len(matrixHigh[0]); k++ {
				if matrixHigh[i][k] < matrixHigh[i][j] {
					break
				}
				width++
			}
			for k := j - 1; k >= 0; k-- {
				if matrixHigh[i][k] < matrixHigh[i][j] {
					break
				}
				width++
			}
			maxArea = max(maxArea, matrixHigh[i][j]*width)
		}
	}

	return maxArea
}
