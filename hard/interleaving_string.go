package hard

// Given s1, s2, s3, find whether s3 is formed by the interleaving of s1 and s2.
//
// Example 1:
// Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
// Output: true
//
// Example 2:
// Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
// Output: false

func isInterleave(s1, s2, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}

	dp := make([][]bool, len(s1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s2)+1)
	}
	dp[0][0] = true
	for i := 1; i < len(dp); i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	for j := 1; j < len(dp[0]); j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) ||
				(dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

// The official solution:
//
// Approach 1: Brute Force
//
// public class Solution {
//     public boolean is_Interleave(String s1,int i,String s2,int j,String res,String s3)
//     {
//         if(res.equals(s3) && i==s1.length() && j==s2.length())
//             return true;
//         boolean ans=false;
//         if(i<s1.length())
//             ans|=is_Interleave(s1,i+1,s2,j,res+s1.charAt(i),s3);
//         if(j<s2.length())
//             ans|=is_Interleave(s1,i,s2,j+1,res+s2.charAt(j),s3);
//         return ans;
//
//     }
//     public boolean isInterleave(String s1, String s2, String s3) {
//         return is_Interleave(s1,0,s2,0,"",s3);
//     }
// }
//
// Approach 2: Recursion with memoization
//
// public class Solution {
//     public boolean is_Interleave(String s1, int i, String s2, int j, String s3, int k, int[][] memo) {
//         if (i == s1.length()) {
//             return s2.substring(j).equals(s3.substring(k));
//         }
//         if (j == s2.length()) {
//             return s1.substring(i).equals(s3.substring(k));
//         }
//         if (memo[i][j] >= 0) {
//             return memo[i][j] == 1 ? true : false;
//         }
//         boolean ans = false;
//         if (s3.charAt(k) == s1.charAt(i) && is_Interleave(s1, i + 1, s2, j, s3, k + 1, memo)
//                 || s3.charAt(k) == s2.charAt(j) && is_Interleave(s1, i, s2, j + 1, s3, k + 1, memo)) {
//             ans = true;
//         }
//         memo[i][j] = ans ? 1 : 0;
//         return ans;
//     }
//     public boolean isInterleave(String s1, String s2, String s3) {
//         int memo[][] = new int[s1.length()][s2.length()];
//         for (int i = 0; i < s1.length(); i++) {
//             for (int j = 0; j < s2.length(); j++) {
//                 memo[i][j] = -1;
//             }
//         }
//         return is_Interleave(s1, 0, s2, 0, s3, 0, memo);
//     }
// }
//
// Approach 3: Using 2D Dynamic Programming
//
// public class Solution {
//     public boolean isInterleave(String s1, String s2, String s3) {
//         if (s3.length() != s1.length() + s2.length()) {
//             return false;
//         }
//         boolean dp[][] = new boolean[s1.length() + 1][s2.length() + 1];
//         for (int i = 0; i <= s1.length(); i++) {
//             for (int j = 0; j <= s2.length(); j++) {
//                 if (i == 0 && j == 0) {
//                     dp[i][j] = true;
//                 } else if (i == 0) {
//                     dp[i][j] = dp[i][j - 1] && s2.charAt(j - 1) == s3.charAt(i + j - 1);
//                 } else if (j == 0) {
//                     dp[i][j] = dp[i - 1][j] && s1.charAt(i - 1) == s3.charAt(i + j - 1);
//                 } else {
//                     dp[i][j] = (dp[i - 1][j] && s1.charAt(i - 1) == s3.charAt(i + j - 1)) || (dp[i][j - 1] && s2.charAt(j - 1) == s3.charAt(i + j - 1));
//                 }
//             }
//         }
//         return dp[s1.length()][s2.length()];
//     }
// }
//
// Approach 4: Using 1D Dynamic Programming
//
// public class Solution {
//     public boolean isInterleave(String s1, String s2, String s3) {
//         if (s3.length() != s1.length() + s2.length()) {
//             return false;
//         }
//         boolean dp[] = new boolean[s2.length() + 1];
//         for (int i = 0; i <= s1.length(); i++) {
//             for (int j = 0; j <= s2.length(); j++) {
//                 if (i == 0 && j == 0) {
//                     dp[j] = true;
//                 } else if (i == 0) {
//                     dp[j] = dp[j - 1] && s2.charAt(j - 1) == s3.charAt(i + j - 1);
//                 } else if (j == 0) {
//                     dp[j] = dp[j] && s1.charAt(i - 1) == s3.charAt(i + j - 1);
//                 } else {
//                     dp[j] = (dp[j] && s1.charAt(i - 1) == s3.charAt(i + j - 1)) || (dp[j - 1] && s2.charAt(j - 1) == s3.charAt(i + j - 1));
//                 }
//             }
//         }
//         return dp[s2.length()];
//     }
// }
