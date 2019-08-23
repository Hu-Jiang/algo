package medium

import (
	"testing"
)

func TestIsCompleteTree(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want bool
	}{
		{NewPreorderTree(1, 2, 4, -1, -1, 5, -1, -1, 3, 6), true},
		{NewPreorderTree(1, 2, 4, -1, -1, 5, -1, -1, 3, -1, 7), false},
	}

	for i, tt := range tests {
		got := isCompleteTree(tt.root)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
