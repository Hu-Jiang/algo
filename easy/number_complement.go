package easy

// Given a positive integer, output its complement number. The complement
// strategy is to flip the bits of its binary representation.
//
// Note:
// The given integer is guaranteed to fit within the range of a 32-bit signed integer.
// You could assume no leading zero bit in the integer’s binary representation.
//
// Example 1:
// Input: 5
// Output: 2
// Explanation: The binary representation of 5 is 101 (no leading zero bits), and its
// complement is 010. So you need to output 2.
//
// Example 2:
// Input: 1
// Output: 0
// Explanation: The binary representation of 1 is 1 (no leading zero bits), and its
// complement is 0. So you need to output 0.

func findComplement(num int) int {
	ans := 0
	flag := false // flag first 1 bit
	for i := uint(0); i < 32; i++ {
		if !flag && (((num << i) & 0x80000000) == 0x80000000) {
			flag = true
		}
		if flag {
			if ((num << i) & 0x80000000) == 0x80000000 {
				ans &= ^(1 << (31 - i))
			} else {
				ans |= 1 << (31 - i)
			}
		}
	}
	if flag == false { // num == 0
		return 1
	}

	return ans
}
