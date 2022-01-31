package hard

import (
	"reflect"
	"testing"
)

func TestSolveNQueues(t *testing.T) {
	tests := []struct {
		n    int
		want [][]string
	}{
		{0, [][]string{}},
		{1, [][]string{{"Q"}}},
		{2, [][]string{}},
		{3, [][]string{}},
		{4, [][]string{
			{
				".Q..",
				"...Q",
				"Q...",
				"..Q.",
			},
			{
				"..Q.",
				"Q...",
				"...Q",
				".Q..",
			},
		}},
	}

	for i, tt := range tests {
		got := solveNQueens(tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
