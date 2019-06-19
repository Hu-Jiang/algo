package medium

// Given a positive integer n, generate a square matrix filled
// with elements from 1 to n^2 in spiral order.
//
// Example:
// Input: 3
// Output:
// [
//  [ 1, 2, 3 ],
//  [ 8, 9, 4 ],
//  [ 7, 6, 5 ]
// ]

func generateMatrix(n int) [][]int {
	var matrix [][]int
	for i := 0; i < n; i++ {
		matrix = append(matrix, make([]int, n))
	}

	total := n * n
	var cnt int

	for round := 0; cnt < total; round += 1 {
		for i := round; cnt < total && i < n-round; i++ {
			cnt++
			matrix[round][i] = cnt
		}

		for i := round + 1; cnt < total && i < n-round; i++ {
			cnt++
			matrix[i][n-1-round] = cnt
		}

		for i := n - 2 - round; cnt < total && i >= round; i-- {
			cnt++
			matrix[n-1-round][i] = cnt
		}

		for i := n - 2 - round; cnt < total && i > round; i-- {
			cnt++
			matrix[i][round] = cnt
		}
	}

	return matrix
}
