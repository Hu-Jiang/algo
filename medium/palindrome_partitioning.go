package medium

// Given a string s, partition s such that every substring
// of the partition is a palindrome.
//
// Return all possible palindrome partitioning of s.
//
// Example:
// Input: "aab"
// Output:
// [
//   ["aa","b"],
//   ["a","a","b"]
// ]

func partitionString(s string) [][]string {
	if len(s) == 0 {
		return nil
	}

	var res [][]string
	backtracking(&res, []string{}, s)
	return res
}

func backtracking(res *[][]string, prev []string, remain string) {
	if len(remain) == 0 {
		*res = append(*res, append([]string{}, prev...))
		return
	}

	for i := 0; i < len(remain); i++ { // find partition position
		if !isPalindromeString(remain[:i+1]) {
			continue
		}
		prev = append(prev, remain[:i+1])
		backtracking(res, prev, remain[i+1:])
		prev = prev[:len(prev)-1]
	}
}

func isPalindromeString(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
