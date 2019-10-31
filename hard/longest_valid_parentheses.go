package hard

// Given a string containing just the characters '(' and ')', find
// the length of the longest valid (well-formed) parentheses substring.
//
// Example 1:
// Input: "(()"
// Output: 2
// Explanation: The longest valid parentheses substring is "()"
//
// Example 2:
// Input: ")()())"
// Output: 4
// Explanation: The longest valid parentheses substring is "()()"

func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}

	var (
		maxMatch int
		dp       = make([]int, len(s))
	)

	dp[0] = 0
	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			dp[i] = 0
		} else {
			if s[i-1] == '(' {
				if i-2 >= 0 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else {
				if i-dp[i-1]-1 < 0 {
					dp[i] = 0
				} else {
					if s[i-dp[i-1]-1] == '(' {
						if i-dp[i-1]-2 < 0 {
							dp[i] = dp[i-1] + 2
						} else {
							dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
						}
					} else {
						dp[i] = 0
					}
				}
			}
		}
		if dp[i] > maxMatch {
			maxMatch = dp[i]
		}
	}

	return maxMatch
}

// The official solution:
//
// Approach 1: Brute Force
//
// public class Solution {
//     public boolean isValid(String s) {
//         Stack<Character> stack = new Stack<Character>();
//         for (int i = 0; i < s.length(); i++) {
//             if (s.charAt(i) == '(') {
//                 stack.push('(');
//             } else if (!stack.empty() && stack.peek() == '(') {
//                 stack.pop();
//             } else {
//                 return false;
//             }
//         }
//         return stack.empty();
//     }
//     public int longestValidParentheses(String s) {
//         int maxlen = 0;
//         for (int i = 0; i < s.length(); i++) {
//             for (int j = i + 2; j <= s.length(); j+=2) {
//                 if (isValid(s.substring(i, j))) {
//                     maxlen = Math.max(maxlen, j - i);
//                 }
//             }
//         }
//         return maxlen;
//     }
// }
//
// Approach 2: Using Dynamic Programming
//
// public class Solution {
//     public int longestValidParentheses(String s) {
//         int maxans = 0;
//         int dp[] = new int[s.length()];
//         for (int i = 1; i < s.length(); i++) {
//             if (s.charAt(i) == ')') {
//                 if (s.charAt(i - 1) == '(') {
//                     dp[i] = (i >= 2 ? dp[i - 2] : 0) + 2;
//                 } else if (i - dp[i - 1] > 0 && s.charAt(i - dp[i - 1] - 1) == '(') {
//                     dp[i] = dp[i - 1] + ((i - dp[i - 1]) >= 2 ? dp[i - dp[i - 1] - 2] : 0) + 2;
//                 }
//                 maxans = Math.max(maxans, dp[i]);
//             }
//         }
//         return maxans;
//     }
// }
//
// Approach 3: Using Stack
//
// public class Solution {
//
//     public int longestValidParentheses(String s) {
//         int maxans = 0;
//         Stack<Integer> stack = new Stack<>();
//         stack.push(-1);
//         for (int i = 0; i < s.length(); i++) {
//             if (s.charAt(i) == '(') {
//                 stack.push(i);
//             } else {
//                 stack.pop();
//                 if (stack.empty()) {
//                     stack.push(i);
//                 } else {
//                     maxans = Math.max(maxans, i - stack.peek());
//                 }
//             }
//         }
//         return maxans;
//     }
// }
//
// Approach 4: Without extra space
//
// public class Solution {
//     public int longestValidParentheses(String s) {
//         int left = 0, right = 0, maxlength = 0;
//         for (int i = 0; i < s.length(); i++) {
//             if (s.charAt(i) == '(') {
//                 left++;
//             } else {
//                 right++;
//             }
//             if (left == right) {
//                 maxlength = Math.max(maxlength, 2 * right);
//             } else if (right >= left) {
//                 left = right = 0;
//             }
//         }
//         left = right = 0;
//         for (int i = s.length() - 1; i >= 0; i--) {
//             if (s.charAt(i) == '(') {
//                 left++;
//             } else {
//                 right++;
//             }
//             if (left == right) {
//                 maxlength = Math.max(maxlength, 2 * left);
//             } else if (left >= right) {
//                 left = right = 0;
//             }
//         }
//         return maxlength;
//     }
// }
