package medium

// Given an array of non-negative integers, you are initially positioned at the
// first index of the array.
//
// Each element in the array represents your maximum jump length at that position.
//
// Determine if you are able to reach the last index.
//
// Example 1:
// Input: [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
//
// Example 2:
// Input: [3,2,1,0,4]
// Output: false
// Explanation: You will always arrive at index 3 no matter what. Its maximum
// 			 jump length is 0, which makes it impossible to reach the last index.

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	maxIndex := make([]int, len(nums))
	maxIndex[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		maxIndex[i] = maxInt(maxIndex[i-1], nums[i]+i)
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			if maxIndex[i] <= i {
				return false
			}
		}
	}
	return true
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// The official soloution:
//
// Approach 1: Backtracking
//
// public class Solution {
//     public boolean canJumpFromPosition(int position, int[] nums) {
//         if (position == nums.length - 1) {
//             return true;
//         }
//
//         int furthestJump = Math.min(position + nums[position], nums.length - 1);
//         for (int nextPosition = position + 1; nextPosition <= furthestJump; nextPosition++) {
//             if (canJumpFromPosition(nextPosition, nums)) {
//                 return true;
//             }
//         }
//
//         return false;
//     }
//
//     public boolean canJump(int[] nums) {
//         return canJumpFromPosition(0, nums);
//     }
// }
//
// Approach 2: Dynamic Programming Top-down
//
// enum Index {
//     GOOD, BAD, UNKNOWN
// }
//
// public class Solution {
//     Index[] memo;
//
//     public boolean canJumpFromPosition(int position, int[] nums) {
//         if (memo[position] != Index.UNKNOWN) {
//             return memo[position] == Index.GOOD ? true : false;
//         }
//
//         int furthestJump = Math.min(position + nums[position], nums.length - 1);
//         for (int nextPosition = position + 1; nextPosition <= furthestJump; nextPosition++) {
//             if (canJumpFromPosition(nextPosition, nums)) {
//                 memo[position] = Index.GOOD;
//                 return true;
//             }
//         }
//
//         memo[position] = Index.BAD;
//         return false;
//     }
//
//     public boolean canJump(int[] nums) {
//         memo = new Index[nums.length];
//         for (int i = 0; i < memo.length; i++) {
//             memo[i] = Index.UNKNOWN;
//         }
//         memo[memo.length - 1] = Index.GOOD;
//         return canJumpFromPosition(0, nums);
//     }
// }
//
// Approach 3: Dynamic Programming Bottom-up
//
// enum Index {
//     GOOD, BAD, UNKNOWN
// }
//
// public class Solution {
//     public boolean canJump(int[] nums) {
//         Index[] memo = new Index[nums.length];
//         for (int i = 0; i < memo.length; i++) {
//             memo[i] = Index.UNKNOWN;
//         }
//         memo[memo.length - 1] = Index.GOOD;
//
//         for (int i = nums.length - 2; i >= 0; i--) {
//             int furthestJump = Math.min(i + nums[i], nums.length - 1);
//             for (int j = i + 1; j <= furthestJump; j++) {
//                 if (memo[j] == Index.GOOD) {
//                     memo[i] = Index.GOOD;
//                     break;
//                 }
//             }
//         }
//
//         return memo[0] == Index.GOOD;
//     }
// }
//
// Approach 4: Greedy
//
// public class Solution {
//     public boolean canJump(int[] nums) {
//         int lastPos = nums.length - 1;
//         for (int i = nums.length - 1; i >= 0; i--) {
//             if (i + nums[i] >= lastPos) {
//                 lastPos = i;
//             }
//         }
//         return lastPos == 0;
//     }
// }
