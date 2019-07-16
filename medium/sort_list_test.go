package medium

import (
	"reflect"
	"testing"
)

func TestSortList(t *testing.T) {
	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		{nil, nil},
		{NewList(4, 2, 1, 3), NewList(1, 2, 3, 4)},
		{NewList(-1, 5, 3, 4, 0), NewList(-1, 0, 3, 4, 5)},
	}

	for i, tt := range tests {
		got := sortList(tt.head)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%d. got %v, want %v", i, got, tt.want)
		}
	}
}
