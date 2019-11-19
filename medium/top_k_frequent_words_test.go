package medium

import (
	"reflect"
	"testing"
)

func TestTopKFrequentWords(t *testing.T) {
	tests := []struct {
		words []string
		k     int
		want  []string
	}{
		{
			words: nil,
			k:     10,
			want:  nil,
		},
		{
			words: []string{"hello"},
			k:     0,
			want:  nil,
		},
		{
			words: []string{"a", "b", "c", "d"},
			k:     2,
			want:  []string{"a", "b"},
		},
		{
			words: []string{"i", "love", "leetcode", "i", "love", "coding"},
			k:     2,
			want:  []string{"i", "love"},
		},
		{
			words: []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"},
			k:     4,
			want:  []string{"the", "is", "sunny", "day"},
		},
	}

	for i, tt := range tests {
		got := topKFrequentWords(tt.words, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
