package medium

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      [][]int
	}{
		{nil, nil},
		{
			[][]int{
				{1, 4},
				{2, 3},
			},
			[][]int{
				{1, 4},
			},
		},
		{
			[][]int{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			},
			[][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			[][]int{
				{1, 4},
				{4, 5},
			},
			[][]int{
				{1, 5},
			},
		},
	}

	for i, tt := range tests {
		got := merge(tt.intervals)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%d. got: %v; want: %v", i, got, tt.want)
		}
	}
}
