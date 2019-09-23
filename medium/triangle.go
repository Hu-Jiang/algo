package medium

// Given a triangle, find the minimum path sum from top to bottom.
// Each step you may move to adjacent numbers on the row below.
//
// For example, given the following triangle
// [
//      [2],
//     [3,4],
//    [6,5,7],
//   [4,1,8,3]
// ]
// The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
//
// Note:
// Bonus point if you are able to do this using only O(n) extra space,
// where n is the total number of rows in the triangle.

func minimumTotal(triangle [][]int) int {
	if triangle == nil || triangle[0] == nil {
		return -1
	}
	triMinPathSum(triangle)
	return minArrVal(triangle[len(triangle)-1])
}

// triPathSum record triangle minimum path sum in place.
func triMinPathSum(triangle [][]int) {
	for i := 1; i < len(triangle); i++ {
		triangle[i][0] += triangle[i-1][0]
		for j := 1; j < i; j++ {
			triangle[i][j] += minNum(triangle[i-1][j-1], triangle[i-1][j])
		}
		triangle[i][i] += triangle[i-1][i-1]
	}
}

func minNum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minArrVal(arr []int) int {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

// Trick: If using bottom to up dp, we do not need minArrVal() opreation.
//
// int minimumTotal(vector<vector<int> > &triangle) {
//     int n = triangle.size();
//     vector<int> minlen(triangle.back());
//     for (int layer = n-2; layer >= 0; layer--) // For each layer
//     {
//         for (int i = 0; i <= layer; i++) // Check its every 'node'
//         {
//             // Find the lesser of its two children, and sum the current value in the triangle with it.
//             minlen[i] = min(minlen[i], minlen[i+1]) + triangle[layer][i];
//         }
//     }
//     return minlen[0];
// }
