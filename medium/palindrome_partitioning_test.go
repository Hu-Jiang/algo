package medium

import (
	"reflect"
	"testing"
)

func TestPartitionString(t *testing.T) {
	tests := []struct {
		s    string
		want [][]string
	}{
		{
			s:    "",
			want: nil,
		},
		{
			s: "aab",
			want: [][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
	}

	for i, tt := range tests {
		got := partitionString(tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
