package medium

import (
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		{nil, nil},
		{NewList(1, 1), nil},
		{NewList(1, 1, 1), nil},
		{NewList(1, 2, 3, 3, 4, 4, 5), NewList(1, 2, 5)},
		{NewList(1, 1, 1, 2, 3), NewList(2, 3)},
	}

	for i, tt := range tests {
		got := deleteDuplicates(tt.head)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
