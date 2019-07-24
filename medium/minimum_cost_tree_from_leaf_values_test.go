package medium

import "testing"

func TestMCTFromLeafValues(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{nil, 0},
		{[]int{3}, 3},
		{[]int{6, 2, 4}, 32},
	}

	for i, tt := range tests {
		got := mctFromLeafValues(tt.arr)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
