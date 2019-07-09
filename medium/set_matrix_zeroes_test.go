package medium

import (
	"reflect"
	"testing"
)

func TestSetZeros(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   [][]int
	}{
		{nil, nil},
		{
			[][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			[][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			[][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			[][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
	}

	for i, tt := range tests {
		setZeroes(tt.matrix)
		if !reflect.DeepEqual(tt.matrix, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, tt.matrix, tt.want)
		}
	}
}
