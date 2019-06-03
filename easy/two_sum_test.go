package easy

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tester := []struct {
		arr     []int
		target  int
		wantIdx []int
	}{
		{
			arr:     []int{1, 2, 3, 4, 5},
			target:  7,
			wantIdx: []int{1, 4},
		},
		{
			arr:     []int{},
			target:  10,
			wantIdx: nil,
		},
		{
			arr:     []int{1, 2, 3},
			target:  100,
			wantIdx: nil,
		},
	}

	for i, test := range tester {
		gotIdx := twoSum(test.arr, test.target)
		if !reflect.DeepEqual(test.wantIdx, gotIdx) {
			t.Fatalf("test %d: arr: %v, target: %d, want: %v, got: %v", i, test.arr, test.target, test.wantIdx, gotIdx)
		}
	}
}
