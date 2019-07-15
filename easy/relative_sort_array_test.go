package easy

import (
	"reflect"
	"testing"
)

func TestRelativeSortArray(t *testing.T) {
	tests := []struct {
		arr1, arr2 []int
		want       []int
	}{
		{nil, nil, nil},
		{
			[]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
			[]int{2, 1, 4, 3, 9, 6},
			[]int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
	}

	for i, tt := range tests {
		got := relativeSortArray(tt.arr1, tt.arr2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
