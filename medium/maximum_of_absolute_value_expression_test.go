package medium

import "testing"

func TestMaxAbsValExpr(t *testing.T) {
	tests := []struct {
		arr1, arr2 []int
		want       int
	}{
		{nil, nil, 0},
		{[]int{1, 2, 3, 4}, []int{-1, 4, 5, 6}, 13},
		{[]int{1, -2, -5, 0, 10}, []int{0, -2, -1, -7, -4}, 20},
	}
	for i, tt := range tests {
		got := maxAbsValExpr(tt.arr1, tt.arr2)
		if got != tt.want {
			t.Errorf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
