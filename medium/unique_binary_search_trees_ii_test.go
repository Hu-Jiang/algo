package medium

import (
	"reflect"
	"testing"
)

func TestGenerateTrees(t *testing.T) {
	tests := []struct {
		n    int
		want []*TreeNode
	}{
		{
			n:    0,
			want: nil,
		},
		{
			n: 1,
			want: []*TreeNode{
				NewPreorderTree(1),
			},
		},
		{
			n: 2,
			want: []*TreeNode{
				NewPreorderTree(1, -1, 2),
				NewPreorderTree(2, 1),
			},
		},
		{
			n: 3,
			want: []*TreeNode{
				NewPreorderTree(1, -1, 2, -1, 3),
				NewPreorderTree(1, -1, 3, 2),
				NewPreorderTree(2, 1, -1, -1, 3),
				NewPreorderTree(3, 1, -1, 2),
				NewPreorderTree(3, 2, 1),
			},
		},
	}

	for i, tt := range tests {
		got := generateTrees(tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
