package medium

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	tests := []struct {
		root, want *TreeNode
	}{
		{
			root: nil,
			want: nil,
		},
		{
			root: NewPreorderTree(1, 2, 3, -1, -1, 4, 5, -1, 6),
			want: NewPreorderTree(1, -1, 2, -1, 3, -1, 4, -1, 5, -1, 6),
		},
	}

	for i, tt := range tests {
		flatten(tt.root)
		if !reflect.DeepEqual(tt.root, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, tt.root, tt.want)
		}
	}
}
