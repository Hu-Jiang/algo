package medium

// Given n, how many structurally unique BST's (binary search trees)
// that store values 1 ... n?
//
// Example:
// Input: 3
// Output: 5
// Explanation:
// Given n = 3, there are a total of 5 unique BST's:
//    1         3     3      2      1
//     \       /     /      / \      \
//      3     2     1      1   3      2
//     /     /       \                 \
//    2     1         2                 3

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= i-1; j++ {
			dp[i] += dp[j] * dp[i-1-j]
		}
	}
	return dp[n]
}

// Intuition:
// G(n): the answer
// F(i, n): use i as root in total n node, how many unique BST?
//
// Then:
// G(n) = F(1, n) + F(2, n) + ... + F(n, n)
// F(i, n) = G(i-1) * G(n-i)
//
// Thus:
// G(n) = G(0) * G(n-1) + G(1) * G(n-2) + ... + G(n-1) * G(0)
