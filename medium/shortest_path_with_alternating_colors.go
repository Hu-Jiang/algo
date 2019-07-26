package medium

import (
	"math"
)

// Consider a directed graph, with nodes labelled 0, 1, ..., n-1.
// In this graph, each edge is either red or blue, and there could
// be self-edges or parallel edges.
//
// Each [i, j] in red_edges denotes a red directed edge from node i
// to node j.  Similarly, each [i, j] in blue_edges denotes a blue
// directed edge from node i to node j.
//
// Return an array answer of length n, where each answer[X] is the
// length of the shortest path from node 0 to node X such that the
// edge colors alternate along the path (or -1 if such a path doesn't
// exist).
//
// Example 1:
// Input: n = 3, red_edges = [[0,1],[1,2]], blue_edges = []
// Output: [0,1,-1]
//
// Example 2:
// Input: n = 3, red_edges = [[0,1]], blue_edges = [[2,1]]
// Output: [0,1,-1]
//
// Example 3:
// Input: n = 3, red_edges = [[1,0]], blue_edges = [[2,1]]
// Output: [0,-1,-1]
//
// Example 4:
// Input: n = 3, red_edges = [[0,1]], blue_edges = [[1,2]]
// Output: [0,1,2]
//
// Example 5:
// Input: n = 3, red_edges = [[0,1],[0,2]], blue_edges = [[1,0]]
// Output: [0,1,1]
//
// Constraints:
// 1 <= n <= 100
// red_edges.length <= 400
// blue_edges.length <= 400
// red_edges[i].length == blue_edges[i].length == 2
// 0 <= red_edges[i][j], blue_edges[i][j] < n

type color int

const (
	Red color = 1 << iota
	Blue
)

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	if n == 0 {
		return nil
	}

	graph := make([][]color, n)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]color, n)
	}
	for i := 0; i < len(redEdges); i++ {
		graph[redEdges[i][0]][redEdges[i][1]] |= Red
	}
	for i := 0; i < len(blueEdges); i++ {
		graph[blueEdges[i][0]][blueEdges[i][1]] |= Blue
	}

	result := []int{0}
	for i := 1; i < n; i++ {
		result = append(result, math.MaxInt32)
	}
	bfs(graph, result)
	for i := 1; i < n; i++ {
		if result[i] == math.MaxInt32 {
			result[i] = -1
		}
	}

	return result
}

// an incomeNodeColor represents which color of edge incoming this node
type incomeNodeColor struct {
	node int
	red  bool
	blue bool
}

func bfs(graph [][]color, res []int) {
	if len(graph) == 0 {
		return
	}

	var q []incomeNodeColor
	q = append(q, incomeNodeColor{node: 0, red: true, blue: true})
	visited := make([]color, len(graph))
	path := 1
	for len(q) > 0 {
		curItems := len(q)
		for i := 0; i < curItems; i++ {
			// dequeue
			inc := q[0]
			q = q[1:]

			for i := 0; i < len(graph); i++ {
				if graph[inc.node][i] != 0 { // have edge
					if inc.blue { // need red edge
						if graph[inc.node][i]&Red != 0 && (visited[i]&Red) == 0 { // have red edge and has not been visited
							q = append(q, incomeNodeColor{node: i, red: true})
							res[i] = minPath(res[i], path)
							visited[i] |= Red
						}
					}
					if inc.red { // need blue edge
						if graph[inc.node][i]&Blue != 0 && (visited[i]&Blue) == 0 { // have blue edge and has not been visited
							q = append(q, incomeNodeColor{node: i, blue: true})
							res[i] = minPath(res[i], path)
							visited[i] |= Blue
						}
					}
				}
			}
		}
		path++
	}
}

func minPath(a, b int) int {
	if a == -1 {
		return b
	}
	if b == -1 {
		return a
	}
	if a <= b {
		return a
	}
	return b
}
