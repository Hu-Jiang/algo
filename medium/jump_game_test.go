package medium

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		// {nil, true},
		{[]int{1}, true},
		{[]int{0}, true},
		{[]int{0, 1}, false},
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
	}

	for i, tt := range tests {
		got := canJump(tt.nums)
		if got != tt.want {
			t.Fatalf("%d. got %v, want %v", i, got, tt.want)
		}
	}
}
