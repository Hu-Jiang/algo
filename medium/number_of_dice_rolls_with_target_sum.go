package medium

import "math"

// You have d dice, and each die has f faces numbered 1, 2, ..., f.
//
// Return the number of possible ways (out of fd total ways) modulo 10^9 + 7
// to roll the dice so the sum of the face up numbers equals target.
//
// Example 1:
// Input: d = 1, f = 6, target = 3
// Output: 1
// Explanation:
// You throw one die with 6 faces.  There is only one way to get a sum of 3.
//
// Example 2:
// Input: d = 2, f = 6, target = 7
// Output: 6
// Explanation:
// You throw two dice, each with 6 faces.  There are 6 ways to get a sum of 7:
// 1+6, 2+5, 3+4, 4+3, 5+2, 6+1.
//
// Example 3:
// Input: d = 2, f = 5, target = 10
// Output: 1
// Explanation:
// You throw two dice, each with 5 faces.  There is only one way to get a sum of 10: 5+5.
//
// Example 4:
// Input: d = 1, f = 2, target = 3
// Output: 0
// Explanation:
// You throw one die with 2 faces.  There is no way to get a sum of 3.
//
// Example 5:
// Input: d = 30, f = 30, target = 500
// Output: 222616187
// Explanation:
// The answer must be returned modulo 10^9 + 7.
//
// Constraints:
// 1 <= d, f <= 30
// 1 <= target <= 1000

func numRollsToTarget(d int, f int, target int) int {
	if target < d || target > d*f {
		return 0
	}

	mod := int(math.Pow10(9) + 7)

	var dp [31][1001]int
	dp[0][0] = 1
	for i := 1; i <= d; i++ {
		for num := i; num <= i*f; num++ {
			for j := 1; j <= f && num-j >= 0; j++ {
				dp[i][num] = (dp[i][num] + dp[i-1][num-j]) % mod
			}
		}
	}

	return dp[d][target]
}

// My Time Limit Exceeded Version: [time complexity: O(f^d)]
//
// func numRollsToTarget(d int, f int, target int) int {
// 	var res int
// 	backtrack(d, f, target, 0, &res)
// 	return res
// }
//
// func backtrack(d int, f int, target int, prevSum int, res *int) {
// 	if target == prevSum {
// 		*res = (*res + 1) % (int(math.Pow10(9)) + 7)
// 		return
// 	}
// 	if d == 0 || prevSum > prevSum {
// 		return
// 	}
//
// 	for i := 1; i <= f; i++ {
// 		backtrack(d-1, f, target, prevSum+i, res)
// 	}
// }
