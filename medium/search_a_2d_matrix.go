package medium

// Write an efficient algorithm that searches for a value in an m x n matrix.
// This matrix has the following properties:
//
// Integers in each row are sorted from left to right.
// The first integer of each row is greater than the last integer of the previous row.
//
// Example 1:
// Input:
// matrix = [
//   [1,   3,  5,  7],
//   [10, 11, 16, 20],
//   [23, 30, 34, 50]
// ]
// target = 3
// Output: true
//
// Example 2:
// Input:
// matrix = [
//   [1,   3,  5,  7],
//   [10, 11, 16, 20],
//   [23, 30, 34, 50]
// ]
// target = 13
// Output: false

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	row := 0
	for row < len(matrix) {
		if len(matrix[row]) == 0 {
			return false
		}
		if matrix[row][0] > target {
			break
		}
		row++
	}
	if row != 0 {
		row--
	}

	l, r := 0, len(matrix[row])-1
	for l <= r {
		mid := l + (r-l)/2
		if matrix[row][mid] > target {
			r = mid - 1
		} else if matrix[row][mid] < target {
			l = mid + 1
		} else {
			return true
		}
	}
	return false
}

// Trick: use twice binary search is better
