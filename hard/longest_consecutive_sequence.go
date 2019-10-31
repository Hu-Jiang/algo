package hard

// Given an unsorted array of integers, find the length of the
// longest consecutive elements sequence.
//
// Your algorithm should run in O(n) complexity.
//
// Example:
// Input: [100, 4, 200, 1, 3, 2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4].
// Therefore its length is 4.

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		set[nums[i]] = true
	}

	var res int
	for i := 0; i < len(nums); i++ {
		if !set[nums[i]-1] {
			next := nums[i] + 1
			for set[next] {
				next++
			}
			res = maxConsecutive(res, next-nums[i])
		}
	}
	return res
}

func maxConsecutive(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// The official solution:
//
// Approach 1: Brute Force
//
// class Solution {
//     private boolean arrayContains(int[] arr, int num) {
//         for (int i = 0; i < arr.length; i++) {
//             if (arr[i] == num) {
//                 return true;
//             }
//         }
//
//         return false;
//     }
//     public int longestConsecutive(int[] nums) {
//         int longestStreak = 0;
//
//         for (int num : nums) {
//             int currentNum = num;
//             int currentStreak = 1;
//
//             while (arrayContains(nums, currentNum + 1)) {
//                 currentNum += 1;
//                 currentStreak += 1;
//             }
//
//             longestStreak = Math.max(longestStreak, currentStreak);
//         }
//
//         return longestStreak;
//     }
// }
//
// Approach 2: Sorting
//
// class Solution {
//     public int longestConsecutive(int[] nums) {
//         if (nums.length == 0) {
//             return 0;
//         }
//
//         Arrays.sort(nums);
//
//         int longestStreak = 1;
//         int currentStreak = 1;
//
//         for (int i = 1; i < nums.length; i++) {
//             if (nums[i] != nums[i-1]) {
//                 if (nums[i] == nums[i-1]+1) {
//                     currentStreak += 1;
//                 }
//                 else {
//                     longestStreak = Math.max(longestStreak, currentStreak);
//                     currentStreak = 1;
//                 }
//             }
//         }
//
//         return Math.max(longestStreak, currentStreak);
//     }
// }
//
// Approach 3: HashSet and Intelligent Sequence Building
//
// class Solution {
//     public int longestConsecutive(int[] nums) {
//         Set<Integer> num_set = new HashSet<Integer>();
//         for (int num : nums) {
//             num_set.add(num);
//         }
//
//         int longestStreak = 0;
//
//         for (int num : num_set) {
//             if (!num_set.contains(num-1)) {
//                 int currentNum = num;
//                 int currentStreak = 1;
//
//                 while (num_set.contains(currentNum+1)) {
//                     currentNum += 1;
//                     currentStreak += 1;
//                 }
//
//                 longestStreak = Math.max(longestStreak, currentStreak);
//             }
//         }
//
//         return longestStreak;
//     }
// }
