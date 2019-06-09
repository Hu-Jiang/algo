package medium

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height   []int
		wantArea int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 2, 3}, 2},
		{nil, 0},
	}

	for i, tt := range tests {
		got := maxArea(tt.height)
		if got != tt.wantArea {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.wantArea)
		}
	}
}

// The official soloution:
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
