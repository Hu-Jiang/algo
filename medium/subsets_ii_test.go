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
				{},
			},
		},
		{
			nums: []int{1},
			want: [][]int{
				{1},
				{},
			},
		},
		{
			nums: []int{1, 2, 2},
			want: [][]int{
				{1, 2, 2},
				{1, 2},
				{1},
				{2, 2},
				{2},
				{},
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
