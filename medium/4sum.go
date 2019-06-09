package medium

import (
	"sort"
)

// Given an array nums of n integers and an integer target, are there elements
// a, b, c, and d in nums such that a + b + c + d = target? Find all unique
// quadruplets in the array which gives the sum of target.
//
// Note:
// The solution set must not contain duplicate quadruplets.
//
// Example:
// Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.
// A solution set is:
// [
//   [-1,  0, 0, 1],
//   [-2, -1, 1, 2],
//   [-2,  0, 0, 2]
// ]

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	sort.Ints(nums)

	var ans [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		threeSum := _threeSum(nums[i+1:], target-nums[i])
		if threeSum != nil {
			for j := 0; j < len(threeSum); j++ {
				var tmp []int
				tmp = append(tmp, nums[i])
				tmp = append(tmp, threeSum[j]...)
				ans = append(ans, tmp)
			}
		}
	}

	return ans
}

func _threeSum(nums []int, target int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	var res [][]int
	l := len(nums)
	for i := 0; i < l-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, l-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				for ; j < k && nums[j] == nums[j-1]; j++ {
				}
			} else if sum < target {
				j++
			} else {
				k--
			}
		}
	}

	return res
}
