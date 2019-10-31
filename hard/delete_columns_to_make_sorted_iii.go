package hard

// We are given an array A of N lowercase letter strings, all of the same length.
//
// Now, we may choose any set of deletion indices, and for each string, we delete
// all the characters in those indices.
//
// For example, if we have an array A = ["babca","bbazb"] and deletion indices {0, 1, 4},
// then the final array after deletions is ["bc","az"].
//
// Suppose we chose a set of deletion indices D such that after deletions, the final
// array has every element (row) in lexicographic order.
//
// For clarity, A[0] is in lexicographic order (ie. A[0][0] <= A[0][1] <= ... <=
// A[0][A[0].length - 1]), A[1] is in lexicographic order (ie. A[1][0] <= A[1][1]
// <= ... <= A[1][A[1].length - 1]), and so on.
//
// Return the minimum possible value of D.length.
//
// Example 1:
// Input: ["babca","bbazb"]
// Output: 3
// Explanation: After deleting columns 0, 1, and 4, the final array is A = ["bc", "az"].
// Both these rows are individually in lexicographic order (ie. A[0][0] <= A[0][1]
// and A[1][0] <= A[1][1]).
// Note that A[0] > A[1] - the array A isn't necessarily in lexicographic order.
//
// Example 2:
// Input: ["edcba"]
// Output: 4
// Explanation: If we delete less than 4 columns, the only row won't be lexicographically sorted.
//
// Example 3:
// Input: ["ghi","def","abc"]
// Output: 0
// Explanation: All rows are already lexicographically sorted.
//
// Note:
// 1 <= A.length <= 100
// 1 <= A[i].length <= 100

func minDeletionSize(A []string) int {
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	dp := make([]int, len(A[0]))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(dp); i++ {
		for j := i - 1; j >= 0; j-- {
			goodPos := true
			for k := 0; k < len(A) && goodPos; k++ {
				if A[k][j] > A[k][i] {
					goodPos = false
				}
			}
			if goodPos {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	maxSorted := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] > maxSorted {
			maxSorted = dp[i]
		}
	}

	return len(A[0]) - maxSorted
}

// The official solution:
//
// Approach 1: Dynamic Programming
//
// class Solution {
//     public int minDeletionSize(String[] A) {
//         int W = A[0].length();
//         int[] dp = new int[W];
//         Arrays.fill(dp, 1);
//         for (int i = W-2; i >= 0; --i)
//             search: for (int j = i+1; j < W; ++j) {
//                 for (String row: A)
//                     if (row.charAt(i) > row.charAt(j))
//                         continue search;
//
//                 dp[i] = Math.max(dp[i], 1 + dp[j]);
//             }
//
//         int kept = 0;
//         for (int x: dp)
//             kept = Math.max(kept, x);
//         return W - kept;
//     }
// }
//
// Another version from discuss in leetcode:
//
// public int minDeletionSize(String[] A) {
// 	int m = A.length, n = A[0].length(), res = n - 1, k;
// 	int[] dp = new int[n];
// 	Arrays.fill(dp, 1);
// 	for (int j = 0; j < n; ++j) {
// 		for (int i = 0; i < j; ++i) {
// 			for (k = 0; k < m; ++k) {
// 				if (A[k].charAt(i) > A[k].charAt(j)) break;
// 			}
// 			if (k == m && dp[i] + 1 > dp[j])
// 				dp[j] = dp[i] + 1;
// 		}
// 		res = Math.min(res, n - dp[j]);
// 	}
// 	return res;
// }
