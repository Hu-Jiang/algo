package medium

// A message containing letters from A-Z is being encoded to
// numbers using the following mapping:
//
// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
//
// Given a non-empty string containing only digits, determine
// the total number of ways to decode it.
//
// Example 1:
// Input: "12"
// Output: 2
// Explanation: It could be decoded as "AB" (1 2) or "L" (12).
//
// Example 2:
// Input: "226"
// Output: 3
// Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6),
// or "BBF" (2 2 6).

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	last := []int{int(s[0] - '0')}
	for i := 1; i < len(s); i++ {
		var newLast []int
		for j := 0; j < len(last); j++ {
			split := int(s[i] - '0')
			if split != 0 {
				newLast = append(newLast, split)
			}
			appendTail := int(s[i]) - '0' + last[j]*10
			if appendTail <= 26 {
				newLast = append(newLast, appendTail)
			}
		}
		last = newLast
	}

	return len(last)
}

// Another answer from leetcode in discuss:
// I used a dp array of size n + 1 to save subproblem solutions.
// dp[0] means an empty string will have one way to decode, dp[1]
// means the way to decode a string of size 1. I then check one
// digit and two digit combination and save the results along the
// way. In the end, dp[n] will be the end result.
//
// public class Solution {
//     public int numDecodings(String s) {
//         if(s == null || s.length() == 0) {
//             return 0;
//         }
//         int n = s.length();
//         int[] dp = new int[n+1];
//         dp[0] = 1;
//         dp[1] = s.charAt(0) != '0' ? 1 : 0;
//         for(int i = 2; i <= n; i++) {
//             int first = Integer.valueOf(s.substring(i-1, i));
//             int second = Integer.valueOf(s.substring(i-2, i));
//             if(first >= 1 && first <= 9) {
//                dp[i] += dp[i-1];
//             }
//             if(second >= 10 && second <= 26) {
//                 dp[i] += dp[i-2];
//             }
//         }
//         return dp[n];
//     }
// }
