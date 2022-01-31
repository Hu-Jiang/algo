package medium

import (
	"reflect"
	"testing"
)

func TestMaxScoreIndices(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{0, 0, 1, 0}, []int{2, 4}},
		{[]int{0, 0, 0}, []int{3}},
		{[]int{1, 1}, []int{0}},
	}

	for i, tt := range tests {
		got := maxScoreIndices(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("case %d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
