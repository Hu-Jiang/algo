package medium

// Given a string, your task is to count how many palindromic substrings in this string.
//
// The substrings with different start indexes or end indexes are counted as different
// substrings even they consist of same characters.
//
// Example 1:
// Input: "abc"
// Output: 3
// Explanation: Three palindromic strings: "a", "b", "c".
//
// Example 2:
// Input: "aaa"
// Output: 6
// Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
//
// Note:
// The input string length won't exceed 1000.

func countSubstrings(s string) int {
	// dp[j][i] represents whether s[j,i] is a palindromic string.
	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}

	var cnt int
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] && (j-i < 3 || dp[i+1][j-1]) {
				dp[i][j] = true
				cnt++
			}
		}
	}

	return cnt
}

// The official solution:
//
// Approach #1: Expand Around Center [Accepted]
//
// class Solution {
//     public int countSubstrings(String S) {
//         int N = S.length(), ans = 0;
//         for (int center = 0; center <= 2*N-1; ++center) {
//             int left = center / 2;
//             int right = left + center % 2;
//             while (left >= 0 && right < N && S.charAt(left) == S.charAt(right)) {
//                 ans++;
//                 left--;
//                 right++;
//             }
//         }
//         return ans;
//     }
// }
//
// Approach #2: Manacher's Algorithm [Accepted]
//
// class Solution {
//     public int countSubstrings(String S) {
//         char[] A = new char[2 * S.length() + 3];
//         A[0] = '@';
//         A[1] = '#';
//         A[A.length - 1] = '$';
//         int t = 2;
//         for (char c: S.toCharArray()) {
//             A[t++] = c;
//             A[t++] = '#';
//         }
//
//         int[] Z = new int[A.length];
//         int center = 0, right = 0;
//         for (int i = 1; i < Z.length - 1; ++i) {
//             if (i < right)
//                 Z[i] = Math.min(right - i, Z[2 * center - i]);
//             while (A[i + Z[i] + 1] == A[i - Z[i] - 1])
//                 Z[i]++;
//             if (i + Z[i] > right) {
//                 center = i;
//                 right = i + Z[i];
//             }
//         }
//         int ans = 0;
//         for (int v: Z) ans += (v + 1) / 2;
//         return ans;
//     }
// }
