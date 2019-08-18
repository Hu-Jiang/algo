package medium

import (
	"testing"
)

func TestMaxLevelSum(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{nil, 0},
		{
			NewPreorderTree(1, 7, 7, -1, -1, -8, -1, -1, 0),
			2,
		},
	}

	for i, tt := range tests {
		got := maxLevelSum(tt.root)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
