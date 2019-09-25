package hard

import "math"

// Say you have an array for which the i-th element is the price of a given stock on day i.
//
// Design an algorithm to find the maximum profit. You may complete at most two transactions.
//
// Note: You may not engage in multiple transactions at the same time (i.e., you must sell
// the stock before you buy again).
//
// Example 1:
// Input: [3,3,5,0,0,3,1,4]
// Output: 6
// Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
//              Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
//
// Example 2:
// Input: [1,2,3,4,5]
// Output: 4
// Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
//              Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
//              engaging multiple transactions at the same time. You must sell before buying again.
//
// Example 3:
// Input: [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transaction is done, i.e. max profit = 0.

func maxProfit(prices []int) int {
	// define:
	// dp[i][k][j] -- max profit in i day, pass maximum k transaction,
	// 				whether hold on stock.
	//
	// range:
	// i: [1, len(prices)]
	// k: [0, 2]
	// j: [0, 1]
	//
	// transfer equation:
	// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
	// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
	if len(prices) == 0 {
		return 0
	}

	dp := make([][][]int, len(prices)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, 3)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	for k := 0; k < 3; k++ {
		dp[0][k][1] = math.MinInt32
	}
	for i := 1; i <= len(prices); i++ {
		for k := 1; k <= 2; k++ {
			dp[i][k][0] = maxPrice(dp[i-1][k][0], dp[i-1][k][1]+prices[i-1])
			dp[i][k][1] = maxPrice(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i-1])
		}
	}

	return dp[len(prices)][2][0]
}

func maxPrice(a, b int) int {
	if a > b {
		return a
	}
	return b
}
