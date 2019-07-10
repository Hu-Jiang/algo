package medium

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{nil, nil},
		{[]int{1}, [][]int{[]int{}, []int{1}}},
		{[]int{1, 2, 3}, [][]int{
			[]int{},
			[]int{1},
			[]int{2},
			[]int{3},
			[]int{1, 2},
			[]int{1, 3},
			[]int{2, 3},
			[]int{1, 2, 3},
		}},
	}

	for i, tt := range tests {
		got := subsets(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
