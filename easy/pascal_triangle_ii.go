package easy

// Given a non-negative index k where k ≤ 33, return the k-th
// index row of the Pascal's triangle.
//
// Note that the row index starts from 0.
//
// In Pascal's triangle, each number is the sum of the two
// numbers directly above it.
//      [1],
//     [1,1],
//    [1,2,1],
//   [1,3,3,1],
//  [1,4,6,4,1]
//
// Example:
// Input: 3
// Output: [1,3,3,1]
//
// Follow up:
// Could you optimize your algorithm to use only O(k) extra space?

func getRow(rowIndex int) []int {
	if rowIndex < 0 {
		return nil
	}

	res := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		res[i] = 1
		for j := i - 1; j > 0; j-- {
			res[j] += res[j-1]
		}
	}

	return res
}
