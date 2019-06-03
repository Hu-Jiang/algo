package medium

// Implement pow(x, n), which calculates x raised to the power n (xn).
//
// Example 1:
// Input: 2.00000, 10
// Output: 1024.00000
//
// Example 2:
// Input: 2.10000, 3
// Output: 9.26100
//
// Example 3:
// Input: 2.00000, -2
// Output: 0.25000
// Explanation: 2^(-2) = 1/(2^2) = 1/4 = 0.25
//
// Note:
// -100.0 < x < 100.0
// n is a 32-bit signed integer, within the range [−2^31, 2^31 − 1]

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / pow(x, -n)
	} else {
		return pow(x, n)
	}
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	if n&1 == 0 { // even
		val := pow(x, n>>1)
		return val * val
	} else {
		val := pow(x, (n-1)>>1)
		return val * val * x
	}
}

func myPow_v1(x float64, n int) float64 {
	if n < 0 {
		return 1 / iterPow(x, -n)
	} else {
		return iterPow(x, n)
	}
}

func iterPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	var val float64 = 1
	for n > 0 {
		val *= x
		n--
	}
	return val
}
