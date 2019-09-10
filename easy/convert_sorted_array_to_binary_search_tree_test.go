package easy

import (
	"reflect"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		nums []int
		want *TreeNode
	}{
		{
			nums: nil,
			want: nil,
		},
		{
			nums: []int{-10, -3, 0, 5, 9},
			want: NewPreorderTree(0, -3, -10, -1, -1, -1, 9, 5),
		},
		{
			nums: []int{0, 1, 2, 3, 4, 5},
			want: NewPreorderTree(3, 1, 0, -1, -1, 2, -1, -1, 5, 4),
		},
	}

	for i, tt := range tests {
		got := sortedArrayToBST(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
