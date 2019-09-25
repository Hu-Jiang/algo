package easy

import (
	"testing"
)

func TestMaxProfitII(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{
			prices: nil,
			want:   0,
		},
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   7,
		},
		{
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
	}

	for i, tt := range tests {
		got := maxProfitII(tt.prices)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
