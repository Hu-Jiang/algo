package medium

import (
	"math"
	"sort"
)

// Given an array nums of n integers and an integer target, find three integers
// in nums such that the sum is closest to target. Return the sum of the three integers.
// You may assume that each input would have exactly one solution.
//
// Example:
// Given array nums = [-1, 2, 1, -4], and target = 1.
// The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums)

	closest := nums[0] + nums[1] + nums[2]
	eaual := false
	for i := 0; !eaual && i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if math.Abs(float64(sum-target)) < math.Abs(float64(closest-target)) {
				closest = sum
			}

			if sum == target {
				eaual = true
				break
			} else if sum < target {
				l++
			} else {
				r--
			}
		}
	}

	return closest
}
