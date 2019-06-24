package hard

// Given an input string (s) and a pattern (p), implement wildcard pattern matching
// with support for '?' and '*'.
//
// '?' Matches any single character.
// '*' Matches any sequence of characters (including the empty sequence).
// The matching should cover the entire input string (not partial).
//
// Note:
// s could be empty and contains only lowercase letters a-z.
// p could be empty and contains only lowercase letters a-z, and characters like ? or *.
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
// p = "*"
// Output: true
// Explanation: '*' matches any sequence.
//
// Example 3:
// Input:
// s = "cb"
// p = "?a"
// Output: false
// Explanation: '?' matches 'c', but the second letter is 'a', which does not match 'b'.
//
// Example 4:
// Input:
// s = "adceb"
// p = "*a*b"
// Output: true
// Explanation: The first '*' matches the empty sequence, while the second '*' matches
// the substring "dce".
//
// Example 5:
// Input:
// s = "acdcb"
// p = "a*c?b"
// Output: false

func isMatching(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[len(s)][len(p)] = true
	for i := len(s); i >= 0; i-- {
		for j := len(p); j >= 0; j-- {
			if i == len(s) && j == len(p) {
				continue
			}
			if j != len(p) && ((i != len(s)) && (s[i] == p[j] || p[j] == '?')) {
				dp[i][j] = dp[i+1][j+1]
			} else if (j != len(p)) && p[j] == '*' {
				dp[i][j] = ((i != len(s)) && (dp[i+1][j]) || dp[i][j+1])
			} else {
				dp[i][j] = false
			}
		}
	}

	return dp[0][0]
}

// NOTE: this question similar to hard/regular_expression_matching
