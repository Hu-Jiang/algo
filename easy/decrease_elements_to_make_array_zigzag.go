package easy

// Given an array nums of integers, a move consists of choosing any
// element and decreasing it by 1.
//
// An array A is a zigzag array if either:
// Every even-indexed element is greater than adjacent elements,
// ie. A[0] > A[1] < A[2] > A[3] < A[4] > ...
// OR, every odd-indexed element is greater than adjacent elements,
// ie. A[0] < A[1] > A[2] < A[3] > A[4] < ...
//
// Return the minimum number of moves to transform the given array nums
// into a zigzag array.
//
// Example 1:
// Input: nums = [1,2,3]
// Output: 2
// Explanation: We can decrease 2 to 0 or 3 to 1.
//
// Example 2:
// Input: nums = [9,6,1,6,2]
// Output: 4
//
// Constraints:
// 1 <= nums.length <= 1000
// 1 <= nums[i] <= 1000

func movesToMakeZigzag(nums []int) int {
	type Parity int

	const (
		even Parity = 0
		odd         = 1
	)

	decSum := func(nums []int, p Parity) (decSum int) {
		for i := 0; i < len(nums); i++ {
			if i%2 == int(p) {
				if i-1 >= 0 && nums[i-1] >= nums[i] {
					dec := nums[i-1] - nums[i] + 1
					decSum += dec
					nums[i-1] -= dec
				}
				if i+1 < len(nums) && nums[i+1] >= nums[i] {
					dec := nums[i+1] - nums[i] + 1
					decSum += dec
					nums[i+1] -= dec
				}
			}
		}
		return
	}

	tmp := make([]int, len(nums))
	copy(tmp, nums)
	evenDecSum := decSum(tmp, even)

	copy(tmp, nums)
	oddDecSum := decSum(tmp, odd)

	if evenDecSum > oddDecSum {
		return oddDecSum
	}
	return evenDecSum
}
