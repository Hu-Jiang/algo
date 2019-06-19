package medium

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   []int
	}{
		{nil, nil},
		{[][]int{[]int{1}}, []int{1}},
		{[][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9},
		}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{[][]int{
			[]int{1, 2, 3, 4},
			[]int{5, 6, 7, 8},
			[]int{9, 10, 11, 12},
		}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}

	for i, tt := range tests {
		got := spiralOrder(tt.matrix)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
