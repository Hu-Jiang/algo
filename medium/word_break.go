package medium

// Given a non-empty string s and a dictionary wordDict containing a list of non-empty words,
// determine if s can be segmented into a space-separated sequence of one or more dictionary words.
//
// Note:
// The same word in the dictionary may be reused multiple times in the segmentation.
// You may assume the dictionary does not contain duplicate words.
//
// Example 1:
// Input: s = "leetcode", wordDict = ["leet", "code"]
// Output: true
// Explanation: Return true because "leetcode" can be segmented as "leet code".
//
// Example 2:
// Input: s = "applepenapple", wordDict = ["apple", "pen"]
// Output: true
// Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
//              Note that you are allowed to reuse a dictionary word.
//
// Example 3:
// Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
// Output: false

func wordBreak(s string, wordDict []string) bool {
	// dp[i]:
	// 		Represents whether s[:i+1] can be segmented into a
	// 		space-separated sequence of one or more dictionary words.

	dp := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		for _, word := range wordDict {
			if (i-len(word)+1 >= 0 && s[i-len(word)+1:i+1] == word) &&
				(i-len(word) < 0 || dp[i-len(word)]) {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(dp)-1]
}
