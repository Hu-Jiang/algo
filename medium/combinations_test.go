package medium

import (
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	tests := []struct {
		n, k int
		want [][]int
	}{
		{0, 0, nil},
		{5, 10, nil},
		{4, 2,
			[][]int{
				[]int{1, 2},
				[]int{1, 3},
				[]int{1, 4},
				[]int{2, 3},
				[]int{2, 4},
				[]int{3, 4},
			},
		},
	}

	for i, tt := range tests {
		got := combine(tt.n, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
