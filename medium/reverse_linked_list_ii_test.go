package medium

import (
	"reflect"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		head *ListNode
		m, n int
		want *ListNode
	}{
		{
			head: NewList(1),
			m:    1,
			n:    1,
			want: NewList(1),
		},
		{
			head: NewList(1, 2, 3),
			m:    1,
			n:    2,
			want: NewList(2, 1, 3),
		},
		{
			head: NewList(1, 2, 3, 4, 5),
			m:    2,
			n:    4,
			want: NewList(1, 4, 3, 2, 5),
		},
	}

	for i, tt := range tests {
		got := reverseBetween(tt.head, tt.m, tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
