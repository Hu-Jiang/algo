package medium

import "testing"

func TestSearchMatrixII(t *testing.T) {
	tests := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{nil, 10, false},
		{[][]int{{1}}, 1, true},
		{
			[][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			5,
			true,
		},
		{
			[][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			20,
			false,
		},
	}

	for i, tt := range tests {
		got := searchMatrixII(tt.matrix, tt.target)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
