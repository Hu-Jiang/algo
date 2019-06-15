package medium

// Given a set of candidate numbers (candidates) (without duplicates)
// and a target number (target), find all unique combinations in
// candidates where the candidate numbers sums to target.
//
// The same repeated number may be chosen from candidates unlimited
// number of times.
//
// Note:
// All numbers (including target) will be positive integers.
// The solution set must not contain duplicate combinations.
//
// Example 1:
// Input: candidates = [2,3,6,7], target = 7,
// A solution set is:
// [
//   [7],
//   [2,2,3]
// ]
//
// Example 2:
// Input: candidates = [2,3,5], target = 8,
// A solution set is:
// [
//   [2,2,2,2],
//   [2,3,3],
//   [3,5]
// ]

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}

	ans := make([][]int, 0)
	doCombination(candidates, target, 0, nil, &ans)
	return ans
}

func doCombination(candidates []int, target int, prevSum int, prev []int, ans *[][]int) {
	if prevSum == target {
		*ans = append(*ans, append([]int{}, prev...))
		return
	}

	if prevSum > target {
		return
	}

	for i, v := range candidates {
		prev = append(prev, v)
		doCombination(candidates[i:], target, prevSum+v, prev, ans)
		prev = prev[:len(prev)-1]
	}
}
