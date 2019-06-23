package hard

// Given a set of non-overlapping intervals, insert a new interval
// into the intervals (merge if necessary).
//
// You may assume that the intervals were initially sorted according
// to their start times.
//
// Example 1:
// Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
// Output: [[1,5],[6,9]]
//
// Example 2:
// Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
// Output: [[1,2],[3,10],[12,16]]
// Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		if len(newInterval) != 0 {
			return append(intervals, newInterval)
		} else {
			return nil
		}
	}

	startIndex := -1
	for i := 0; i < len(intervals); i++ {
		if newInterval[0] <= intervals[i][0] {
			startIndex = i
			break
		}
	}

	if startIndex == -1 {
		intervals = append(intervals, newInterval)
	} else {
		intervals = append(append(append([][]int{},
			intervals[0:startIndex]...),
			newInterval),
			intervals[startIndex:]...)
	}

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= intervals[i-1][1] {
			intervals[i-1][1] = maxInt(intervals[i][1], intervals[i-1][1])
			intervals = append(intervals[0:i], intervals[i+1:]...)
			i--
		}
	}
	return intervals
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
