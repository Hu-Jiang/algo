package medium

import (
	"sort"
)

// Given an array of strings, group anagrams together.
//
// Example:
// Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
// Output:
// [
//   ["ate","eat","tea"],
//   ["nat","tan"],
//   ["bat"]
// ]
// Note:
// All inputs will be in lowercase.
// The order of your output does not matter.

func groupAnagrams(strs []string) [][]string {
	strToStrs := make(map[string][]string)
	for _, str := range strs {
		bs := []byte(str)
		sort.Sort(byteSlice(bs))
		strToStrs[string(bs)] = append(strToStrs[string(bs)], str)
	}

	var ans [][]string
	for _, group := range strToStrs {
		ans = append(ans, group)
	}

	return ans
}

type byteSlice []byte

func (p byteSlice) Len() int           { return len(p) }
func (p byteSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p byteSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Trick 1:
// b := []byte(str)
// sort.Slice(b, func(i, j int) bool {
// 	return b[i] < b[j]
// })
// key := string(b)
//
// Trick 2:
// ss := strings.Split(str, "")
// sort.Strings(ss)
// key := strings.Join(ss, "")
