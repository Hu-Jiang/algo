package medium

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{nil, nil, nil},
		{NewList(1, 2, 3), nil, NewList(1, 2, 3)},
		{NewList(2, 4, 3), NewList(5, 6, 4), NewList(7, 0, 8)},
	}

	for i, tt := range tests {
		got := addTwoNumbers(tt.l1, tt.l2)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("%d. got %v, want %v", i, got, tt.want)
		}
	}
}
