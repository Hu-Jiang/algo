package medium

import (
	"reflect"
	"testing"
)

func TestRemoveNthFormEnd(t *testing.T) {
	tests := []struct {
		head *ListNode
		n    int
		want *ListNode
	}{
		{NewList(1, 2, 3, 4, 5), 2, NewList(1, 2, 3, 5)},
		{NewList(1, 2, 3, 4, 5), 5, NewList(2, 3, 4, 5)},
		{NewList(1, 2, 3, 4, 5), 1, NewList(1, 2, 3, 4)},
	}

	for i, tt := range tests {
		got := removeNthFromEnd(tt.head, tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%d. got %v, want %v", i, got, tt.want)
		}
	}
}
