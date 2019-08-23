package easy

import (
	"testing"
)

func TestRepeatdNTimes(t *testing.T) {
	tests := []struct {
		A    []int
		want int
	}{
		{[]int{1, 2, 3, 3}, 3},
		{[]int{2, 1, 2, 5, 3, 2}, 2},
		{[]int{5, 1, 5, 2, 5, 3, 5, 4}, 5},
	}

	for i, tt := range tests {
		got := repeatedNTimes(tt.A)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
