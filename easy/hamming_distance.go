package easy

// The Hamming distance between two integers is the number
// of positions at which the corresponding bits are different.
//
// Given two integers x and y, calculate the Hamming distance.
//
// Note:
// 0 ≤ x, y < 2^31.
//
// Example:
// Input: x = 1, y = 4
// Output: 2
// Explanation:
// 1   (0 0 0 1)
// 4   (0 1 0 0)
//        ↑   ↑
// The above arrows point to positions where the corresponding
// bits are different.

func hammingDistance(x int, y int) int {
	t := x ^ y
	var bits int
	for n := uint(0); n < 32; n++ {
		if (t>>n)&1 == 1 {
			bits++
		}
	}
	return bits
}
