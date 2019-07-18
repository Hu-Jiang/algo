package medium

// Suppose an array sorted in ascending order is rotated at some
// pivot unknown to you beforehand.
//
// (i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).
//
// You are given a target value to search. If found in the array
// return true, otherwise return false.
//
// Example 1:
// Input: nums = [2,5,6,0,0,1,2], target = 0
// Output:
//
// Example 2:
// Input: nums = [2,5,6,0,0,1,2], target = 3
// Output: false
//
// Follow up:
// This is a follow up problem to Search in Rotated Sorted Array,
// where nums may contain duplicates.
// Would this affect the run-time complexity? How and why?

func searchII(nums []int, target int) bool {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return true
		}
		if nums[low] < nums[mid] {
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[low] == nums[mid] {
			low++
		} else {
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return false
}
