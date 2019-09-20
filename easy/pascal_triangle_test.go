package easy

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		numRows int
		want    [][]int
	}{
		{
			numRows: -1,
			want:    nil,
		},
		{
			numRows: 1,
			want: [][]int{
				[]int{1},
			},
		},
		{
			numRows: 5,
			want: [][]int{
				[]int{1},
				[]int{1, 1},
				[]int{1, 2, 1},
				[]int{1, 3, 3, 1},
				[]int{1, 4, 6, 4, 1},
			},
		},
	}

	for i, tt := range tests {
		got := generate(tt.numRows)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
