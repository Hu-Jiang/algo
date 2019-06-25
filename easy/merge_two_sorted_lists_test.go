package easy

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{nil, nil, nil},
		{NewList(1, 2, 3), nil, NewList(1, 2, 3)},
		{NewList(1, 2, 4), NewList(1, 3, 4), NewList(1, 1, 2, 3, 4, 4)},
	}

	for i, test := range tests {
		got := mergeTwoLists(test.l1, test.l2)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("%d. want %v, got %v", i, test.want, got)
		}
	}
}
