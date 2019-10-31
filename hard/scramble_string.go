package hard

// Given a string s1, we may represent it as a binary tree by
// partitioning it to two non-empty substrings recursively.
//
// Below is one possible representation of s1 = "great":
//     great
//    /    \
//   gr    eat
//  / \    /  \
// g   r  e   at
//            / \
//           a   t
// To scramble the string, we may choose any non-leaf node and
// swap its two children.
//
// For example, if we choose the node "gr" and swap its two children,
// it produces a scrambled string "rgeat".
//     rgeat
//    /    \
//   rg    eat
//  / \    /  \
// r   g  e   at
//            / \
//           a   t
// We say that "rgeat" is a scrambled string of "great".
//
// Similarly, if we continue to swap the children of nodes "eat" and
// "at", it produces a scrambled string "rgtae".
//     rgtae
//    /    \
//   rg    tae
//  / \    /  \
// r   g  ta  e
//        / \
//       t   a
// We say that "rgtae" is a scrambled string of "great".
//
// Given two strings s1 and s2 of the same length, determine if s2
// is a scrambled string of s1.
//
// Example 1:
// Input: s1 = "great", s2 = "rgeat"
// Output: true
//
// Example 2:
// Input: s1 = "abcde", s2 = "caebd"
// Output: false

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}

	s1CharCnt := make(map[rune]int)
	s2CharCnt := make(map[rune]int)
	for _, r := range s1 {
		s1CharCnt[r]++
	}
	for _, r := range s2 {
		s2CharCnt[r]++
	}
	if !equal(s1CharCnt, s2CharCnt) {
		return false
	}

	for i := 1; i < len(s1); i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		if isScramble(s1[:i], s2[len(s2)-i:]) && isScramble(s1[i:], s2[:len(s2)-i]) {
			return true
		}
	}
	return false
}

func equal(a, b map[rune]int) bool {
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

// Another solution from leetcode discuss:
//
// func isScramble(s1 string, s2 string) bool {
//
// 	if s1 == s2 {
// 		return true
// 	} else if len(s1) != len(s2) {
// 		return false
// 	}
//
// 	letters := make([]int, 26)
//
// 	for i := 0; i < len(s1); i++ {
// 		letters[int(s1[i])-int('a')]++
// 		letters[int(s2[i])-int('a')]--
// 	}
//
// 	for i := 0; i < len(s1); i++ {
// 		if letters[int(s1[i])-int('a')] != 0 {
// 			return false
// 		}
// 	}
//
// 	for i := 1; i < len(s1); i++ {
//
// 		if isScramble(s1[0:i], s2[0:i]) && isScramble(s1[i:], s2[i:]) {
// 			return true
// 		}
// 		if isScramble(s1[0:i], s2[len(s2)-i:]) && isScramble(s1[i:], s2[0:len(s2)-i]) {
// 			return true
// 		}
// 	}
// 	return false
// }
