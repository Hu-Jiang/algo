package medium

import (
	"reflect"
	"testing"
)

func TestMaxDepthAfterSplit(t *testing.T) {
	tests := []struct {
		seq  string
		want []int
	}{
		{"", []int{}},
		{"(()())", []int{0, 1, 0, 0, 1, 0}},
		{"()(())()", []int{0, 0, 1, 0, 1, 0, 1, 1}},
	}

	for i, tt := range tests {
		got := maxDepthAfterSplit(tt.seq)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
