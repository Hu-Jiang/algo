package medium

import "testing"

func TestDivide(t *testing.T) {
	tests := []struct {
		divident int
		divisor  int
		want     int
	}{
		{10, 3, 3},
		{7, -3, -2},
		{12, 2, 6},
		{-2147483648, -1, 2147483647},
	}

	for i, tt := range tests {
		got := divide(tt.divident, tt.divisor)
		if got != tt.want {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
