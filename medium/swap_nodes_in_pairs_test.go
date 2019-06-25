package medium

import (
	"reflect"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		{nil, nil},
		{NewList(1, 2, 3, 4), NewList(2, 1, 4, 3)},
		{NewList(1, 2, 3), NewList(2, 1, 3)},
	}

	for i, tt := range tests {
		got := swapPairs(tt.head)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%d. got %v, want %v", i, got, tt.want)
		}
	}
}
