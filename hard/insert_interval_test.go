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
				[]int{5, 7},
			},
		},
		{
			[][]int{
				[]int{1, 2}, []int{3, 4},
			},
			[]int{5, 6},
			[][]int{
				[]int{1, 2}, []int{3, 4}, []int{5, 6},
			},
		},
		{
			[][]int{
				[]int{3, 4}, []int{5, 6},
			},
			[]int{1, 2},
			[][]int{
				[]int{1, 2}, []int{3, 4}, []int{5, 6},
			},
		},
		{
			[][]int{
				[]int{1, 3}, []int{6, 9},
			},
			[]int{2, 5},
			[][]int{
				[]int{1, 5}, []int{6, 9},
			},
		},
		{
			[][]int{
				[]int{1, 5}, []int{10, 15},
			},
			[]int{8, 12},
			[][]int{
				[]int{1, 5}, []int{8, 15},
			},
		},
		{
			[][]int{
				[]int{1, 2}, []int{3, 5}, []int{6, 7}, []int{8, 10}, []int{12, 16},
			},
			[]int{4, 8},
			[][]int{
				[]int{1, 2}, []int{3, 10}, []int{12, 16},
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
