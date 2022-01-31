package medium

import (
	"sort"
)

// Given a collection of integers that might contain duplicates,
// nums, return all possible subsets (the power set).
//
// Note: The solution set must not contain duplicate subsets.
//
// Example:
// Input: [1,2,2]
// Output:
// [
//   [2],
//   [1],
//   [1,2,2],
//   [2,2],
//   [1,2],
//   []
// ]

func subsetsWithDup(nums []int) [][]int {
	s := newSets(nums, true)
	s.generate(0, []int{})
	return s.subsets
}

type sets struct {
	nums    []int // nums will be sorted when new sets
	nodup   bool
	subsets [][]int
}

func newSets(nums []int, nodup bool) *sets {
	sort.Ints(nums)
	return &sets{
		nums:    nums,
		nodup:   nodup,
		subsets: make([][]int, 0),
	}
}

func (s *sets) generate(idx int, prev []int) {
	if idx != len(s.nums) {
		s.generate(idx+1, append(prev, s.nums[idx]))
		s.generate(idx+1, prev)
		return
	}

	if !s.nodup {
		s.append(prev)
		return
	}

	if s.find(prev) {
		return
	}

	s.append(prev)
}

func (s *sets) append(set []int) {
	// here need to deep copy with set
	s.subsets = append(s.subsets, append([]int{}, set...))
}

func (s *sets) find(item []int) bool {
	for i := 0; i < len(s.subsets); i++ {
		if len(s.subsets[i]) != len(item) {
			continue
		}
		find := true
		for j := 0; j < len(item); j++ {
			if s.subsets[i][j] != item[j] {
				find = false
				break
			}
		}
		if find {
			return true
		}
	}
	return false
}
