package medium

// Given a non-empty array of integers, every element appears three times
// except for one, which appears exactly once. Find that single one.
//
// Note:
// Your algorithm should have a linear runtime complexity. Could you implement
// it without using extra memory?
//
// Example 1:
// Input: [2,2,3,2]
// Output: 3
//
// Example 2:
// Input: [0,1,0,1,0,1,99]
// Output: 99

func singleNumber(nums []int) int {
	var res int32
	var i uint
	for ; i < 32; i++ {
		var cnt uint
		for k := 0; k < len(nums); k++ {
			cnt += (uint(nums[k]) >> i) & 1
		}
		if cnt%3 != 0 {
			res = res | (1 << i)
		}
	}
	return int(res)
}
