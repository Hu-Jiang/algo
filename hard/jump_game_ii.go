package hard

// Given an array of non-negative integers, you are initially positioned
// at the first index of the array.
//
// Each element in the array represents your maximum jump length at that position.
//
// Your goal is to reach the last index in the minimum number of jumps.
//
// Example:
// Input: [2,3,1,1,4]
// Output: 2
// Explanation: The minimum number of jumps to reach the last index is 2.
//     Jump 1 step from index 0 to 1, then 3 steps to the last index.
// Note:
// You can assume that you can always reach the last index.

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	var maxIdx, i, step int
	for i < len(nums)-1 {
		nextIdx, tmpIdx := 0, 0
		for j := 0; j < nums[i]; j++ {
			tmpIdx = i + j + 1
			if tmpIdx < len(nums) && tmpIdx+nums[tmpIdx] > maxIdx {
				maxIdx = tmpIdx + nums[tmpIdx]
				nextIdx = tmpIdx
			}
			if tmpIdx == len(nums)-1 {
				nextIdx = tmpIdx
			}
		}

		i = nextIdx
		step++
	}

	return step
}
