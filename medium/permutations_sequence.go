package medium

import "strconv"

// The set [1,2,3,...,n] contains a total of n! unique permutations.
//
// By listing and labeling all of the permutations in order, we get
// the following sequence for n = 3:
// 1. "123"
// 2. "132"
// 3. "213"
// 4. "231"
// 5. "312"
// 6. "321"
//
// Given n and k, return the k-th permutation sequence.
//
// Note:
// Given n will be between 1 and 9 inclusive.
// Given k will be between 1 and n! inclusive.
//
// Example 1:
// Input: n = 3, k = 3
// Output: "213"
//
// Example 2:
// Input: n = 4, k = 9
// Output: "2314"

func getPermutation(n int, k int) string {
	var nums []int
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}

	var ans string
	var pos int
	permuteK(nums, &pos, k, "", &ans)

	return ans
}

func permuteK(nums []int, pos *int, k int, prev string, res *string) {
	if len(*res) != 0 {
		return
	}

	if len(nums) == 0 {
		*pos++
		if *pos == k {
			*res = prev
		}
		return
	}

	for i := 0; i < len(nums); i++ {
		permuteK(append(append([]int{}, nums[:i]...), nums[i+1:]...),
			pos, k, prev+strconv.Itoa(nums[i]), res)
	}
}
