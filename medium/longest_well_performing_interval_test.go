package medium

import "testing"

func TestLongestWPI(t *testing.T) {
	tests := []struct {
		hours []int
		want  int
	}{
		{nil, 0},
		{[]int{1, 1, 1}, 0},
		{[]int{9, 1, 1}, 1},
		{[]int{9, 9, 6, 0, 6, 6, 9}, 3},
	}

	for i, tt := range tests {
		got := longestWPI(tt.hours)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
