package hard

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{nil, 0},
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
	}

	for i, tt := range tests {
		got := trap(tt.height)
		if got != tt.want {
			t.Fatalf("%d. got %v, want %v", i, got, tt.want)
		}
	}
}
