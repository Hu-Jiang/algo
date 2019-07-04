package medium

import (
	"reflect"
	"testing"
)

func TestBuildTrees(t *testing.T) {
	tests := []struct {
		inorder, postorder []int
		want               *TreeNode
	}{
		{nil, nil, nil},
		{
			[]int{9, 3, 15, 20, 7},
			[]int{9, 15, 7, 20, 3},
			NewPreorderTree(3, 9, -1, -1, 20, 15, -1, -1, 7),
		},
		{
			[]int{10, 9, 3, 15, 20, 7},
			[]int{10, 9, 15, 7, 20, 3},
			NewPreorderTree(3, 9, 10, -1, -1, -1, 20, 15, -1, -1, 7),
		},
		{
			[]int{9, 10, 3, 15, 20, 7},
			[]int{10, 9, 15, 7, 20, 3},
			NewPreorderTree(3, 9, -1, 10, -1, -1, 20, 15, -1, -1, 7),
		},
	}

	for i, tt := range tests {
		got := buildTrees(tt.inorder, tt.postorder)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
