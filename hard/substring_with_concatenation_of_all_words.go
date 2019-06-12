package hard

// You are given a string, s, and a list of words, words, that are all of the same length.
// Find all starting indices of substring(s) in s that is a concatenation of each word in
// words exactly once and without any intervening characters.
//
// Example 1:
// Input:
//   s = "barfoothefoobarman",
//   words = ["foo","bar"]
// Output: [0,9]
// Explanation: Substrings starting at index 0 and 9 are "barfoor" and "foobar" respectively.
// The output order does not matter, returning [9,0] is fine too.
//
// Example 2:
// Input:
//   s = "wordgoodgoodgoodbestword",
//   words = ["word","good","best","word"]
// Output: []

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return nil
	}

	wordsCnt := genWordsMap(words)
	totalLen := charsLen(words)
	wordLen := len(words[0])
	var ans []int

	for i := 0; i <= len(s)-totalLen; i++ {
		sub := s[i : i+totalLen]
		subWordsCnt := make(map[string]int)
		for j := 0; j < len(sub); j += wordLen {
			subWordsCnt[sub[j:j+wordLen]] += subWordsCnt[sub[j:j+wordLen]] + 1
		}
		if equalWordsCnt(wordsCnt, subWordsCnt) {
			ans = append(ans, i)
		}
	}

	return ans
}

func genWordsMap(words []string) map[string]int {
	wordsCnt := make(map[string]int)
	for _, word := range words {
		wordsCnt[word] += wordsCnt[word] + 1
	}

	return wordsCnt
}

func charsLen(words []string) int {
	var length int
	for _, word := range words {
		length += len(word)
	}

	return length
}

func equalWordsCnt(wc1, wc2 map[string]int) bool {
	if len(wc1) != len(wc2) {
		return false
	}

	for k, v := range wc1 {
		if wc2[k] != v {
			return false
		}
	}

	return true
}
