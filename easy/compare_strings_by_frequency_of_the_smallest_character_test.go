package easy

import (
	"reflect"
	"testing"
)

func TestNumSmallerByFrequency(t *testing.T) {
	tests := []struct {
		queries []string
		words   []string
		want    []int
	}{
		{
			queries: []string{"cbd"},
			words:   []string{"zaaaz"},
			want:    []int{1},
		},
		{
			queries: []string{"bbb", "cc"},
			words:   []string{"a", "aa", "aaa", "aaaa"},
			want:    []int{1, 2},
		},
	}

	for i, tt := range tests {
		got := numSmallerByFrequency(tt.queries, tt.words)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
