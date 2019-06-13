package medium

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{nil, 10, []int{-1, -1}},
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{1}, 1, []int{0, 0}},
		{[]int{1, 1}, 1, []int{0, 1}},
		{[]int{1, 1, 1, 1, 1}, 1, []int{0, 4}},
	}

	for i, tt := range tests {
		got := searchRange(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
