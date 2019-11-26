package hard

import (
	"reflect"
	"testing"
)

func TestWordBreak(t *testing.T) {
	tests := []struct {
		s        string
		wordDict []string
		want     []string
	}{
		{
			s:        "catsanddog",
			wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			want:     []string{"cats and dog", "cat sand dog"},
		},
		{
			s:        "pineapplepenapple",
			wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			want:     []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"},
		},
		{
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			want:     nil,
		},
	}

	for i, tt := range tests {
		got := wordBreak(tt.s, tt.wordDict)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
