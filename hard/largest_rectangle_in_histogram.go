package hard

// Given n non-negative integers representing the histogram's
// bar height where the width of each bar is 1, find the area
// of largest rectangle in the histogram.
//
// Image location: [/image/medium/largest_rectangle_in_histogram_1.png]
// Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].
//
// Image location: [/image/medium/largest_rectangle_in_histogram_2.png]
// The largest rectangle is shown in the shaded area, which has area = 10 unit.
//
// Example:
// Input: [2,1,5,6,2,3]
// Output: 10

func largestRectangleArea(heights []int) int {
	var stack []int
	stack = append(stack, -1)
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		for len(stack) != 1 && heights[i] <= heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
			maxArea = maxVal(maxArea, heights[top]*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}

	for len(stack) != 1 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		maxArea = maxVal(maxArea, heights[top]*(len(heights)-stack[len(stack)-1]-1))
	}

	return maxArea
}

func maxVal(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// The official soloution:
//
// 方法 1：暴力
//
// public class Solution {
// 	public int largestRectangleArea(int[] heights) {
// 		int maxarea = 0;
// 		for (int i = 0; i < heights.length; i++) {
// 			for (int j = i; j < heights.length; j++) {
// 				int minheight = Integer.MAX_VALUE;
// 				for (int k = i; k <= j; k++)
// 					minheight = Math.min(minheight, heights[k]);
// 				maxarea = Math.max(maxarea, minheight * (j - i + 1));
// 			}
// 		}
// 		return maxarea;
// 	}
// }
//
// 方法 2：优化的暴力
//
// public class Solution {
// 	public int largestRectangleArea(int[] heights) {
// 		int maxarea = 0;
// 		for (int i = 0; i < heights.length; i++) {
// 			int minheight = Integer.MAX_VALUE;
// 			for (int j = i; j < heights.length; j++) {
// 				minheight = Math.min(minheight, heights[j]);
// 				maxarea = Math.max(maxarea, minheight * (j - i + 1));
// 			}
// 		}
// 		return maxarea;
// 	}
// }
//
// 方法 3：分治
//
// public class Solution {
//     public int calculateArea(int[] heights, int start, int end) {
//         if (start > end)
//             return 0;
//         int minindex = start;
//         for (int i = start; i <= end; i++)
//             if (heights[minindex] > heights[i])
//                 minindex = i;
//         return Math.max(heights[minindex] * (end - start + 1), Math.max(calculateArea(heights, start, minindex - 1), calculateArea(heights, minindex + 1, end)));
//     }
//     public int largestRectangleArea(int[] heights) {
//         return calculateArea(heights, 0, heights.length - 1);
//     }
// }
//
// 方法 4：优化的分治
//
// class SegTreeNode {
// 	public:
// 	  int start;
// 	  int end;
// 	  int min;
// 	  SegTreeNode *left;
// 	  SegTreeNode *right;
// 	  SegTreeNode(int start, int end) {
// 		this->start = start;
// 		this->end = end;
// 		left = right = NULL;
// 	  }
// 	};
//
// 	class Solution {
// 	public:
// 	  int largestRectangleArea(vector<int>& heights) {
// 		if (heights.size() == 0) return 0;
// 		// first build a segment tree
// 		SegTreeNode *root = buildSegmentTree(heights, 0, heights.size() - 1);
// 		// next calculate the maximum area recursively
// 		return calculateMax(heights, root, 0, heights.size() - 1);
// 	  }
//
// 	  int calculateMax(vector<int>& heights, SegTreeNode* root, int start, int end) {
// 		if (start > end) {
// 		  return -1;
// 		}
// 		if (start == end) {
// 		  return heights[start];
// 		}
// 		int minIndex = query(root, heights, start, end);
// 		int leftMax = calculateMax(heights, root, start, minIndex - 1);
// 		int rightMax = calculateMax(heights, root, minIndex + 1, end);
// 		int minMax = heights[minIndex] * (end - start + 1);
// 		return max( max(leftMax, rightMax), minMax );
// 	  }
//
// 	  SegTreeNode *buildSegmentTree(vector<int>& heights, int start, int end) {
// 		if (start > end) return NULL;
// 		SegTreeNode *root = new SegTreeNode(start, end);
// 		if (start == end) {
// 			root->min = start;
// 		  return root;
// 		} else {
// 		  int middle = (start + end) / 2;
// 		  root->left = buildSegmentTree(heights, start, middle);
// 		  root->right = buildSegmentTree(heights, middle + 1, end);
// 		  root->min = heights[root->left->min] < heights[root->right->min] ? root->left->min : root->right->min;
// 		  return root;
// 		}
// 	  }
//
// 	  int query(SegTreeNode *root, vector<int>& heights, int start, int end) {
// 		if (root == NULL || end < root->start || start > root->end) return -1;
// 		if (start <= root->start && end >= root->end) {
// 		  return root->min;
// 		}
// 		int leftMin = query(root->left, heights, start, end);
// 		int rightMin = query(root->right, heights, start, end);
// 		if (leftMin == -1) return rightMin;
// 		if (rightMin == -1) return leftMin;
// 		return heights[leftMin] < heights[rightMin] ? leftMin : rightMin;
// 	  }
// 	};
//
// 方法 5：单调栈
//
// public class Solution {
//     public int largestRectangleArea(int[] heights) {
//         Stack < Integer > stack = new Stack < > ();
//         stack.push(-1);
//         int maxarea = 0;
//         for (int i = 0; i < heights.length; ++i) {
//             while (stack.peek() != -1 && heights[stack.peek()] >= heights[i])
//                 maxarea = Math.max(maxarea, heights[stack.pop()] * (i - stack.peek() - 1));
//             stack.push(i);
//         }
//         while (stack.peek() != -1)
//             maxarea = Math.max(maxarea, heights[stack.pop()] * (heights.length - stack.peek() -1));
//         return maxarea;
//     }
// }
