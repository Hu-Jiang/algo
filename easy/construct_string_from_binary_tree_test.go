package easy

import "testing"

func TestTree2Str(t *testing.T) {
	tests := []struct {
		t    *TreeNode
		want string
	}{
		{nil, ""},
		{NewPreorderTree(1, 2, 4, -1, -1, -1, 3), "1(2(4))(3)"},
		{NewPreorderTree(1, 2, -1, 4, -1, -1, 3), "1(2()(4))(3)"},
	}

	for i, tt := range tests {
		got := tree2str(tt.t)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
