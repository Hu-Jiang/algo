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
				[]int{1, 4},
				[]int{2, 3},
			},
			[][]int{
				[]int{1, 4},
			},
		},
		{
			[][]int{
				[]int{1, 3},
				[]int{2, 6},
				[]int{8, 10},
				[]int{15, 18},
			},
			[][]int{
				[]int{1, 6},
				[]int{8, 10},
				[]int{15, 18},
			},
		},
		{
			[][]int{
				[]int{1, 4},
				[]int{4, 5},
			},
			[][]int{
				[]int{1, 5},
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
