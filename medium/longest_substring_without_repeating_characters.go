package medium

import (
	"strings"
)

// Given a string, find the length of the longest substring without repeating characters.
//
// Example 1:
// Input: "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
//
// Example 2:
// Input: "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
//
// Example 3:
// Input: "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
//              Note that the answer must be a substring,
//				"pwke" is a subsequence and not a substring.

func lengthOfLongestSubstring(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		if ans >= len(s)-i {
			return ans
		}

		var tmp string
		for j := i; j < len(s); j++ {
			if strings.ContainsRune(tmp, rune(s[j])) {
				break
			}
			tmp += string(s[j])
		}

		if len(tmp) > ans {
			ans = len(tmp)
		}
	}

	return ans
}

// The official solution:
//
// Approach 1: Brute Force
//
// public class Solution {
//     public int lengthOfLongestSubstring(String s) {
//         int n = s.length();
//         int ans = 0;
//         for (int i = 0; i < n; i++)
//             for (int j = i + 1; j <= n; j++)
//                 if (allUnique(s, i, j)) ans = Math.max(ans, j - i);
//         return ans;
//     }
//
//     public boolean allUnique(String s, int start, int end) {
//         Set<Character> set = new HashSet<>();
//         for (int i = start; i < end; i++) {
//             Character ch = s.charAt(i);
//             if (set.contains(ch)) return false;
//             set.add(ch);
//         }
//         return true;
//     }
// }
//
// Approach 2: Sliding Window
//
// public class Solution {
//     public int lengthOfLongestSubstring(String s) {
//         int n = s.length();
//         Set<Character> set = new HashSet<>();
//         int ans = 0, i = 0, j = 0;
//         while (i < n && j < n) {
//             // try to extend the range [i, j]
//             if (!set.contains(s.charAt(j))){
//                 set.add(s.charAt(j++));
//                 ans = Math.max(ans, j - i);
//             }
//             else {
//                 set.remove(s.charAt(i++));
//             }
//         }
//         return ans;
//     }
// }
//
// Approach 3: Sliding Window Optimized
//
// public class Solution {
//     public int lengthOfLongestSubstring(String s) {
//         int n = s.length(), ans = 0;
//         Map<Character, Integer> map = new HashMap<>(); // current index of character
//         // try to extend the range [i, j]
//         for (int j = 0, i = 0; j < n; j++) {
//             if (map.containsKey(s.charAt(j))) {
//                 i = Math.max(map.get(s.charAt(j)), i);
//             }
//             ans = Math.max(ans, j - i + 1);
//             map.put(s.charAt(j), j + 1);
//         }
//         return ans;
//     }
// }
//
// Approach 3 variation: Assuming ASCII 128
//
// public class Solution {
//     public int lengthOfLongestSubstring(String s) {
//         int n = s.length(), ans = 0;
//         int[] index = new int[128]; // current index of character
//         // try to extend the range [i, j]
//         for (int j = 0, i = 0; j < n; j++) {
//             i = Math.max(index[s.charAt(j)], i);
//             ans = Math.max(ans, j - i + 1);
//             index[s.charAt(j)] = j + 1;
//         }
//         return ans;
//     }
// }
