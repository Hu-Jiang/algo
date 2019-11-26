package medium

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		s, want string
	}{
		{" 		", ""},
		{"the sky is blue", "blue is sky the"},
		{"	hello world!	", "world! hello"},
		{"a good	example", "example good a"},
	}

	for i, tt := range tests {
		got := reverseWords(tt.s)
		if got != tt.want {
			t.Fatalf("#%d. got: %v, want: %v", i, got, tt.want)
		}
	}
}
