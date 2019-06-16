package easy

import (
	"reflect"
	"testing"
)

func TestDuplicateZeros(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{nil, nil},
		{[]int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
	}

	for i, tt := range tests {
		duplicateZeros(tt.arr)
		if !reflect.DeepEqual(tt.arr, tt.want) {
			t.Fatalf("%d: got: %v; want: %v", i, tt.arr, tt.want)
		}
	}
}
