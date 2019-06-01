package medium

import (
	"reflect"
	"sort"
)

// Given an array nums of n integers, are there elements a, b, c in nums
// such that a + b + c = 0? Find all unique triplets in the array which
// gives the sum of zero.
//
// Note:
//
// The solution set must not contain duplicate triplets.
//
// Example:
//
// Given array nums = [-1, 0, 1, 2, -1, -4],
//
// A solution set is:
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] > 0 {
			return res
		}
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, l-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				for ; j < k && nums[j] == nums[j-1]; j++ {
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}

	return res
}

func threeSum_v2(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int
	l := len(nums)
	for i := 0; i < l; {
		if nums[i] > 0 {
			return res
		}
		for j := i + 1; j < l; {
			if nums[i]+nums[j] > 0 {
				j = l
				break
			}
			for k := j + 1; k < l; {
				sum := nums[i] + nums[j] + nums[k]
				if sum > 0 {
					k = l
					break
				}
				if sum == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}

				for k = k + 1; k < l; k++ {
					if nums[k] != nums[k-1] {
						break
					}
				}
			}

			for j = j + 1; j < l; j++ {
				if nums[j] != nums[j-1] {
					break
				}
			}

		}

		for i = i + 1; i < l; i++ {
			if nums[i] != nums[i-1] {
				break
			}
		}
	}

	return res
}

func threeSum_v1(nums []int) [][]int {
	var res [][]int

	len := len(nums)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			for k := j + 1; k < len; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					elem := []int{nums[i], nums[j], nums[k]}
					sort.Ints(elem)
					if !hasElement(res, elem) {
						res = append(res, elem)
					}
				}
			}
		}
	}

	return res
}

func hasElement(res [][]int, elem []int) bool {
	for _, v := range res {
		if reflect.DeepEqual(v, elem) {
			return true
		}
	}
	return false
}
