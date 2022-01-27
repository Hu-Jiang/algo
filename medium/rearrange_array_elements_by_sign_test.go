package medium

import (
	"reflect"
	"testing"
)

func TestReArrayArray(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{3, 1, -2, -5, 2, -4}, []int{3, -2, 1, -5, 2, -4}},
		{[]int{-1, 1}, []int{1, -1}},
	}

	for i, tt := range tests {
		got := rearrangeArray(tt.input)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("case %d: got %v; want: %v", i, got, tt.want)
		}
	}
}
