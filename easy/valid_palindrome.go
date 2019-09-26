package easy

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	ls := strings.ToLower(alphanumeric(s))
	for i := 0; i < len(ls)/2; i++ {
		if ls[i] != ls[len(ls)-1-i] {
			return false
		}
	}
	return true
}

// alphanumeric returns all alphanumeric characters in s.
func alphanumeric(s string) string {
	var bs strings.Builder
	for _, r := range s {
		if unicode.IsNumber(r) || unicode.IsLetter(r) {
			bs.WriteRune(r)
		}
	}
	return bs.String()
}
