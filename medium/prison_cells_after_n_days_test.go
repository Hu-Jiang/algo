package medium

import (
	"reflect"
	"testing"
)

func TestPrisonAfterNDays(t *testing.T) {
	tests := []struct {
		cells []int
		N     int
		want  []int
	}{
		{
			cells: []int{0, 1, 0, 1, 1, 0, 0, 1},
			N:     7,
			want:  []int{0, 0, 1, 1, 0, 0, 0, 0},
		},
		{
			cells: []int{1, 0, 0, 1, 0, 0, 1, 0},
			N:     1000000000,
			want:  []int{0, 0, 1, 1, 1, 1, 1, 0},
		},
	}

	for i, tt := range tests {
		got := prisonAfterNDays(tt.cells, tt.N)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
