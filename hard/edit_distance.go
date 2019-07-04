package hard

// Given two words word1 and word2, find the minimum number
// of operations required to convert word1 to word2.
//
// You have the following 3 operations permitted on a word:
// Insert a character
// Delete a character
// Replace a character
//
// Example 1:
// Input: word1 = "horse", word2 = "ros"
// Output: 3
// Explanation:
// horse -> rorse (replace 'h' with 'r')
// rorse -> rose (remove 'r')
// rose -> ros (remove 'e')
//
// Example 2:
// Input: word1 = "intention", word2 = "execution"
// Output: 5
// Explanation:
// intention -> inention (remove 't')
// inention -> enention (replace 'i' with 'e')
// enention -> exention (replace 'n' with 'x')
// exention -> exection (replace 'n' with 'c')
// exection -> execution (insert 'u')

func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}

	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = j
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			diagonal := 1
			if word1[i-1] == word2[j-1] {
				diagonal = 0
			}
			dp[i][j] = minDis(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+diagonal)
		}
	}
	return dp[len(word1)][len(word2)]
}

func minDis(a, b, c int) int {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

// My Time Limit Exceeded Recursion Version:
//
// func minDistance(word1 string, word2 string) int {
// 	if len(word1) == 0 {
// 		return len(word2)
// 	}
// 	if len(word2) == 0 {
// 		return len(word1)
// 	}
//
// 	m1 := minDistance(word1[:len(word1)-1], word2) + 1
// 	m2 := minDistance(word1, word2[:len(word2)-1]) + 1
// 	m3 := minDistance(word1[:len(word1)-1], word2[:len(word2)-1])
// 	if word1[len(word1)-1] != word2[len(word2)-1] {
// 		m3 += 1
// 	}
// 	return minDis(m1, m2, m3)
// }
