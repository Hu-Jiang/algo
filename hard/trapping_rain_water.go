package hard

// Given n non-negative integers representing an elevation map
// where the width of each bar is 1, compute how much water it
// is able to trap after raining.
//
// Image location: [/image/hard/trapping_rain_water.png]
// The above elevation map is represented by array
// [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water
// (blue section) are being trapped. Thanks Marcos for contributing
// this image!
//
// Example:
// Input: [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	left := make([]int, len(height))
	left[0] = height[0]
	maxLeft := left[0]
	for i := 1; i < len(height); i++ {
		if height[i] > maxLeft {
			maxLeft = height[i]
		}
		left[i] = maxLeft
	}

	right := make([]int, len(height))
	right[len(height)-1] = height[len(height)-1]
	maxRight := right[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		if height[i] > maxRight {
			maxRight = height[i]
		}
		right[i] = maxRight
	}

	cap := 0
	for i := 0; i < len(height); i++ {
		cap += minHeight(left[i], right[i]) - height[i]
	}

	return cap
}

func minHeight(h1, h2 int) int {
	if h1 <= h2 {
		return h1
	}
	return h2
}

// The official solution:
//
// Approach 1: Brute force [no code]
//
// Approach 2: Dynamic Programming
//
// int trap(vector<int>& height)
// {
// 	if(height == null)
// 		return 0;
//     int ans = 0;
//     int size = height.size();
//     vector<int> left_max(size), right_max(size);
//     left_max[0] = height[0];
//     for (int i = 1; i < size; i++) {
//         left_max[i] = max(height[i], left_max[i - 1]);
//     }
//     right_max[size - 1] = height[size - 1];
//     for (int i = size - 2; i >= 0; i--) {
//         right_max[i] = max(height[i], right_max[i + 1]);
//     }
//     for (int i = 1; i < size - 1; i++) {
//         ans += min(left_max[i], right_max[i]) - height[i];
//     }
//     return ans;
// }
//
// Approach 3: Using stacks
//
// int trap(vector<int>& height)
// {
//     int ans = 0, current = 0;
//     stack<int> st;
//     while (current < height.size()) {
//         while (!st.empty() && height[current] > height[st.top()]) {
//             int top = st.top();
//             st.pop();
//             if (st.empty())
//                 break;
//             int distance = current - st.top() - 1;
//             int bounded_height = min(height[current], height[st.top()]) - height[top];
//             ans += distance * bounded_height;
//         }
//         st.push(current++);
//     }
//     return ans;
// }
//
// Approach 4: Using 2 pointers
//
// int trap(vector<int>& height)
// {
//     int left = 0, right = height.size() - 1;
//     int ans = 0;
//     int left_max = 0, right_max = 0;
//     while (left < right) {
//         if (height[left] < height[right]) {
//             height[left] >= left_max ? (left_max = height[left]) : ans += (left_max - height[left]);
//             ++left;
//         }
//         else {
//             height[right] >= right_max ? (right_max = height[right]) : ans += (right_max - height[right]);
//             --right;
//         }
//     }
//     return ans;
// }
