package medium

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	tests := []struct {
		preorder, inorder []int
		want              *TreeNode
	}{
		{nil, nil, nil},
		{
			[]int{3, 9, 20, 15, 7},
			[]int{9, 3, 15, 20, 7},
			NewPreorderTree(3, 9, -1, -1, 20, 15, -1, -1, 7),
		},
		{
			[]int{3, 9, 10, 20, 15, 7},
			[]int{10, 9, 3, 15, 20, 7},
			NewPreorderTree(3, 9, 10, -1, -1, -1, 20, 15, -1, -1, 7),
		},
		{
			[]int{3, 9, 10, 20, 15, 7},
			[]int{9, 10, 3, 15, 20, 7},
			NewPreorderTree(3, 9, -1, 10, -1, -1, 20, 15, -1, -1, 7),
		},
	}

	for i, tt := range tests {
		got := buildTree(tt.preorder, tt.inorder)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
