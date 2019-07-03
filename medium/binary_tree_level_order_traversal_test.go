package medium

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want [][]int
	}{
		{nil, nil},
		{
			NewPreorderTree(3, 9, -1, -1, 20, 15, -1, -1, 7),
			[][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
	}

	for i, tt := range tests {
		got := levelOrder(tt.root)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
