package easy

import "testing"

func TestSymmetricTree(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want bool
	}{
		{NewPreorderTree(1, 2, 3, -1, -1, 4, -1, -1, 2, 4, -1, -1, 3), true},
		{NewPreorderTree(1, 2, -1, 3, -1, -1, 2, -1, 3), false},
	}

	for i, tt := range tests {
		got := isSymmetric(tt.root)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
