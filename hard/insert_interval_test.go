package hard

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{
			nil, nil, nil,
		},
		{
			nil,
			[]int{5, 7},
			[][]int{
				{5, 7},
			},
		},
		{
			[][]int{
				{1, 2}, {3, 4},
			},
			[]int{5, 6},
			[][]int{
				{1, 2}, {3, 4}, {5, 6},
			},
		},
		{
			[][]int{
				{3, 4}, {5, 6},
			},
			[]int{1, 2},
			[][]int{
				{1, 2}, {3, 4}, {5, 6},
			},
		},
		{
			[][]int{
				{1, 3}, {6, 9},
			},
			[]int{2, 5},
			[][]int{
				{1, 5}, {6, 9},
			},
		},
		{
			[][]int{
				{1, 5}, {10, 15},
			},
			[]int{8, 12},
			[][]int{
				{1, 5}, {8, 15},
			},
		},
		{
			[][]int{
				{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
			},
			[]int{4, 8},
			[][]int{
				{1, 2}, {3, 10}, {12, 16},
			},
		},
	}

	for i, tt := range tests {
		got := insert(tt.intervals, tt.newInterval)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%d. got: %v; want: %d", i, got, tt.want)
		}
	}
}
