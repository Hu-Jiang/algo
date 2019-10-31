package medium

// Given n non-negative integers a1, a2, ..., an , where each represents
// a point at coordinate (i, ai). n vertical lines are drawn such that
// the two endpoints of line i is at (i, ai) and (i, 0). Find two lines,
// which together with x-axis forms a container, such that the container
// contains the most water.
//
// Note: You may not slant the container and n is at least 2.
//
// Image location: [/image/medium/container_with_most_water.jpg]
//
// The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7].
// In this case, the max area of water (blue section) the container can
// contain is 49.
//
// Example:
//
// Input: [1,8,6,2,5,4,8,3,7]
// Output: 49

func maxArea(height []int) int {
	var maxArea int
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			a := area(height[i], height[j], j-i)
			if a > maxArea {
				maxArea = a
			}
		}
	}
	return maxArea
}

func area(hl, hr, width int) int {
	if hl > hr {
		return hr * width
	} else {
		return hl * width
	}
}

// The official solution:
//
// Approach 1: Brute Force
//
// public class Solution {
//     public int maxArea(int[] height) {
//         int maxarea = 0;
//         for (int i = 0; i < height.length; i++)
//             for (int j = i + 1; j < height.length; j++)
//                 maxarea = Math.max(maxarea, Math.min(height[i], height[j]) * (j - i));
//         return maxarea;
//     }
// }
//
// Approach 2: Two Pointer Approach
//
// public class Solution {
//     public int maxArea(int[] height) {
//         int maxarea = 0, l = 0, r = height.length - 1;
//         while (l < r) {
//             maxarea = Math.max(maxarea, Math.min(height[l], height[r]) * (r - l));
//             if (height[l] < height[r])
//                 l++;
//             else
//                 r--;
//         }
//         return maxarea;
//     }
// }
