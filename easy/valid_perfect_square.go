package easy

// Given a positive integer num, write a function which returns
// True if num is a perfect square else False.
//
// Note: Do not use any built-in library function such as sqrt.
//
// Example 1:
// Input: 16
// Output: true
//
// Example 2:
// Input: 14
// Output: false

func isPerfectSquare(num int) bool {
	var i int
	for i*i <= num {
		i++
	}

	return (i-1)*(i-1) == num
}

// Trick: think binary search
