package hard

// Given an input string (s) and a pattern (p), implement regular expression matching
// with support for '.' and '*'.
//
// '.' Matches any single character.
// '*' Matches zero or more of the preceding element.
//
// The matching should cover the entire input string (not partial).
//
// Note:
// s could be empty and contains only lowercase letters a-z.
// p could be empty and contains only lowercase letters a-z, and characters like . or *.
//
// Example 1:
// Input:
// s = "aa"
// p = "a"
// Output: false
// Explanation: "a" does not match the entire string "aa".
//
// Example 2:
// Input:
// s = "aa"
// p = "a*"
// Output: true
// Explanation: '*' means zero or more of the precedeng element, 'a'.
// Therefore, by repeating 'a' once, it becomes "aa".
//
// Example 3:
// Input:
// s = "ab"
// p = ".*"
// Output: true
// Explanation: ".*" means "zero or more (*) of any character (.)".
//
// Example 4:
// Input:
// s = "aab"
// p = "c*a*b"
// Output: true
// Explanation: c can be repeated 0 times, a can be repeated 1 time.
// Therefore it matches "aab".
//
// Example 5:
// Input:
// s = "mississippi"
// p = "mis*is*p*."
// Output: false

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	firstMatch := (len(s) != 0) && (s[0] == p[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) ||
			(firstMatch && isMatch(s[1:], p))
	} else {
		return firstMatch && isMatch(s[1:], p[1:])
	}
}

// The official soloution:
//
// Approach 1: Recursion
// class Solution {
//     public boolean isMatch(String text, String pattern) {
//         if (pattern.isEmpty()) return text.isEmpty();
//         boolean first_match = (!text.isEmpty() &&
//                                (pattern.charAt(0) == text.charAt(0) || pattern.charAt(0) == '.'));
//
//         if (pattern.length() >= 2 && pattern.charAt(1) == '*'){
//             return (isMatch(text, pattern.substring(2)) ||
//                     (first_match && isMatch(text.substring(1), pattern)));
//         } else {
//             return first_match && isMatch(text.substring(1), pattern.substring(1));
//         }
//     }
// }
//
// Approach 2: Dynamic Programming
//
// Top-Down Variation:
//
// enum Result {
//     TRUE, FALSE
// }
//
// class Solution {
//     Result[][] memo;
//
//     public boolean isMatch(String text, String pattern) {
//         memo = new Result[text.length() + 1][pattern.length() + 1];
//         return dp(0, 0, text, pattern);
//     }
//
//     public boolean dp(int i, int j, String text, String pattern) {
//         if (memo[i][j] != null) {
//             return memo[i][j] == Result.TRUE;
//         }
//         boolean ans;
//         if (j == pattern.length()){
//             ans = i == text.length();
//         } else{
//             boolean first_match = (i < text.length() &&
//                                    (pattern.charAt(j) == text.charAt(i) ||
//                                     pattern.charAt(j) == '.'));
//
//             if (j + 1 < pattern.length() && pattern.charAt(j+1) == '*'){
//                 ans = (dp(i, j+2, text, pattern) ||
//                        first_match && dp(i+1, j, text, pattern));
//             } else {
//                 ans = first_match && dp(i+1, j+1, text, pattern);
//             }
//         }
//         memo[i][j] = ans ? Result.TRUE : Result.FALSE;
//         return ans;
//     }
// }
// Bottom-Up Variation:
//
// class Solution(object):
//     def isMatch(self, text, pattern):
//         dp = [[False] * (len(pattern) + 1) for _ in range(len(text) + 1)]
//
//         dp[-1][-1] = True
//         for i in range(len(text), -1, -1):
//             for j in range(len(pattern) - 1, -1, -1):
//                 first_match = i < len(text) and pattern[j] in {text[i], '.'}
//                 if j+1 < len(pattern) and pattern[j+1] == '*':
//                     dp[i][j] = dp[i][j+2] or first_match and dp[i+1][j]
//                 else:
//                     dp[i][j] = first_match and dp[i+1][j+1]
//
//         return dp[0][0]
