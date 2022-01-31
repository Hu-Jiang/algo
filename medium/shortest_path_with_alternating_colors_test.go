package medium

import (
	"reflect"
	"testing"
)

func TestShortestAlternatingPaths(t *testing.T) {
	tests := []struct {
		n         int
		redEdges  [][]int
		blueEdges [][]int
		want      []int
	}{
		{0, nil, nil, nil},
		{3, nil, nil, []int{0, -1, -1}},
		{
			n: 3,
			redEdges: [][]int{
				{0, 1},
				{1, 2},
			},
			blueEdges: [][]int{},
			want:      []int{0, 1, -1},
		},
		{
			n: 3,
			redEdges: [][]int{
				{0, 1},
			},
			blueEdges: [][]int{
				{2, 1},
			},
			want: []int{0, 1, -1},
		},
		{
			n: 3,
			redEdges: [][]int{
				{1, 0},
			},
			blueEdges: [][]int{
				{2, 1},
			},
			want: []int{0, -1, -1},
		},
		{
			n: 3,
			redEdges: [][]int{
				{0, 1},
				{0, 2},
			},
			blueEdges: [][]int{
				{1, 0},
			},
			want: []int{0, 1, 1},
		},
		{
			n: 4,
			redEdges: [][]int{
				{0, 1},
			},
			blueEdges: [][]int{
				{1, 1},
				{1, 2},
			},
			want: []int{0, 1, 2, -1},
		},
		{
			n: 5,
			redEdges: [][]int{
				{0, 1},
				{1, 2},
				{2, 3},
				{3, 4},
			},
			blueEdges: [][]int{
				{1, 2},
				{2, 3},
				{3, 1},
			},
			want: []int{0, 1, 2, 3, 7},
		},
		{
			n: 5,
			redEdges: [][]int{
				{2, 2},
				{0, 1},
				{0, 3},
				{0, 0},
				{0, 4},
				{2, 1},
				{2, 0},
				{1, 4},
				{3, 4},
			},
			blueEdges: [][]int{
				{1, 3},
				{0, 0},
				{0, 3},
				{4, 2},
				{1, 0},
			},
			want: []int{0, 1, 2, 1, 1},
		},
	}

	for i, tt := range tests {
		got := shortestAlternatingPaths(tt.n, tt.redEdges, tt.blueEdges)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
