package hard

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{
			prices: nil,
			want:   0,
		},
		{
			prices: []int{3, 3, 5, 0, 0, 3, 1, 4},
			want:   6,
		},
		{
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
	}

	for i, tt := range tests {
		got := maxProfit(tt.prices)
		if got != tt.want {
			t.Errorf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
