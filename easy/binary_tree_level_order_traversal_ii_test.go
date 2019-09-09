package easy

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
				{15, 7},
				{9, 20},
				{3},
			},
		},
	}

	for i, tt := range tests {
		got := levelOrderBottom(tt.root)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
