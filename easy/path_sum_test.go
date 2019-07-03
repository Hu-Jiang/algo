package easy

import (
	"testing"
)

func TestPathSum(t *testing.T) {
	tests := []struct {
		root *TreeNode
		num  int
		want bool
	}{
		{nil, 100, false},
		{
			NewPreorderTree(5, 4, 11, 7, -1, -1, 2, -1, -1, -1, 8, 13, -1, -1, 4, -1, 1),
			22,
			true,
		},
		{
			NewPreorderTree(5, 4, 11, 7, -1, -1, 2, -1, -1, -1, 8, 13, -1, -1, 4, -1, 1),
			9,
			false,
		},
	}

	for i, tt := range tests {
		got := hasPathSum(tt.root, tt.num)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
