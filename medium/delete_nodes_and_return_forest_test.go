package medium

import (
	"reflect"
	"testing"
)

func TestDelNodes(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		toDelete []int
		want     []*TreeNode
	}{
		{nil, nil, nil},
		{
			root:     NewPreorderTree(1, 2, 4, -1, -1, 5, -1, -1, 3, 6, -1, -1, 7),
			toDelete: []int{3, 5},
			want:     []*TreeNode{NewPreorderTree(1, 2, 4), NewPreorderTree(6), NewPreorderTree(7)},
		},
	}

	for i, tt := range tests {
		got := delNodes(tt.root, tt.toDelete)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
