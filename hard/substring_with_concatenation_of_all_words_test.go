package hard

import (
	"reflect"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	tests := []struct {
		s     string
		words []string
		want  []int
	}{
		{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, nil},
	}

	for i, tt := range tests {
		got := findSubstring(tt.s, tt.words)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
