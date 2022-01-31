package medium

import (
	"testing"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		obstacleGrid [][]int
		want         int
	}{
		{nil, 0},
		{
			[][]int{
				{1, 0},
			},
			0,
		},
		{
			[][]int{
				{0, 1, 0},
			},
			0,
		},
		{
			[][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
			3,
		},
		{
			[][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0}},
			2,
		},
		{
			[][]int{
				{0, 0, 0},
				{0, 0, 1},
				{0, 1, 0},
			},
			0,
		},
	}

	for i, tt := range tests {
		got := uniquePathsWithObstacles(tt.obstacleGrid)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
