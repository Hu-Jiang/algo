package medium

// Implement next permutation, which rearranges numbers into the lexicographically
// next greater permutation of numbers.
//
// If such arrangement is not possible, it must rearrange it as the lowest possible
// order (ie, sorted in ascending order).
//
// The replacement must be in-place and use only constant extra memory.
//
// Here are some examples. Inputs are in the left-hand column and its corresponding
// outputs are in the right-hand column.
//
// 1,2,3 → 1,3,2
// 3,2,1 → 1,2,3
// 1,1,5 → 1,5,1

func nextPermutation(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}

	i := len(nums) - 1
	for ; i > 0; i-- {
		if nums[i] > nums[i-1] {
			break
		}
	}

	if i == 0 {
		for j, k := 0, len(nums)-1; j < k; j, k = j+1, k-1 {
			nums[j], nums[k] = nums[k], nums[j]
		}
		return
	}

	j := len(nums) - 1
	for ; j >= i; j-- {
		if nums[j] > nums[i-1] {
			break
		}
	}

	nums[i-1], nums[j] = nums[j], nums[i-1]

	for l, r := i, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}

// The official solution:
//
// Approach 1: Brute Force [no code]
//
// Approach 2: Single Pass Approach
//
// public class Solution {
//     public void nextPermutation(int[] nums) {
//         int i = nums.length - 2;
//         while (i >= 0 && nums[i + 1] <= nums[i]) {
//             i--;
//         }
//         if (i >= 0) {
//             int j = nums.length - 1;
//             while (j >= 0 && nums[j] <= nums[i]) {
//                 j--;
//             }
//             swap(nums, i, j);
//         }
//         reverse(nums, i + 1);
//     }
//
//     private void reverse(int[] nums, int start) {
//         int i = start, j = nums.length - 1;
//         while (i < j) {
//             swap(nums, i, j);
//             i++;
//             j--;
//         }
//     }
//
//     private void swap(int[] nums, int i, int j) {
//         int temp = nums[i];
//         nums[i] = nums[j];
//         nums[j] = temp;
//     }
// }
