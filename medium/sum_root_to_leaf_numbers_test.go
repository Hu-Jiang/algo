package medium

import "testing"

func TestSumNumbers(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{nil, 0},
		{NewPreorderTree(1, 2, -1, -1, 3), 25},
		{NewPreorderTree(4, 9, 5, -1, -1, 1, -1, -1, 0), 1026},
	}

	for i, tt := range tests {
		got := sumNumbers(tt.root)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
