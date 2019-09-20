package easy

// Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.
//
// In Pascal's triangle, each number is the sum of the two numbers directly above it.
//
// Example:
// Input: 5
// Output:
// [
//      [1],
//     [1,1],
//    [1,2,1],
//   [1,3,3,1],
//  [1,4,6,4,1]
// ]

func generate(numRows int) [][]int {
	if numRows < 0 {
		return nil
	}

	var res [][]int
	for i := 0; i < numRows; i++ {
		curRow := make([]int, i+1)
		curRow[0] = 1
		for j := 1; j < i; j++ {
			curRow[j] = res[i-1][j-1] + res[i-1][j]
		}
		curRow[i] = 1

		res = append(res, curRow)
	}
	return res
}

// The official soloution:
//
// Approach 1: Dynamic Programming
// class Solution:
//     def generate(self, num_rows):
//         triangle = []
//
//         for row_num in range(num_rows):
//             # The first and last row elements are always 1.
//             row = [None for _ in range(row_num+1)]
//             row[0], row[-1] = 1, 1
//
//             # Each triangle element is equal to the sum of the elements
//             # above-and-to-the-left and above-and-to-the-right.
//             for j in range(1, len(row)-1):
//                 row[j] = triangle[row_num-1][j-1] + triangle[row_num-1][j]
//
//             triangle.append(row)
//
//         return triangle
