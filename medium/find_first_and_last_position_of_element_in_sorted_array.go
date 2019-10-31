package medium

// Given an array of integers nums sorted in ascending order, find the
// starting and ending position of a given target value.
//
// Your algorithm's runtime complexity must be in the order of O(log n).
//
// If the target is not found in the array, return [-1, -1].
//
// Example 1:
// Input: nums = [5,7,7,8,8,10], target = 8
// Output: [3,4]
//
// Example 2:
// Input: nums = [5,7,7,8,8,10], target = 6
// Output: [-1,-1]

func searchRange(nums []int, target int) []int {
	left := binarySearchLeft(nums, target)
	if left == -1 {
		return []int{-1, -1}
	}

	right := binarySearchRight(nums, target)

	return []int{left, right}
}

func binarySearchLeft(nums []int, target int) (index int) {
	if len(nums) == 0 {
		return -1
	}

	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			high = mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if nums[low] == target {
		return low
	}
	return -1
}

func binarySearchRight(nums []int, target int) (index int) {
	if len(nums) == 0 {
		return -1
	}

	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if nums[high] == target {
		return high
	}
	return -1
}

// The official solution:
//
// Approach 1: Linear Scan
//
// class Solution {
//     public int[] searchRange(int[] nums, int target) {
//         int[] targetRange = {-1, -1};
//
//         // find the index of the leftmost appearance of `target`.
//         for (int i = 0; i < nums.length; i++) {
//             if (nums[i] == target) {
//                 targetRange[0] = i;
//                 break;
//             }
//         }
//
//         // if the last loop did not find any index, then there is no valid range
//         // and we return [-1, -1].
//         if (targetRange[0] == -1) {
//             return targetRange;
//         }
//
//         // find the index of the rightmost appearance of `target` (by reverse
//         // iteration). it is guaranteed to appear.
//         for (int j = nums.length-1; j >= 0; j--) {
//             if (nums[j] == target) {
//                 targetRange[1] = j;
//                 break;
//             }
//         }
//
//         return targetRange;
//     }
// }
//
// Approach 2: Binary Search
//
// class Solution {
//     // returns leftmost (or rightmost) index at which `target` should be
//     // inserted in sorted array `nums` via binary search.
//     private int extremeInsertionIndex(int[] nums, int target, boolean left) {
//         int lo = 0;
//         int hi = nums.length;
//
//         while (lo < hi) {
//             int mid = (lo + hi) / 2;
//             if (nums[mid] > target || (left && target == nums[mid])) {
//                 hi = mid;
//             }
//             else {
//                 lo = mid+1;
//             }
//         }
//
//         return lo;
//     }
//
//     public int[] searchRange(int[] nums, int target) {
//         int[] targetRange = {-1, -1};
//
//         int leftIdx = extremeInsertionIndex(nums, target, true);
//
//         // assert that `leftIdx` is within the array bounds and that `target`
//         // is actually in `nums`.
//         if (leftIdx == nums.length || nums[leftIdx] != target) {
//             return targetRange;
//         }
//
//         targetRange[0] = leftIdx;
//         targetRange[1] = extremeInsertionIndex(nums, target, false)-1;
//
//         return targetRange;
//     }
// }
