package medium

import (
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want []int
	}{
		{
			root: nil,
			want: nil,
		},
		{
			root: NewPreorderTree(1),
			want: []int{1},
		},
		{
			root: NewPreorderTree(1, -1, 2, 3),
			want: []int{1, 2, 3},
		},
		{
			root: NewPreorderTree(3, 1, -1, -1, 2),
			want: []int{3, 1, 2},
		},
	}

	for i, tt := range tests {
		got := preorderTraversal(tt.root)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
