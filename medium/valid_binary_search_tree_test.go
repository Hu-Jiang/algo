package medium

import (
	"testing"
)

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want bool
	}{
		{
			root: nil,
			want: true,
		},
		{
			root: NewPreorderTree(2147483647),
			want: true,
		},
		{
			root: NewPreorderTree(2, 1, -1, -1, 3),
			want: true,
		},
		{
			root: NewPreorderTree(5, 1, -1, -1, 4, 3, -1, -1, 6),
			want: false,
		},
		{
			root: NewPreorderTree(10, 5, -1, -1, 15, 6, -1, -1, 20),
			want: false,
		},
	}

	for i, tt := range tests {
		got := isValidBST(tt.root)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
