package easy

// Given an array of integers, return indices of the two numbers
// such that they add up to a specific target.
//
// You may assume that each input would have exactly one solution,
// and you may not use the same element twice.
//
// Example:
// Given nums = [2, 7, 11, 15], target = 9,
//
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

func twoSum(nums []int, target int) (indices []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				indices = append(indices, i, j)
				return
			}
		}
	}
	return
}

// The official solution:
//
// Approach 1: Brute Force
//
// public int[] twoSum(int[] nums, int target) {
//     for (int i = 0; i < nums.length; i++) {
//         for (int j = i + 1; j < nums.length; j++) {
//             if (nums[j] == target - nums[i]) {
//                 return new int[] { i, j };
//             }
//         }
//     }
//     throw new IllegalArgumentException("No two sum solution");
// }
//
// Approach 2: Two-pass Hash Table
//
// public int[] twoSum(int[] nums, int target) {
//     Map<Integer, Integer> map = new HashMap<>();
//     for (int i = 0; i < nums.length; i++) {
//         map.put(nums[i], i);
//     }
//     for (int i = 0; i < nums.length; i++) {
//         int complement = target - nums[i];
//         if (map.containsKey(complement) && map.get(complement) != i) {
//             return new int[] { i, map.get(complement) };
//         }
//     }
//     throw new IllegalArgumentException("No two sum solution");
// }
//
// Approach 3: One-pass Hash Table
//
// public int[] twoSum(int[] nums, int target) {
//     Map<Integer, Integer> map = new HashMap<>();
//     for (int i = 0; i < nums.length; i++) {
//         int complement = target - nums[i];
//         if (map.containsKey(complement)) {
//             return new int[] { map.get(complement), i };
//         }
//         map.put(nums[i], i);
//     }
//     throw new IllegalArgumentException("No two sum solution");
// }
