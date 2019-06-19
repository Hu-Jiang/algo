package medium

// Given a matrix of m x n elements (m rows, n columns), return all
// elements of the matrix in spiral order.
//
// Example 1:
// Input:
// [
//  [ 1, 2, 3 ],
//  [ 4, 5, 6 ],
//  [ 7, 8, 9 ]
// ]
// Output: [1,2,3,6,9,8,7,4,5]
//
// Example 2:
// Input:
// [
//   [1, 2, 3, 4],
//   [5, 6, 7, 8],
//   [9,10,11,12]
// ]
// Output: [1,2,3,4,8,12,11,10,9,5,6,7]

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	rows, cols := len(matrix), len(matrix[0])
	total := rows * cols
	var cnt int
	var ans []int

	for round := 0; cnt < total; round += 1 {
		for i := round; cnt < total && i < cols-round; i++ {
			ans = append(ans, matrix[round][i])
			cnt++
		}

		for i := round + 1; cnt < total && i < rows-round; i++ {
			ans = append(ans, matrix[i][cols-1-round])
			cnt++
		}

		for i := cols - 2 - round; cnt < total && i >= round; i-- {
			ans = append(ans, matrix[rows-1-round][i])
			cnt++
		}

		for i := rows - 2 - round; cnt < total && i > round; i-- {
			ans = append(ans, matrix[i][round])
			cnt++
		}
	}

	return ans
}

// The official soloution:
//
// Approach 1: Simulation
//
// class Solution {
//     public List<Integer> spiralOrder(int[][] matrix) {
//         List ans = new ArrayList();
//         if (matrix.length == 0) return ans;
//         int R = matrix.length, C = matrix[0].length;
//         boolean[][] seen = new boolean[R][C];
//         int[] dr = {0, 1, 0, -1};
//         int[] dc = {1, 0, -1, 0};
//         int r = 0, c = 0, di = 0;
//         for (int i = 0; i < R * C; i++) {
//             ans.add(matrix[r][c]);
//             seen[r][c] = true;
//             int cr = r + dr[di];
//             int cc = c + dc[di];
//             if (0 <= cr && cr < R && 0 <= cc && cc < C && !seen[cr][cc]){
//                 r = cr;
//                 c = cc;
//             } else {
//                 di = (di + 1) % 4;
//                 r += dr[di];
//                 c += dc[di];
//             }
//         }
//         return ans;
//     }
// }
//
// Approach 2: Layer-by-Layer
//
// class Solution {
//     public List < Integer > spiralOrder(int[][] matrix) {
//         List ans = new ArrayList();
//         if (matrix.length == 0)
//             return ans;
//         int r1 = 0, r2 = matrix.length - 1;
//         int c1 = 0, c2 = matrix[0].length - 1;
//         while (r1 <= r2 && c1 <= c2) {
//             for (int c = c1; c <= c2; c++) ans.add(matrix[r1][c]);
//             for (int r = r1 + 1; r <= r2; r++) ans.add(matrix[r][c2]);
//             if (r1 < r2 && c1 < c2) {
//                 for (int c = c2 - 1; c > c1; c--) ans.add(matrix[r2][c]);
//                 for (int r = r2; r > r1; r--) ans.add(matrix[r][c1]);
//             }
//             r1++;
//             r2--;
//             c1++;
//             c2--;
//         }
//         return ans;
//     }
// }
