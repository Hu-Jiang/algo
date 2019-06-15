package medium

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   [][]int
	}{
		{nil, nil},
		{[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}, [][]int{
			{7, 4, 1},
			{8, 5, 2},
			{9, 6, 3},
		}},
		{[][]int{
			{5, 1, 9, 11},
			{2, 4, 8, 10},
			{13, 3, 6, 7},
			{15, 14, 12, 16},
		}, [][]int{
			{15, 13, 2, 5},
			{14, 3, 4, 1},
			{12, 6, 8, 9},
			{16, 7, 10, 11},
		}},
	}

	for i, tt := range tests {
		rotate(tt.matrix)
		if !reflect.DeepEqual(tt.matrix, tt.want) {
			t.Fatalf("%d: got: %v; want: %v", i, tt.matrix, tt.want)
		}
	}
}
