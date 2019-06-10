package medium

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{3, []string{
			"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()"},
		},
		{0, nil},
	}

	for i, tt := range tests {
		got := generateParenthesis(tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}

}
