package medium

import (
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	tests := []struct {
		triangle [][]int
		want     int
	}{
		{
			triangle: nil,
			want:     -1,
		},
		{
			triangle: [][]int{
				[]int{2},
				[]int{3, 4},
				[]int{6, 5, 7},
				[]int{4, 1, 8, 3},
			},
			want: 11,
		},
	}

	for i, tt := range tests {
		got := minimumTotal(tt.triangle)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
