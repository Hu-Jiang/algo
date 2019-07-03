package easy

import "testing"

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{nil, 0},
		{NewPreorderTree(1), 1},
		{NewPreorderTree(3, 9, -1, -1, 20, 15, -1, -1, 7), 3},
	}

	for i, tt := range tests {
		got := maxDepth(tt.root)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
