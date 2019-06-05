package hard

// There are two sorted arrays nums1 and nums2 of size m and n respectively.
// Find the median of the two sorted arrays. The overall run time complexity
// should be O(log (m+n)).
//
// You may assume nums1 and nums2 cannot be both empty.
//
// Example 1:
// nums1 = [1, 3]
// nums2 = [2]
//
// The median is 2.0
//
// Example 2:
// nums1 = [1, 2]
// nums2 = [3, 4]
//
// The median is (2 + 3)/2 = 2.5

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergeArr := merge(nums1, nums2)
	if len(mergeArr) == 0 {
		return 0
	}

	mid := len(mergeArr) / 2
	if len(mergeArr)%2 == 0 {
		return float64(mergeArr[mid-1]+mergeArr[mid]) / 2
	} else {
		return float64(mergeArr[mid])
	}
}

// merge returns sorted array merged by sorted array of nums1 and nums2.
func merge(nums1 []int, nums2 []int) []int {
	len1, len2 := len(nums1), len(nums2)
	i, j := 0, 0
	ans := make([]int, 0, len1+len2)
	for i < len1 && j < len2 {
		if nums1[i] < nums2[j] {
			ans = append(ans, nums1[i])
			i++
		} else {
			ans = append(ans, nums2[j])
			j++
		}
	}

	for i < len1 {
		ans = append(ans, nums1[i])
		i++
	}

	for j < len2 {
		ans = append(ans, nums2[j])
		j++
	}

	return ans
}

// The official soloution:
//
// Recursive Approach: Time complexity: O(log(min(m,n))).
//
// class Solution {
//     public double findMedianSortedArrays(int[] A, int[] B) {
//         int m = A.length;
//         int n = B.length;
//         if (m > n) { // to ensure m<=n
//             int[] temp = A; A = B; B = temp;
//             int tmp = m; m = n; n = tmp;
//         }
//         int iMin = 0, iMax = m, halfLen = (m + n + 1) / 2;
//         while (iMin <= iMax) {
//             int i = (iMin + iMax) / 2;
//             int j = halfLen - i;
//             if (i < iMax && B[j-1] > A[i]){
//                 iMin = i + 1; // i is too small
//             }
//             else if (i > iMin && A[i-1] > B[j]) {
//                 iMax = i - 1; // i is too big
//             }
//             else { // i is perfect
//                 int maxLeft = 0;
//                 if (i == 0) { maxLeft = B[j-1]; }
//                 else if (j == 0) { maxLeft = A[i-1]; }
//                 else { maxLeft = Math.max(A[i-1], B[j-1]); }
//                 if ( (m + n) % 2 == 1 ) { return maxLeft; }
//
//                 int minRight = 0;
//                 if (i == m) { minRight = B[j]; }
//                 else if (j == n) { minRight = A[i]; }
//                 else { minRight = Math.min(B[j], A[i]); }
//
//                 return (maxLeft + minRight) / 2.0;
//             }
//         }
//         return 0.0;
//     }
// }
