package hard

import (
	"reflect"
	"testing"
)

func TestRecoverTree(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want *TreeNode
	}{
		{
			root: nil,
			want: nil,
		},
		{
			root: NewPreorderTree(1, 3, -1, 2),
			want: NewPreorderTree(3, 1, -1, 2),
		},
		{
			root: NewPreorderTree(3, 1, -1, -1, 4, 2),
			want: NewPreorderTree(2, 1, -1, -1, 4, 3),
		},
	}

	for i, tt := range tests {
		recoverTree(tt.root)
		if !reflect.DeepEqual(tt.root, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, tt.root, tt.want)
		}
	}
}
