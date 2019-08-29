package medium

import (
	"reflect"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{
			nums: nil,
			want: [][]int{
				[]int{},
			},
		},
		{
			nums: []int{1},
			want: [][]int{
				[]int{1},
				[]int{},
			},
		},
		{
			nums: []int{1, 2, 2},
			want: [][]int{
				[]int{1, 2, 2},
				[]int{1, 2},
				[]int{1},
				[]int{2, 2},
				[]int{2},
				[]int{},
			},
		},
	}

	for i, tt := range tests {
		got := subsetsWithDup(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
