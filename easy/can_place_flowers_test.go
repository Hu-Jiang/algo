package easy

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	tests := []struct {
		flowerbed []int
		n         int
		want      bool
	}{
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 2, false},
		{[]int{}, 0, true},
		{[]int{}, 1, false},
		{[]int{0}, 1, true},
		{[]int{1}, 1, false},
	}

	for i, tt := range tests {
		got := canPlaceFlowers(tt.flowerbed, tt.n)
		if got != tt.want {
			t.Fatalf("case %d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
