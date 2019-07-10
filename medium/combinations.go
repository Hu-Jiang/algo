package medium

// Given two integers n and k, return all possible combinations
// of k numbers out of 1 ... n.
//
// Example:
// Input: n = 4, k = 2
// Output:
// [
//   [2,4],
//   [3,4],
//   [2,3],
//   [1,2],
//   [1,3],
//   [1,4],
// ]

func combine(n int, k int) [][]int {
	var res [][]int
	for i := 1; i <= n-k+1; i++ {
		doCombine(&res, []int{i}, n, k)
	}
	return res
}

func doCombine(res *[][]int, prev []int, n, k int) {
	if len(prev) == k {
		*res = append(*res, prev)
		return
	}
	if prev[len(prev)-1] == n {
		return
	}
	// remain: n-v
	// have: len(prev)
	// want: k-len(prev)
	// need ensure: remain >= want
	for v := prev[len(prev)-1]; v < n && n-v >= k-len(prev); v++ {
		doCombine(res, append(append([]int{}, prev...), v+1), n, k)
	}
}
