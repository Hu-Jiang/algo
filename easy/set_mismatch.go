package easy

// The set S originally contains numbers from 1 to n.
// But unfortunately, due to the data error, one of
// the numbers in the set got duplicated to another
// number in the set, which results in repetition of
// one number and loss of another number.
//
// Given an array nums representing the data status
// of this set after the error. Your task is to firstly
// find the number occurs twice and then find the number
// that is missing. Return them in the form of an array.
//
// Example 1:
// Input: nums = [1,2,2,4]
// Output: [2,3]
//
// Note:
// The given array size will in the range [2, 10000].
// The given array's numbers won't have any order.

func findErrorNums(nums []int) []int {
	var ans []int
	appear := make(map[int]bool)
	for _, v := range nums {
		if _, ok := appear[v]; ok {
			ans = append(ans, v)
		} else {
			appear[v] = true
		}
	}
	n := len(nums)
	for n > 0 {
		if !appear[n] {
			ans = append(ans, n)
		}
		n--
	}
	return ans
}

// The official soloution:
//
// Approach 1: Brute Force
//
// public class Solution {
//     public int[] findErrorNums(int[] nums) {
//         int dup = -1, missing = -1;
//         for (int i = 1; i <= nums.length; i++) {
//             int count = 0;
//             for (int j = 0; j < nums.length; j++) {
//                 if (nums[j] == i)
//                     count++;
//             }
//             if (count == 2)
//                 dup = i;
//             else if (count == 0)
//                 missing = i;
//         }
//         return new int[] {dup, missing};
//     }
// }
//
// Approach 2: Better Brute Force
//
// public class Solution {
//     public int[] findErrorNums(int[] nums) {
//         int dup = -1, missing = -1;;
//         for (int i = 1; i <= nums.length; i++) {
//             int count = 0;
//             for (int j = 0; j < nums.length; j++) {
//                 if (nums[j] == i)
//                     count++;
//             }
//             if (count == 2)
//                 dup = i;
//             else if (count == 0)
//                 missing = i;
//             if (dup > 0 && missing > 0)
//                 break;
//         }
//         return new int[] {dup, missing};
//     }
// }
//
// Approach 3: Using Sorting
//
// public class Solution {
//     public int[] findErrorNums(int[] nums) {
//         Arrays.sort(nums);
//         int dup = -1, missing = 1;
//         for (int i = 1; i < nums.length; i++) {
//             if (nums[i] == nums[i - 1])
//                 dup = nums[i];
//             else if (nums[i] > nums[i - 1] + 1)
//                 missing = nums[i - 1] + 1;
//         }
//         return new int[] {dup, nums[nums.length - 1] != nums.length ? nums.length : missing};
//     }
// }
//
// Approach 4: Using Map
//
// public class Solution {
//     public int[] findErrorNums(int[] nums) {
//         Map < Integer, Integer > map = new HashMap();
//         int dup = -1, missing = 1;
//         for (int n: nums) {
//             map.put(n, map.getOrDefault(n, 0) + 1);
//         }
//         for (int i = 1; i <= nums.length; i++) {
//             if (map.containsKey(i)) {
//                 if (map.get(i) == 2)
//                     dup = i;
//             } else
//                 missing = i;
//         }
//         return new int[]{dup, missing};
//     }
// }
//
// Approach 5: Using Extra Array
//
// public class Solution {
//     public int[] findErrorNums(int[] nums) {
//         int[] arr = new int[nums.length + 1];
//         int dup = -1, missing = 1;
//         for (int i = 0; i < nums.length; i++) {
//             arr[nums[i]] += 1;
//         }
//         for (int i = 1; i < arr.length; i++) {
//             if (arr[i] == 0)
//                 missing = i;
//             else if (arr[i] == 2)
//                 dup = i;
//         }
//         return new int[]{dup, missing};
//     }
// }
//
// Approach 6: Using Constant Space
//
// public class Solution {
//     public int[] findErrorNums(int[] nums) {
//         int dup = -1, missing = 1;
//         for (int n: nums) {
//             if (nums[Math.abs(n) - 1] < 0)
//                 dup = Math.abs(n);
//             else
//                 nums[Math.abs(n) - 1] *= -1;
//         }
//         for (int i = 1; i < nums.length; i++) {
//             if (nums[i] > 0)
//                 missing = i + 1;
//         }
//         return new int[]{dup, missing};
//     }
// }
//
// Approach 7: Using XOR
//
// public class Solution {
//     public int[] findErrorNums(int[] nums) {
//         int xor = 0, xor0 = 0, xor1 = 0;
//         for (int n: nums)
//             xor ^= n;
//         for (int i = 1; i <= nums.length; i++)
//             xor ^= i;
//         int rightmostbit = xor & ~(xor - 1);
//         for (int n: nums) {
//             if ((n & rightmostbit) != 0)
//                 xor1 ^= n;
//             else
//                 xor0 ^= n;
//         }
//         for (int i = 1; i <= nums.length; i++) {
//             if ((i & rightmostbit) != 0)
//                 xor1 ^= i;
//             else
//                 xor0 ^= i;
//         }
//         for (int i = 0; i < nums.length; i++) {
//             if (nums[i] == xor0)
//                 return new int[]{xor0, xor1};
//         }
//         return new int[]{xor1, xor0};
//     }
// }
