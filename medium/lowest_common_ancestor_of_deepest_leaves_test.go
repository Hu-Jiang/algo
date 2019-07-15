package medium

import (
	"reflect"
	"testing"
)

func TestLcaDeepestLeaves(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want *TreeNode
	}{
		{nil, nil},
		{NewPreorderTree(1, 2, -1, -1, 3), NewPreorderTree(1, 2, -1, -1, 3)},
		{NewPreorderTree(1, 2, 4, -1, -1, -1, 3), NewPreorderTree(4)},
		{NewPreorderTree(1, 2, 4, -1, -1, 5, -1, -1, 3), NewPreorderTree(2, 4, -1, -1, 5)},
	}

	for i, tt := range tests {
		got := lcaDeepestLeaves(tt.root)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
