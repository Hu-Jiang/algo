package medium

import (
	"reflect"
	"testing"
)

func TestTopKFreq(t *testing.T) {
	tests := []struct {
		nums []int
		n    int
		want []int
	}{
		{
			nums: nil,
			n:    0,
			want: nil,
		},
		{
			nums: []int{1},
			n:    -1,
			want: nil,
		},
		{
			nums: []int{1},
			n:    1,
			want: []int{1},
		},
		{
			nums: []int{1, 2, 2, 2, 2, 3, 3, 4, 5},
			n:    2,
			want: []int{2, 3},
		},
		{
			nums: []int{1, 2, 2, 3, 3, 3},
			n:    100,
			want: []int{3, 2, 1},
		},
	}

	for i, tt := range tests {
		got := topKFrequent(tt.nums, tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("case %d. got %v, want %v", i, got, tt.want)
		}
	}
}
