package hard

import (
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{
			root: NewPreorderTree(-3),
			want: -3,
		},
		{
			root: NewPreorderTree(1, 2, -1, -1, 3),
			want: 6,
		},
		{
			root: NewPreorderTree(-10, 9, -1, -1, 20, 15, -1, -1, 7),
			want: 42,
		},
	}

	for i, tt := range tests {
		got := maxPathSum(tt.root)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
