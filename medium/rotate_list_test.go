package medium

import (
	"reflect"
	"testing"
)

func TestRotateRight(t *testing.T) {
	tests := []struct {
		head *ListNode
		k    int
		want *ListNode
	}{
		{nil, 10, nil},
		{NewList(1), 10, NewList(1)},
		{NewList(1, 2), 0, NewList(1, 2)},
		{NewList(1, 2), 2, NewList(1, 2)},
		{NewList(1, 2, 3, 4, 5), 2, NewList(4, 5, 1, 2, 3)},
		{NewList(0, 1, 2), 4, NewList(2, 0, 1)},
	}

	for i, tt := range tests {
		got := rotateRight(tt.head, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%d. got %v, want %v", i, got, tt.want)
		}
	}
}
