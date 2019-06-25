package hard

import (
	"reflect"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		head *ListNode
		k    int
		want *ListNode
	}{
		{nil, 10, nil},
		{NewList(1, 2, 3, 4, 5), 0, NewList(1, 2, 3, 4, 5)},
		{NewList(1, 2, 3, 4, 5), 2, NewList(2, 1, 4, 3, 5)},
		{NewList(1, 2, 3, 4, 5), 3, NewList(3, 2, 1, 4, 5)},
		{NewList(1, 2, 3, 4, 5), 5, NewList(5, 4, 3, 2, 1)},
	}

	for i, tt := range tests {
		got := reverseKGroup(tt.head, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%d. got %v, want %v", i, got, tt.want)
		}
	}
}
