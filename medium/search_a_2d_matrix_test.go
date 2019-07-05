package medium

import "testing"

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{nil, 0, false},
		{[][]int{[]int{}}, 0, false},
		{[][]int{[]int{1}}, 1, true},
		{
			[][]int{
				[]int{1, 3, 5, 7},
				[]int{10, 11, 16, 20},
				[]int{23, 30, 34, 50},
			},
			3,
			true,
		},
		{
			[][]int{
				[]int{1, 3, 5, 7},
				[]int{10, 11, 16, 20},
				[]int{23, 30, 34, 50},
			},
			13,
			false,
		},
	}

	for i, tt := range tests {
		got := searchMatrix(tt.matrix, tt.target)
		if got != tt.want {
			t.Fatalf("%d. got %v, want %v", i, got, tt.want)
		}
	}
}
