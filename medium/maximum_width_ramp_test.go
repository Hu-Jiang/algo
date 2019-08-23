package medium

import "testing"

func TestMaxWidthRamp(t *testing.T) {
	tests := []struct {
		A    []int
		want int
	}{
		{[]int{6, 0, 8, 2, 1, 5}, 4},
		{[]int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}, 7},
	}

	for i, tt := range tests {
		got := maxWidthRamp(tt.A)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
