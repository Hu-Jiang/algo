package medium

// Given a string s, find the longest palindromic substring in s.
// You may assume that the maximum length of s is 1000.
//
// Example 1:
//
// Input: "babad"
// Output: "bab"
// Note: "aba" is also a valid answer.
//
// Example 2:
//
// Input: "cbbd"
// Output: "bb"

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	var longest string
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if j-i+1 > len(longest) && isPalindrome(s[i:j+1]) {
				longest = s[i : j+1]
			}
		}
	}
	return longest
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

// The official solution:
// [https://leetcode.com/articles/longest-palindromic-substring/]
// NOTE: pay attention to dynamic programming
