package medium

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{nil, nil},
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 3, 2}, []int{2, 1, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
	}

	for i, tt := range tests {
		nextPermutation(tt.nums)
		if !reflect.DeepEqual(tt.nums, tt.want) {
			t.Fatalf("%d: got: %v; want: %v", i, tt.nums, tt.want)
		}
	}
}
