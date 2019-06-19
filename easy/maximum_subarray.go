package easy

// Given an integer array nums, find the contiguous subarray
// (containing at least one number) which has the largest sum
// and return its sum.
//
// Example:
// Input: [-2,1,-3,4,-1,2,1,-5,4],
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
//
// Follow up:
// If you have figured out the O(n) solution, try coding another
// solution using the divide and conquer approach, which is more subtle.

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = maxInt(nums[i], nums[i]+dp[i-1])
	}

	return maxArrVal(dp)
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxArrVal(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
