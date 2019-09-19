package medium

import (
	"reflect"
	"testing"
)

func TestPathSum(t *testing.T) {
	tests := []struct {
		root *TreeNode
		num  int
		want [][]int
	}{
		{
			root: nil,
			num:  100,
			want: nil,
		},
		{
			root: NewPreorderTree(5, 4, 11, 7, -1, -1, 2, -1, -1, -1, 8, 13, -1, -1, 4, 5, -1, -1, 1),
			num:  22,
			want: [][]int{
				[]int{5, 4, 11, 2},
				[]int{5, 8, 4, 5},
			},
		},
		{
			root: NewPreorderTree(5, 4, 11, 7, -1, -1, 2, -1, -1, -1, 8, 13, -1, -1, 4, -1, 1),
			num:  9,
			want: nil,
		},
		{
			root: NewPreorderTree(-2, -1, -3),
			num:  -5,
			want: [][]int{
				[]int{-2, -3},
			},
		},
	}

	for i, tt := range tests {
		got := pathSum(tt.root, tt.num)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
