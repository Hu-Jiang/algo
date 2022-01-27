package easy

import "testing"

func TestCountElements(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{11, 7, 2, 15}, 2},
		{[]int{-3, 3, 3, 90}, 2},
		{[]int{3, 3, 3, 90}, 0},
		{[]int{3, 3, 3, 3, 4, 90}, 1},
	}

	for i, tt := range tests {
		got := countElements(tt.input)
		if got != tt.want {
			t.Fatalf("case %d: got %v; want: %v", i, got, tt.want)
		}
	}
}
