package medium

import (
	"reflect"
	"testing"
)

func TestZigzagLevelOrder(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want [][]int
	}{
		{nil, nil},
		{
			NewPreorderTree(3, 9, -1, -1, 20, 15, -1, -1, 7),
			[][]int{
				{3},
				{20, 9},
				{15, 7},
			},
		},
	}

	for i, tt := range tests {
		got := zigzagLevelOrder(tt.root)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
