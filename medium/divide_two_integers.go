package medium

import "math"

// Given two integers dividend and divisor, divide two integers without
// using multiplication, division and mod operator.
//
// Return the quotient after dividing dividend by divisor.
//
// The integer division should truncate toward zero.
//
// Example 1:
// Input: dividend = 10, divisor = 3
// Output: 3
//
// Example 2:
// Input: dividend = 7, divisor = -3
// Output: -2
//
// Note:
// Both dividend and divisor will be 32-bit signed integers.
// The divisor will never be 0.
// Assume we are dealing with an environment which could only store
// integers within the 32-bit signed integer range: [−2^31,  2^31 − 1].
// For the purpose of this problem, assume that your function returns
// 2^31 − 1 when the division result overflows.

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	sign := 1
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		sign = -1
	}

	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	var quotient int
	for dividend-divisor >= 0 {
		quotient++
		dividend -= divisor
	}

	if sign == 1 {
		if quotient >= math.MaxInt32 {
			quotient = math.MaxInt32
		}
		return quotient
	} else {
		return -quotient
	}
}
