package easy

import (
	"reflect"
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
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
	}

	for i, tt := range tests {
		got := maxProfit(tt.prices)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
