package easy

import (
	"testing"
)

func TestMySqrt(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{0, 0},
		{4, 2},
		{8, 2},
	}

	for i, tt := range tests {
		got := mySqrt(tt.x)
		if got != tt.want {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
