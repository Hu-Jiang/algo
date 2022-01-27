package hard

import "testing"

func TestMaximumGood(t *testing.T) {
	tests := []struct {
		input [][]int
		want  int
	}{
		{[][]int{
			{2, 1, 2},
			{1, 2, 2},
			{2, 0, 2},
		}, 2},
		{[][]int{
			{2, 0},
			{0, 2},
		}, 1},
	}

	for i, tt := range tests {
		got := maximumGood(tt.input)
		if got != tt.want {
			t.Fatalf("#%d. got: %d, want: %d", i, got, tt.want)
		}
	}
}
