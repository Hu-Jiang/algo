package easy

// Let's define a function f(s) over a non-empty string s, which calculates
// the frequency of the smallest character in s. For example, if s = "dcce"
// then f(s) = 2 because the smallest character is "c" and its frequency is 2.
//
// Now, given string arrays queries and words, return an integer array answer,
// where each answer[i] is the number of words such that f(queries[i]) < f(W),
// where W is a word in words.
//
// Example 1:
// Input: queries = ["cbd"], words = ["zaaaz"]
// Output: [1]
// Explanation: On the first query we have f("cbd") = 1, f("zaaaz") = 3 so
// f("cbd") < f("zaaaz").
//
// Example 2:
// Input: queries = ["bbb","cc"], words = ["a","aa","aaa","aaaa"]
// Output: [1,2]
// Explanation: On the first query only f("bbb") < f("aaaa"). On the second
// query both f("aaa") and f("aaaa") are both > f("cc").
//
// Constraints:
// 1 <= queries.length <= 2000
// 1 <= words.length <= 2000
// 1 <= queries[i].length, words[i].length <= 10
// queries[i][j], words[i][j] are English lowercase letters.

func numSmallerByFrequency(queries []string, words []string) []int {
	var freqs []int
	for i := 0; i < len(words); i++ {
		freqs = append(freqs, smallestCharFreq(words[i]))
	}

	var res []int
	for i := 0; i < len(queries); i++ {
		qfreq := smallestCharFreq(queries[i])
		cnt := 0
		for j := 0; j < len(freqs); j++ {
			if qfreq < freqs[j] {
				cnt++
			}
		}
		res = append(res, cnt)
	}
	return res
}

func smallestCharFreq(str string) int {
	freq := make([]int, 26)
	for _, r := range str {
		freq[r-'a']++
	}

	for i := 0; i < len(freq); i++ {
		if freq[i] > 0 {
			return freq[i]
		}
	}
	return -1
}

// The official Hint:
// - For each string from words calculate the leading count and store
// it in an array, then sort the integer array.
// - For each string from queries calculate the leading count "p" and
// in base of the sorted array calculated on the step 1 do a binary
// search to count the number of items greater than "p".
