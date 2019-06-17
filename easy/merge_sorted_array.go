package easy

// Given two sorted integer arrays nums1 and nums2, merge nums2 into
// nums1 as one sorted array.
//
// Note:
// The number of elements initialized in nums1 and nums2 are m and n respectively.
// You may assume that nums1 has enough space (size that is greater or equal to m + n)
// to hold additional elements from nums2.
//
// Example:
// Input:
// nums1 = [1,2,3,0,0,0], m = 3
// nums2 = [2,5,6],       n = 3
// Output: [1,2,2,3,5,6]

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, len1 := 0, 0, m
	for i < len1 && j < n {
		n1 := nums1[i]
		n2 := nums2[j]
		if n1 >= n2 {
			for k := len(nums1) - 2; k >= i; k-- {
				nums1[k], nums1[k+1] = nums1[k+1], nums1[k]
			}
			nums1[i] = n2
			i++
			len1++
			j++
		} else {
			i++
		}
	}

	anchor := m + j
	for j < n {
		nums1[anchor] = nums2[j]
		anchor++
		j++
	}
}
