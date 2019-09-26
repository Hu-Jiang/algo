package easy

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{
			s:    "",
			want: true,
		},
		{
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			s:    "race a car",
			want: false,
		},
		{
			s:    "0P",
			want: false,
		},
	}

	for i, tt := range tests {
		got := isPalindrome(tt.s)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
