package medium

import "sort"

// Given a collection of intervals, merge all overlapping intervals.
//
// Example 1:
//
// Input: [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
//
// Example 2:
// Input: [[1,4],[4,5]]
// Output: [[1,5]]
// Explanation: Intervals [1,4] and [4,5] are considered overlapping.

func merge(intervals [][]int) [][]int {
	sort.Sort(Intervals(intervals))
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= intervals[i-1][1] {
			intervals[i-1][1] = max(intervals[i][1], intervals[i-1][1])
			intervals = append(intervals[0:i], intervals[i+1:]...)
			i--
		}
	}
	return intervals
}

type Intervals [][]int

func (p Intervals) Len() int           { return len(p) }
func (p Intervals) Less(i, j int) bool { return p[i][0] < p[j][0] }
func (p Intervals) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// The official solution:
//
// Approach 1: Connected Components
// class Solution {
//     private Map<Interval, List<Interval> > graph;
//     private Map<Integer, List<Interval> > nodesInComp;
//     private Set<Interval> visited;
//
//     // return whether two intervals overlap (inclusive)
//     private boolean overlap(Interval a, Interval b) {
//         return a.start <= b.end && b.start <= a.end;
//     }
//
//     // build a graph where an undirected edge between intervals u and v exists
//     // iff u and v overlap.
//     private void buildGraph(List<Interval> intervals) {
//         graph = new HashMap<>();
//         for (Interval interval : intervals) {
//             graph.put(interval, new LinkedList<>());
//         }
//
//         for (Interval interval1 : intervals) {
//             for (Interval interval2 : intervals) {
//                 if (overlap(interval1, interval2)) {
//                     graph.get(interval1).add(interval2);
//                     graph.get(interval2).add(interval1);
//                 }
//             }
//         }
//     }
//
//     // merges all of the nodes in this connected component into one interval.
//     private Interval mergeNodes(List<Interval> nodes) {
//         int minStart = nodes.get(0).start;
//         for (Interval node : nodes) {
//             minStart = Math.min(minStart, node.start);
//         }
//
//         int maxEnd = nodes.get(0).end;
//         for (Interval node : nodes) {
//             maxEnd= Math.max(maxEnd, node.end);
//         }
//
//         return new Interval(minStart, maxEnd);
//     }
//
//     // use depth-first search to mark all nodes in the same connected component
//     // with the same integer.
//     private void markComponentDFS(Interval start, int compNumber) {
//         Stack<Interval> stack = new Stack<>();
//         stack.add(start);
//
//         while (!stack.isEmpty()) {
//             Interval node = stack.pop();
//             if (!visited.contains(node)) {
//                 visited.add(node);
//
//                 if (nodesInComp.get(compNumber) == null) {
//                     nodesInComp.put(compNumber, new LinkedList<>());
//                 }
//                 nodesInComp.get(compNumber).add(node);
//
//                 for (Interval child : graph.get(node)) {
//                     stack.add(child);
//                 }
//             }
//         }
//     }
//
//     // gets the connected components of the interval overlap graph.
//     private void buildComponents(List<Interval> intervals) {
//         nodesInComp = new HashMap();
//         visited = new HashSet();
//         int compNumber = 0;
//
//         for (Interval interval : intervals) {
//             if (!visited.contains(interval)) {
//                 markComponentDFS(interval, compNumber);
//                 compNumber++;
//             }
//         }
//     }
//
//     public List<Interval> merge(List<Interval> intervals) {
//         buildGraph(intervals);
//         buildComponents(intervals);
//
//         // for each component, merge all intervals into one interval.
//         List<Interval> merged = new LinkedList<>();
//         for (int comp = 0; comp < nodesInComp.size(); comp++) {
//             merged.add(mergeNodes(nodesInComp.get(comp)));
//         }
//
//         return merged;
//     }
// }
//
// Approach 2: Sorting
//
// class Solution {
//     private class IntervalComparator implements Comparator<Interval> {
//         @Override
//         public int compare(Interval a, Interval b) {
//             return a.start < b.start ? -1 : a.start == b.start ? 0 : 1;
//         }
//     }
//
//     public List<Interval> merge(List<Interval> intervals) {
//         Collections.sort(intervals, new IntervalComparator());
//
//         LinkedList<Interval> merged = new LinkedList<Interval>();
//         for (Interval interval : intervals) {
//             // if the list of merged intervals is empty or if the current
//             // interval does not overlap with the previous, simply append it.
//             if (merged.isEmpty() || merged.getLast().end < interval.start) {
//                 merged.add(interval);
//             }
//             // otherwise, there is overlap, so we merge the current and previous
//             // intervals.
//             else {
//                 merged.getLast().end = Math.max(merged.getLast().end, interval.end);
//             }
//         }
//
//         return merged;
//     }
// }
