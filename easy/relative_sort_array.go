package easy

import "sort"

// Given two arrays arr1 and arr2, the elements of arr2 are distinct,
// and all elements in arr2 are also in arr1.
//
// Sort the elements of arr1 such that the relative ordering of items
// in arr1 are the same as in arr2.  Elements that don't appear in arr2
// should be placed at the end of arr1 in ascending order.
//
// Example 1:
// Input: arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
// Output: [2,2,2,1,4,3,3,9,6,7,19]

// Constraints:
// arr1.length, arr2.length <= 1000
// 0 <= arr1[i], arr2[i] <= 1000
// Each arr2[i] is distinct.
// Each arr2[i] is in arr1.
func relativeSortArray(arr1 []int, arr2 []int) []int {
	var res []int
	var cnt int
	for i := 0; i < len(arr2); i++ {
		arr1, cnt = remove(arr1, arr2[i])
		for j := 0; j < cnt; j++ {
			res = append(res, arr2[i])
		}
	}
	sort.Ints(arr1)
	res = append(res, arr1...)
	return res
}

func remove(arr []int, n int) (res []int, cnt int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == n {
			arr = append(arr[0:i], arr[i+1:]...)
			i--
			cnt++
		}
	}
	return arr, cnt
}
