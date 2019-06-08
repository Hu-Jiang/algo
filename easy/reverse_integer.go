package easy

// Given a 32-bit signed integer, reverse digits of an integer.
//
// Example 1:
// Input: 123
// Output: 321
//
// Example 2:
// Input: -123
// Output: -321
//
// Example 3:
// Input: 120
// Output: 21
//
// Note:
// Assume we are dealing with an environment which could only store
// integers within the 32-bit signed integer range: [−2^31,  2^31 − 1].
// For the purpose of this problem, assume that your function returns 0
// when the reversed integer overflows.

func reverse(x int) int {
	return int(innerReverse(int32(x)))
}

func innerReverse(x int32) int32 {
	var result int64 = 0
	for x != 0 {
		result = (result * 10) + int64(x%10)
		x = x / 10
	}

	if int64(int32(result)) != result {
		return 0
	}
	return int32(result)
}
