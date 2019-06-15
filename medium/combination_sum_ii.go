package medium

import (
	"sort"
)

// Given a collection of candidate numbers (candidates) and a target
// number (target), find all unique combinations in candidates where
// the candidate numbers sums to target.
//
// Each number in candidates may only be used once in the combination.
//
// Note:
// All numbers (including target) will be positive integers.
// The solution set must not contain duplicate combinations.
//
// Example 1:
// Input: candidates = [10,1,2,7,6,1,5], target = 8,
// A solution set is:
// [
//   [1, 7],
//   [1, 2, 5],
//   [2, 6],
//   [1, 1, 6]
// ]
//
// Example 2:
// Input: candidates = [2,5,2,1,2], target = 5,
// A solution set is:
// [
//   [1,2,2],
//   [5]
// ]

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}

	sort.Ints(candidates)
	ans := make([][]int, 0)
	doCombination2(candidates, target, 0, nil, &ans)

	return ans
}

func doCombination2(candidates []int, target int, prevSum int, prev []int, ans *[][]int) {
	if prevSum == target {
		*ans = append(*ans, append([]int{}, prev...))
		return
	}

	if prevSum > target {
		return
	}

	for i := 0; i < len(candidates); i++ {
		if i != 0 && candidates[i] == candidates[i-1] {
			continue
		}

		doCombination2(candidates[i+1:], target, prevSum+candidates[i], append(prev, candidates[i]), ans)
	}
}
