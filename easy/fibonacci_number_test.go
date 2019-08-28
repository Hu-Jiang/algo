package easy

import "testing"

func TestFib(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
	}

	for i, tt := range tests {
		got := fib(tt.n)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %d", i, got, tt.want)
		}
	}
}
