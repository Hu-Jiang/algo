package easy

import (
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	tests := []struct {
		rowIndex int
		want     []int
	}{
		{
			rowIndex: -1,
			want:     nil,
		},
		{
			rowIndex: 0,
			want:     []int{1},
		},
		{
			rowIndex: 3,
			want:     []int{1, 3, 3, 1},
		},
	}

	for i, tt := range tests {
		got := getRow(tt.rowIndex)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
