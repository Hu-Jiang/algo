package hard

import (
	"reflect"
	"strings"
	"testing"
)

func TestFullJustify(t *testing.T) {
	tests := []struct {
		words    []string
		maxWidth int
		want     []string
	}{
		{nil, 3, []string{"   "}},
		{
			words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 16,
			want: []string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		},
		{
			words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth: 16,
			want: []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
		},
		{
			words: []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain",
				"to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			maxWidth: 20,
			want: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
		{
			words:    []string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"},
			maxWidth: 16,
			want: []string{
				"ask   not   what",
				"your country can",
				"do  for  you ask",
				"what  you can do",
				"for your country",
			},
		},
	}

	for i, tt := range tests {
		got := fullJustify(tt.words, tt.maxWidth)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got \n%q, \nwant \n%q", i, strings.Join(got, ","), strings.Join(tt.want, ","))
		}
	}
}
