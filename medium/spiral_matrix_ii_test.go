package medium

import (
	"reflect"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	tests := []struct {
		n    int
		want [][]int
	}{
		{0, nil},
		{1, [][]int{[]int{1}}},
		{2, [][]int{
			[]int{1, 2},
			[]int{4, 3},
		}},
		{3, [][]int{
			[]int{1, 2, 3},
			[]int{8, 9, 4},
			[]int{7, 6, 5},
		}},
		{4, [][]int{
			[]int{1, 2, 3, 4},
			[]int{12, 13, 14, 5},
			[]int{11, 16, 15, 6},
			[]int{10, 9, 8, 7},
		}},
	}

	for i, tt := range tests {
		got := generateMatrix(tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
