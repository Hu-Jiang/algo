package easy

// Given a non-negative integer num, repeatedly add all its digits
// until the result has only one digit.
//
// Example:
//
// Input: 38
// Output: 2
// Explanation: The process is like: 3 + 8 = 11, 1 + 1 = 2.
//              Since 2 has only one digit, return it.
// Follow up:
// Could you do it without any loop/recursion in O(1) runtime?

func addDigits(num int) int {
	ans := 0
	for num != 0 {
		ans += num % 10
		num /= 10
	}
	if ans < 10 {
		return ans
	}
	return addDigits(ans)
}
