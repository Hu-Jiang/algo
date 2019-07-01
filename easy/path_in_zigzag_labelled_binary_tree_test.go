package easy

import (
	"reflect"
	"testing"
)

func TestPathInZigZagTree(t *testing.T) {
	tests := []struct {
		label int
		want  []int
	}{
		{-1, nil},
		{0, nil},
		{1, []int{1}},
		{14, []int{1, 3, 4, 14}},
		{26, []int{1, 2, 6, 10, 26}},
	}

	for i, tt := range tests {
		got := pathInZigZagTree(tt.label)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%d. got %v, want %v", i, got, tt.want)
		}
	}
}
