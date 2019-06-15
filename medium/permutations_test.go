package medium

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{nil, nil},
		{[]int{1, 2, 3}, [][]int{
			{1, 2, 3}, {1, 3, 2},
			{2, 1, 3}, {2, 3, 1},
			{3, 1, 2}, {3, 2, 1}},
		},
	}

	for i, tt := range tests {
		got := permute(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}

}
