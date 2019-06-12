package easy

// Implement strStr().
// Return the index of the first occurrence of needle in haystack,
// or -1 if needle is not part of haystack.
//
// Example 1:
// Input: haystack = "hello", needle = "ll"
// Output: 2
//
// Example 2:
// Input: haystack = "aaaaa", needle = "bba"
// Output: -1
//
// Clarification:
// What should we return when needle is an empty string? This is
// a great question to ask during an interview.
//
// For the purpose of this problem, we will return 0 when needle is
// an empty string. This is consistent to C's strstr() and Java's indexOf().

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(haystack) == 0 {
		return -1
	}

	for i := 0; i < len(haystack); i++ {
		var j int
		for j = i; j < len(haystack) && j-i < len(needle); j++ {
			if haystack[j] != needle[j-i] {
				break
			}
		}

		if j-i == len(needle) {
			return i
		}
	}

	return -1
}
