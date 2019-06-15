package medium

import (
	"reflect"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		want       [][]int
	}{
		{nil, 0, nil},
		{[]int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
		{[]int{2, 5, 2, 1, 2}, 5, [][]int{{1, 2, 2}, {5}}},
	}

	for i, tt := range tests {
		got := combinationSum2(tt.candidates, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
