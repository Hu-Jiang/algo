package hard

// Given a non-empty string s and a dictionary wordDict containing a list of non-empty words,
// add spaces in s to construct a sentence where each word is a valid dictionary word.
// Return all such possible sentences.
//
// Note:
// The same word in the dictionary may be reused multiple times in the segmentation.
// You may assume the dictionary does not contain duplicate words.
//
// Example 1:
// Input:
// s = "catsanddog"
// wordDict = ["cat", "cats", "and", "sand", "dog"]
// Output:
// [
//   "cats and dog",
//   "cat sand dog"
// ]
//
// Example 2:
// Input:
// s = "pineapplepenapple"
// wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
// Output:
// [
//   "pine apple pen apple",
//   "pineapple pen apple",
//   "pine applepen apple"
// ]
// Explanation: Note that you are allowed to reuse a dictionary word.
//
// Example 3:
// Input:
// s = "catsandog"
// wordDict = ["cats", "dog", "sand", "and", "cat"]
// Output:
// []

func wordBreak(s string, wordDict []string) []string {
	// dp[i]:
	// 		Represents whether s[:i+1] can be segmented into a
	// 		space-separated sequence of one or more dictionary words.
	//
	// path[i]:
	// 		Record cut position of parent
	var (
		dp   = make([]bool, len(s))
		path = make([][]int, len(s))
	)
	for i := 0; i < len(s); i++ {
		for _, word := range wordDict {
			if (i-len(word)+1 >= 0 && s[i-len(word)+1:i+1] == word) &&
				(i-len(word) < 0 || dp[i-len(word)]) {
				dp[i] = true
				path[i] = append(path[i], i-len(word))
			}
		}
	}

	if !dp[len(dp)-1] {
		return nil
	}

	return outputPath(s, len(dp)-1, path, map[int][]string{})
}

func outputPath(s string, i int, path [][]int, memo map[int][]string) []string {
	if i < 0 {
		return nil
	}

	if v, ok := memo[i]; ok {
		return v
	}

	var res []string
	for j := 0; j < len(path[i]); j++ {
		prevs := outputPath(s, path[i][j], path, memo)
		if len(prevs) == 0 {
			res = append(res, s[:i+1])
			continue
		}
		for k := 0; k < len(prevs); k++ {
			res = append(res, prevs[k]+" "+s[path[i][j]+1:i+1])
		}
	}

	memo[i] = res
	return res
}
