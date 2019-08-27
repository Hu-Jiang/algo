package medium

import (
	"reflect"
	"testing"
)

func TestGrayCode(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{0, []int{0}},
		{1, []int{0, 1}},
		{2, []int{0, 1, 3, 2}},
	}

	for i, tt := range tests {
		got := grayCode(tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
