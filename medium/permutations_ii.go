package medium

import "sort"

// Given a collection of numbers that might contain duplicates,
// return all possible unique permutations.
//
// Example:
// Input: [1,1,2]
// Output:
// [
//   [1,1,2],
//   [1,2,1],
//   [2,1,1]
// ]

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)
	ans := make([][]int, 0)
	backtrackUnique(nums, nil, &ans)

	return ans
}

func backtrackUnique(nums []int, prev []int, ans *[][]int) {
	if len(nums) == 0 {
		*ans = append(*ans, append([]int{}, prev...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		backtrackUnique(append(append([]int{}, nums[0:i]...), nums[i+1:]...),
			append(prev, nums[i]),
			ans)
	}
}
