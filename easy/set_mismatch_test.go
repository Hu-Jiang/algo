package easy

import (
	"reflect"
	"testing"
)

func TestFindErrorNums(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{nil, nil},
		{[]int{1}, nil},
		{[]int{1, 2, 2, 4}, []int{2, 3}},
	}

	for i, tt := range tests {
		got := findErrorNums(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
