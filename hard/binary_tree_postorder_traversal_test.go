package hard

import (
	"reflect"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want []int
	}{
		{
			root: nil,
			want: nil,
		},
		{
			root: NewPreorderTree(1),
			want: []int{1},
		},
		{
			root: NewPreorderTree(1, -1, 2, 3),
			want: []int{3, 2, 1},
		},
	}

	for i, tt := range tests {
		got := postorderTraversal(tt.root)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
