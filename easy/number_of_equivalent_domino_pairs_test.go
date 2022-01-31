package easy

import "testing"

func TestNumEquivDominoPairs(t *testing.T) {
	tests := []struct {
		dominoes [][]int
		want     int
	}{
		{
			nil, 0,
		},
		{
			dominoes: [][]int{
				{1, 2},
				{2, 1},
				{3, 4},
				{5, 6},
			},
			want: 1,
		},
		{
			dominoes: [][]int{
				{1, 2},
				{1, 2},
				{1, 1},
				{1, 2},
				{2, 2},
			},
			want: 3,
		},
	}

	for i, tt := range tests {
		got := numEquivDominoPairs(tt.dominoes)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
