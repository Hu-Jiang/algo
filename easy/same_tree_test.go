package easy

import "testing"

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		p, q *TreeNode
		want bool
	}{
		{
			nil, nil, true,
		},
		{
			p:    NewPreorderTree(1, 2, -1, -1, 3),
			q:    NewPreorderTree(1, 2, -1, -1, 3, -1, -1),
			want: true,
		},
		{
			p:    NewPreorderTree(1, 2),
			q:    NewPreorderTree(1, -1, 2),
			want: false,
		},
		{
			p:    NewPreorderTree(1, 2, -1, -1, 1),
			q:    NewPreorderTree(1, 1, -1, -1, 2),
			want: false,
		},
	}

	for i, tt := range tests {
		got := isSameTree(tt.p, tt.q)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
