package easy

import (
	"testing"
)

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want bool
	}{
		{
			root: NewPreorderTree(3, 9, -1, -1, 20, 15, -1, -1, 7),
			want: true,
		},
		{
			root: NewPreorderTree(1, 2, 3, 4, -1, -1, 4, -1, -1, 3, -1, -1, 2),
			want: false,
		},
	}

	for i, tt := range tests {
		got := isBalanced(tt.root)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
