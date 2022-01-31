package medium

// Given a set of distinct integers, nums, return all possible
// subsets (the power set).
//
// Note: The solution set must not contain duplicate subsets.
//
// Example:
// Input: nums = [1,2,3]
// Output:
// [
//   [3],
//   [1],
//   [2],
//   [1,2,3],
//   [1,3],
//   [2,3],
//   [1,2],
//   []
// ]

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	res := [][]int{{}}
	for k := 1; k <= len(nums); k++ {
		res = append(res, combineK(nums, k)...)
	}

	return res
}

func combineK(nums []int, k int) [][]int {
	var res [][]int
	for pos := 0; pos <= len(nums)-k; pos++ {
		doCombineK(&res, []int{nums[pos]}, nums, k, pos)
	}
	return res
}

func doCombineK(res *[][]int, prev, nums []int, k int, pos int) {
	if len(prev) == k {
		*res = append(*res, prev)
		return
	}
	if pos+1 == len(nums) {
		return
	}

	for pos = pos + 1; pos < len(nums); pos++ {
		doCombineK(res, append(append([]int{}, prev...), nums[pos]), nums, k, pos)
	}
}
