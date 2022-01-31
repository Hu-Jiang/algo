package easy

import "testing"

func TestFindFinalValue(t *testing.T) {
	tests := []struct {
		nums     []int
		original int
		want     int
	}{
		{[]int{5, 3, 6, 1, 12}, 3, 24},
		{[]int{2, 7, 9}, 4, 4},
	}

	for i, tt := range tests {
		got := findFinalValue(tt.nums, tt.original)
		if got != tt.want {
			t.Fatalf("case %d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
