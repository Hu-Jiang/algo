package easy

import "errors"

// For strings S and T, we say "T divides S" if and only if S = T + ... + T
// (T concatenated with itself 1 or more times)
//
// Return the largest string X such that X divides str1 and X divides str2.
//
// Example 1:
// Input: str1 = "ABCABC", str2 = "ABC"
// Output: "ABC"
//
// Example 2:
// Input: str1 = "ABABAB", str2 = "ABAB"
// Output: "AB"
//
// Example 3:
// Input: str1 = "LEET", str2 = "CODE"
// Output: ""
//
// Note:
//
// 1 <= str1.length <= 1000
// 1 <= str2.length <= 1000
// str1[i] and str2[i] are English uppercase letters.

func gcdOfStrings(str1 string, str2 string) string {
	g, err := gcd(len(str1), len(str2))
	if err != nil {
		return ""
	}

	for ; g > 0; g-- {
		if len(str1)%g == 0 && len(str2)%g == 0 {
			pattern := str1[:g]
			if !isNPattern(str1, pattern, g) {
				continue
			}
			if !isNPattern(str2, pattern, g) {
				continue
			}
			return pattern
		}
	}

	return ""
}

func gcd(m, n int) (int, error) {
	if m <= 0 || n <= 0 {
		return 0, errors.New("gcd: invalid args, args <= 0")
	}

	for n != 0 {
		m, n = n, m%n
	}

	return m, nil
}

// isNPattern returns true if s contains pattern concatenated n times,
// otherwise, returns false.
func isNPattern(s, pattern string, n int) bool {
	if n == 0 || len(s)%n != 0 {
		return false
	}

	for i := 0; i < len(s)/n; i++ {
		if s[i*n:(i+1)*n] != pattern {
			return false
		}
	}

	return true
}
