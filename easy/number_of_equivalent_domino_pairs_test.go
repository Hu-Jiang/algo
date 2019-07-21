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
				[]int{1, 2},
				[]int{2, 1},
				[]int{3, 4},
				[]int{5, 6},
			},
			want: 1,
		},
		{
			dominoes: [][]int{
				[]int{1, 2},
				[]int{1, 2},
				[]int{1, 1},
				[]int{1, 2},
				[]int{2, 2},
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
