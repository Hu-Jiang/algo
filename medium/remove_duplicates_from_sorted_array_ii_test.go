package medium

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums     []int
		wantNums []int
		wantCnt  int
	}{
		{nil, nil, 0},
		{[]int{1}, []int{1}, 1},
		{[]int{1, 1}, []int{1, 1}, 2},
		{[]int{1, 1, 1}, []int{1, 1}, 2},
		{[]int{1, 1, 1, 2, 2, 3}, []int{1, 1, 2, 2, 3}, 5},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, []int{0, 0, 1, 1, 2, 3, 3}, 7},
	}

	for i, tt := range tests {
		got := removeDuplicates(tt.nums)
		if got != tt.wantCnt || !reflect.DeepEqual(tt.nums[:got], tt.wantNums) {
			t.Fatalf("#%d. got {%v, %v}, want {%v, %v}", i, got, tt.nums[:got], tt.wantCnt, tt.wantNums)
		}
	}
}
