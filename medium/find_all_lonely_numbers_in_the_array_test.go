package medium

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindLonely(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{10, 6, 5, 8}, []int{10, 8}},
		{[]int{1, 3, 5, 3}, []int{1, 5}},
	}

	for i, tt := range tests {
		got := findLonely(tt.input)
		sort.Ints(tt.want)
		sort.Ints(got)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got: %v, want: %v", i, got, tt.want)
		}
	}
}
