package medium

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{0}, []int{0}},
		{[]int{2, 0}, []int{0, 2}},
		{[]int{0, 1, 2}, []int{0, 1, 2}},
		{[]int{2, 1, 0}, []int{0, 1, 2}},
		{[]int{2, 0, 1}, []int{0, 1, 2}},
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
	}

	for i, tt := range tests {
		sortColors(tt.nums)
		if !reflect.DeepEqual(tt.nums, tt.want) {
			t.Errorf("#%d. got %v, want %v", i, tt.nums, tt.want)
		}
	}
}
