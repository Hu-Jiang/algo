package hard

// Given a string S and a string T, find the minimum window in S which
// will contain all the characters in T in complexity O(n).
//
// Example:
// Input: S = "ADOBECODEBANC", T = "ABC"
// Output: "BANC"
//
// Note:
// If there is no such window in S that covers all characters in T,
// return the empty string "".
// If there is such window, you are guaranteed that there will always
// be only one unique minimum window in S.

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	l, r := 0, 0 // contain left and right
	tcnt := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tcnt[t[i]]++
	}

	scnt := make(map[byte]int)
	scnt[s[0]]++

	res := struct {
		l, r int
		len  int
	}{-1, -1, len(s) + 1}

	for l < len(s) && r < len(s) {
		if contain(scnt, tcnt) {
			if r-l+1 < res.len {
				res.l, res.r, res.len = l, r, r-l+1
			}
			l++
			if l < len(s) {
				scnt[s[l-1]]--
			}
		} else {
			r++
			if r < len(s) {
				scnt[s[r]]++
			}
		}
	}

	if res.len > len(s) {
		return ""
	}

	return s[res.l : res.r+1]
}

func contain(src, dst map[byte]int) bool {
	for k, v := range dst {
		if src[k] < v {
			return false
		}
	}
	return true
}

// The official solution:
//
// Approach 1: Sliding Window
//
// class Solution {
// 	public String minWindow(String s, String t) {
//
// 		if (s.length() == 0 || t.length() == 0) {
// 			return "";
// 		}
//
// 		// Dictionary which keeps a count of all the unique characters in t.
// 		Map<Character, Integer> dictT = new HashMap<Character, Integer>();
// 		for (int i = 0; i < t.length(); i++) {
// 			int count = dictT.getOrDefault(t.charAt(i), 0);
// 			dictT.put(t.charAt(i), count + 1);
// 		}
//
// 		// Number of unique characters in t, which need to be present in the desired window.
// 		int required = dictT.size();
//
// 		// Left and Right pointer
// 		int l = 0, r = 0;
//
// 		// formed is used to keep track of how many unique characters in t
// 		// are present in the current window in its desired frequency.
// 		// e.g. if t is "AABC" then the window must have two A's, one B and one C.
// 		// Thus formed would be = 3 when all these conditions are met.
// 		int formed = 0;
//
// 		// Dictionary which keeps a count of all the unique characters in the current window.
// 		Map<Character, Integer> windowCounts = new HashMap<Character, Integer>();
//
// 		// ans list of the form (window length, left, right)
// 		int[] ans = {-1, 0, 0};
//
// 		while (r < s.length()) {
// 			// Add one character from the right to the window
// 			char c = s.charAt(r);
// 			int count = windowCounts.getOrDefault(c, 0);
// 			windowCounts.put(c, count + 1);
//
// 			// If the frequency of the current character added equals to the
// 			// desired count in t then increment the formed count by 1.
// 			if (dictT.containsKey(c) && windowCounts.get(c).intValue() == dictT.get(c).intValue()) {
// 				formed++;
// 			}
//
// 			// Try and contract the window till the point where it ceases to be 'desirable'.
// 			while (l <= r && formed == required) {
// 				c = s.charAt(l);
// 				// Save the smallest window until now.
// 				if (ans[0] == -1 || r - l + 1 < ans[0]) {
// 					ans[0] = r - l + 1;
// 					ans[1] = l;
// 					ans[2] = r;
// 				}
//
// 				// The character at the position pointed by the
// 				// `Left` pointer is no longer a part of the window.
// 				windowCounts.put(c, windowCounts.get(c) - 1);
// 				if (dictT.containsKey(c) && windowCounts.get(c).intValue() < dictT.get(c).intValue()) {
// 					formed--;
// 				}
//
// 				// Move the left pointer ahead, this would help to look for a new window.
// 				l++;
// 			}
//
// 			// Keep expanding the window once we are done contracting.
// 			r++;
// 		}
//
// 		return ans[0] == -1 ? "" : s.substring(ans[1], ans[2] + 1);
// 	}
// }
