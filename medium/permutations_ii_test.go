package medium

import (
	"reflect"
	"testing"
)

func TestPermuteUniqu(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{nil, nil},
		{[]int{1, 1, 2}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
	}

	for i, tt := range tests {
		got := permuteUnique(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
