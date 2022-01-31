package medium

import (
	"reflect"
	"testing"
)

func TestCorpFlightBookings(t *testing.T) {
	tests := []struct {
		bookings [][]int
		n        int
		want     []int
	}{
		{nil, 0, nil},
		{
			bookings: [][]int{
				{1, 2, 10},
				{2, 3, 20},
				{2, 5, 25},
			},
			n:    5,
			want: []int{10, 55, 45, 25, 25},
		},
	}

	for i, tt := range tests {
		got := corpFlightBookings(tt.bookings, tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
