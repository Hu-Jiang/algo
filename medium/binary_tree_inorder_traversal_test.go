package medium

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want []int
	}{
		{nil, nil},
		{NewPreorderTree(1, -1, 2, 3), []int{1, 3, 2}},
	}

	for i, tt := range tests {
		got := inorderTraversal(tt.root)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
